package service_web

import (
	"fmt"
	"gin-api/app/common/common"
	db "gin-api/app/common/db"
	curl "gin-api/app/common/http"

	. "gin-api/app/models/web"

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
func QueryBarCode(searchInput string) interface{} {
	// 调用商品条形码接口
	// m := make(map[string]interface{})
	params := map[string]interface{}{
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
	jsonString, _ := curl.GetWithHeader(url, header)

	code := gjson.Get(jsonString, "Code")
	// data := gjson.Get(jsonString, "Data")
	items := gjson.Get(jsonString, "Data.Items").Array()

	var product interface{}
	if code.Int() == 1 && len(items) > 0 {
		product = items[0]
	}

	// 查询接口为空则查询数据库
	if product == nil {
		product = []BarCode{}
		db.DB.Where("bar_code = ?", searchInput).Take(&product)
		fmt.Println(product)
	}

	fmt.Println(product)

	return map[string]interface{}{}
}
