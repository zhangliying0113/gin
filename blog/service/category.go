package service

import (
	"fmt"
	"github.com/zhangliying0113/gin/blog/dao/db"
	"github.com/zhangliying0113/gin/blog/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	return
}

