package service_web

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-api/app/common/common"
	db "gin-api/app/common/db"
	curl "gin-api/app/common/http"
	"net"
	"strconv"
	"time"

	. "gin-api/app/models/web"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

/*
  - 查询商品条形码
    1、查询商品接口，如果有则使用此数据，没有则查询数据数据，数据库数据没有则返回空
    2、接口有数据，数据库没有数据则插入，数据库有数据则对比数据是否一致，不一致则更新数据库数据
    3、图片每次查询接口则下载，等比缩小，并同步到数据库
    4、记录每个接口总的查询次数，每个用户总的查询次数，用户每个时段的查询次数。
  - @param {*gin.Context} c
  - @param {string} searchInput
  - @return {*}
*/
func QueryBarCode(searchInput string) map[string]any {
	// 获取商品名称、品牌、供应商、商品分类、规格
	product := getProductBaseInfo(searchInput)

	// 获取商品扩展信息 价格
	productExtend := getProductExtendInfo(searchInput)

	// 拼接数据 处理数据逻辑
	return getFinalyData(searchInput, product, productExtend)
}

// 拼接数据 处理数据逻辑
func getFinalyData(searchInput string, product map[string]any, productExtend map[string]any) map[string]any {
	// var result map[string]any
	resultInterface := map[string]any{}
	resultModel := map[string]any{}

	productLen := len(product)

	// 接口返回数据
	if productLen > 0 {
		resultInterface = getResInterfaceMap(searchInput, product, productExtend)
	}

	// 数据库返回数据
	barCode := BarCode{}
	barCodeModel := db.Model.Where("bar_code = ?", searchInput).First(&barCode)
	barCodeModelEmptyStatus := 1
	if errors.Is(barCodeModel.Error, gorm.ErrRecordNotFound) {
		barCodeModelEmptyStatus = 0

		// 新增商品条码数据
		if productLen > 0 {
			addNewBarCode(resultInterface)
		}
	}

	barCodeMap := map[string]any{}
	barCodeString, _ := json.Marshal(barCode)
	_ = json.Unmarshal([]byte(barCodeString), &barCodeMap)
	barCodeMap["price"] = common.PriceFillZero(fmt.Sprint(barCodeMap["price"]))
	resultModel = barCodeMap

	// 判断返回数据
	if productLen == 0 && barCodeModelEmptyStatus == 0 { // 都为空
		// 记录未查到的商品条形码
		recordSearchNoResultCode(searchInput)

		return map[string]any{}
	} else if productLen > 0 && barCodeModelEmptyStatus == 0 { // 接口有值 数据库没值
		return resultInterface
	} else if productLen == 0 && barCodeModelEmptyStatus == 1 { // 接口没值 数据库有值
		return resultModel
	}

	// 接口和数据库查询都有数据，做逻辑处理，对比如果不一致则更新数据库中数据
	isEqual := common.CompareTwoMapInterface(resultInterface, resultModel)
	if !isEqual {
		cTime := time.Now().Format("2006-01-02 15:04:05")
		resultInterface["updated_at"] = cTime
		db.Model.Table("bar_code").Where("bar_code = ?", barCode.BarCode).Updates(resultInterface)
	}

	return resultInterface
}

// 获取商品名称、品牌、供应商、商品分类接口
func getProductBaseInfo(searchInput string) map[string]any {
	// 调用商品条形码接口
	// params := map[string]any{}
	params := gin.H{
		"PageSize":   "30",
		"PageIndex":  "1",
		"SearchItem": searchInput,
	}
	queryParams := common.HttpBuildQuery(params, "")

	ip := common.GenIpaddr()
	header := map[string]string{
		"CLIENT-IP":       ip,
		"X-FORWARDED-FOR": ip,
	}
	url := "https://bff.gds.org.cn/gds/searching-api/ProductService/ProductListByGTIN?" + queryParams
	productString, _ := curl.GetWithHeader(url, header)

	code := gjson.Get(productString, "Code")
	items := gjson.Get(productString, "Data.Items").Array()

	// var product map[string]any
	product := gin.H{}
	if code.Int() == 1 && len(items) > 0 {
		product = items[0].Value().(map[string]any)
	}

	return product
}

