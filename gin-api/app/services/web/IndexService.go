package service_web

import (
	"fmt"
	"gin-api/app/common/cache"
	"gin-api/app/common/common"
	db "gin-api/app/common/db"
	. "gin-api/app/models/web"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexService struct{}

// 查询api列表数据
func (indexService *IndexService) GetApiList() interface{} {
	list := []List{}
	db.Model.Select("id,name,img").Where("is_show = ?", 1).Order("sort desc").Find(&list)

	for index, value := range list {
		list[index].Img = "/public/static/web/img/" + value.Img
	}

	return list
}

// 获取一条list数据
func (indexService *IndexService) GetOneList(id int) (interface{}, string) {
	info := List{Id: id}
	db.Model.First(&info)

	info.Img = "/public/static/web/img/" + info.Img

	return info, info.SearchInput
}

// 查询校验 查询限制（根据ip），限制查询频率（5秒）、24小时总的查询次数（100次）
func (indexService *IndexService) ValidateSearch(searchInput string, inputIsExist bool, keySuffix string, tips string) (int, string) {
	// 参数是否存在
	if !inputIsExist {
		return 0, ""
	}

	// 输入是否为空
	if searchInput == "" {
		return 0, tips + "不能为空"
	}

	// 24小时总查询次数
	var numKey = "userIpNum:" + keySuffix
	num, _ := cache.RedisClient.Get(numKey).Result()
	int_num, _ := strconv.Atoi(num)
	if int_num > 100 {
		return 0, "今日已到了查询上限"
	}

	// 查询频率
	var timeKey = "userIpTime:" + keySuffix
	res, _ := cache.RedisClient.SetNX(timeKey, 1, 5*time.Second).Result()
	if !res {
		return 0, "查询太过频繁，请稍后在试"
	}

	// 校验正常 可以正常查询 写入redis
	keyExist, _ := cache.RedisClient.Exists(numKey).Result()
	if keyExist == 1 {
		cache.RedisClient.Incr(numKey).Result()
	} else {
		_, todayEndUnix := common.GetTodayStartTimeAndEndTime()
		cTimeUnix := time.Now().Unix()
		todaySurplusTimeUnix := todayEndUnix - cTimeUnix

		cache.RedisClient.Set(numKey, 1, time.Duration(todaySurplusTimeUnix)*time.Second)
	}

	return 1, ""
}

/*
  - 查询接口操作
    先根据id类型查询不同的接口，然后再对比本地数据库数据
    本地如果没有则添加，如果有则判断有变动则修改。
  - @param {int} id
  - @return {*}
*/
func (indexService *IndexService) QueryData(id int, searchInput string) map[string]interface{} {
	var barCodeService BarCodeService

	switch id {
	case 1: // 商品条形码
		return barCodeService.QueryBarCode(searchInput)
	case 2: // IP信息查询
		fmt.Println(2)
	case 3: // 全国天气查询
		fmt.Println(3)
	default:
		fmt.Println("default")
	}

	return gin.H{}
}

// IP信息查询
func (indexService *IndexService) QueryIp(searchInput string) {

}

// 全国天气查询
func (indexService *IndexService) QueryWeather(searchInput string) {

}

// *******************************************************************************************************
