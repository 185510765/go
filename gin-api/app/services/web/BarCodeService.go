package service_web

import (
	"encoding/json"
	"fmt"
	"gin-api/app/common/common"
	"gin-api/app/common/db"
	curl "gin-api/app/common/http"
	. "gin-api/app/models/web"
)

// 查询商品条形码
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
	jsonParse := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &jsonParse)
	if err != nil {
		return map[string]interface{}{}
		// return make([]int, 0)
	}

	Code, CodeExist := jsonParse["Code"].(float64)
	if Code != 1 || !CodeExist {
		return map[string]interface{}{}
	}

	// 查询数据
	barCode := []BarCode{}
	db.DB.Where("bar_code = ?", searchInput).Take(&barCode)

	fmt.Println(jsonParse["Data"])
	fmt.Println(barCode)

	return map[string]interface{}{}
}