// 获取商品价格、规格接口
func getProductExtendInfo(searchInput string) map[string]any {
	params := gin.H{
		"barcode":    searchInput,
		"app_id":     "ohnvf8eponbjhjwv",
		"app_secret": "K0pUajE3a2w3MnlXanFhNU5nREpNdz09",
	}
	queryParams := common.HttpBuildQuery(params, "")
	url := "https://www.mxnzp.com/api/barcode/goods/details?" + queryParams
	productExtendString, _ := curl.Get(url)

	code := gjson.Get(productExtendString, "code")

	// var productExtend map[string]any
	productExtend := gin.H{}
	if code.Int() == 1 {
		data := gjson.Get(productExtendString, "data").Value().(map[string]any)
		productExtend = data
	}

	return productExtend
}

// 处理返回数据
func InitRes(searchRes map[string]any) map[string]any {
	resFieldMap := gin.H{
		"bar_code":       "商品条形码",
		"name":           "商品名称",
		"short_name":     "简称",
		"image":          "图片",
		"brand":          "品牌",
		"supplier":       "供应商",
		"classification": "商品分类",
		"status":         "条码状态",
		"price":          "价格",
		"specification":  "规格",
	}

	result := gin.H{}
	for key, value := range searchRes {
		if key == "status" {
			if fmt.Sprint(value) == "1" {
				value = "有效"
			} else {
				value = "无效"
			}
		}

		if key == "Price" {

		}

		keyName := resFieldMap[key]
		result[fmt.Sprint(keyName)] = value
	}

	return result
}

// 新增商品条码数据
func addNewBarCode(resultInterface map[string]any) {
	int8_status, _ := strconv.Atoi(fmt.Sprint(resultInterface["status"]))
	float64_price, _ := strconv.ParseFloat(fmt.Sprint(resultInterface["price"]), 64)

	db.Model.Create(&BarCode{
		BarCode:        fmt.Sprint(resultInterface["bar_code"]),
		Name:           fmt.Sprint(resultInterface["name"]),
		ShortName:      fmt.Sprint(resultInterface["short_name"]),
		Image:          fmt.Sprint(resultInterface["image"]),
		Brand:          fmt.Sprint(resultInterface["brand"]),
		Supplier:       fmt.Sprint(resultInterface["supplier"]),
		Classification: fmt.Sprint(resultInterface["classification"]),
		Status:         int8(int8_status),
		Specification:  fmt.Sprint(resultInterface["specification"]),
		Price:          float64_price,
	})
}

// 拼接商品接口返回数据
func getResInterfaceMap(searchInput string, product map[string]any, productExtend map[string]any) map[string]any {
	resultInterface := map[string]any{
		"bar_code":       searchInput,
		"name":           product["description"],
		"short_name":     product["keyword"],
		"image":          "",
		"brand":          product["brandcn"],
		"supplier":       product["firm_name"],
		"classification": product["gpcname"],
		"status":         product["gtinstatus"],
		"specification":  product["specification"],
	}

	if len(productExtend) > 0 {
		price := common.GetPrice(fmt.Sprint(productExtend["price"]))
		resultInterface["price"] = price
	}

	return resultInterface
}

// 记录未查到的商品条形码
func recordSearchNoResultCode(searchInput string) {
	barCodeSearchNoResult := BarCodeSearchNoResult{}
	barCodeSearchNoResultModel := db.Model.Where("bar_code = ?", searchInput).First(&barCodeSearchNoResult)
	if errors.Is(barCodeSearchNoResultModel.Error, gorm.ErrRecordNotFound) {
		ipStr := common.GetIp()
		ip := net.ParseIP(ipStr)
		ipInt := common.IpToInt(ip)

		db.Model.Create(&BarCodeSearchNoResult{
			BarCode:  searchInput,
			ListId:   1,
			UserId:   0,
			IpString: fmt.Sprint(ip),
			IpInt:    ipInt,
			Status:   0,
		})
	}
}
