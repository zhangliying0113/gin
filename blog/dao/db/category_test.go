package db

import (
	"testing"
)

func init() {
	// parseTime=true 将 mysql 的时间类型自动解析为 go 结构体中的时间类型，不加报错
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		t.Errorf("get category failed, err:%v\n", err)
		return
	}
	t.Logf("category:%#v", category)

}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}
	if len(categoryIds) != len(categoryList) {
		t.Errorf("get category list failed, len of categorylist:%d ids len:%d\n", len(categoryList), len(categoryIds))
		return
	}
	for _, v := range categoryList{
		t.Logf("id: %d category: %#v\n", v.CategoryId, v)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}
	for _, v := range categoryList{
		t.Logf("category: %#v\n", v)
	}
}