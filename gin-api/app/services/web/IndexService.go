package service_web

import (
	db "gin-api/app/common/db"
	. "gin-api/app/models/web"
)

// 查询api列表数据
func GetApiList() interface{} {
	list := []List{}
	db.DB.Select("id,name,img").Where("is_show = ?", 1).Order("sort desc").Find(&list)

	for index, value := range list {
		list[index].Img = "/public/static/web/img/" + value.Img
	}

	return list
}
