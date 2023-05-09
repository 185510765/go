package service_web

import (
	"fmt"
	"gin-api/app/common/common"
	db "gin-api/app/common/db"
	curl "gin-api/app/common/http"

	. "gin-api/app/models/web"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
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
	// 通过接口获取商品信息
	product := getProductInfo(searchInput)

	return product
}

// 通过接口获取商品信息
func getProductInfo(searchInput string) map[string]any {
	// 获取商品名称、品牌、供应商、商品分类、
	product := getProductBaseInfo(searchInput)

	// 获取商品扩展信息 价格、规格
	productExtend := getProductExtendInfo(searchInput)

	// 拼接数据
	// var result map[string]any
	result := gin.H{}
	if len(product) > 0 {
		result = gin.H{
			"BarCode":        searchInput,
			"Name":           product["description"],
			"ShortName":      product["keyword"],
			"Image":          "",
			"Brand":          product["brandcn"],
			"Supplier":       product["firm_name"],
			"Classification": product["gpcname"],
			"Status":         product["gtinstatus"],
			"Price":          "",
			"Specification":  "",
		}

		if len(productExtend) > 0 {
			price := common.GetPrice(fmt.Sprint(productExtend["price"]))

			result["Price"] = price
			result["Specification"] = productExtend["standard"]
		}
	} else {
		barCode := BarCode{}
		db.DB.Where("bar_code = ?", searchInput).First(&barCode)

		result = gin.H{
			"BarCode":        barCode.BarCode,
			"Name":           barCode.Name,
			"ShortName":      barCode.ShortName,
			"Image":          barCode.Image,
			"Brand":          barCode.Brand,
			"Supplier":       barCode.Supplier,
			"Classification": barCode.Classification,
			"Status":         barCode.Status,
			"Price":          barCode.Price,
			"Specification":  barCode.Specification,
		}
	}

	return result
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
		"BarCode":        "商品条形码",
		"Name":           "商品名称",
		"ShortName":      "简称",
		"Image":          "图片",
		"Brand":          "品牌",
		"Supplier":       "供应商",
		"Classification": "商品分类",
		"Status":         "条码状态",
		"Price":          "价格",
		"Specification":  "规格",
	}

	result := gin.H{}
	for key, value := range searchRes {
		keyName := resFieldMap[key]
		result[fmt.Sprint(keyName)] = value
	}

	return result
}
