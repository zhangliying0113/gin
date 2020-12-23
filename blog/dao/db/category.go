package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhangliying0113/gin/blog/model"
)

func InsertCategory(category *model.Category)(categoryId int64, err error) {
	sqlStr := "insert into category(category_name, category_no) values (?, ?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return categoryId, err
}

// 获取单个文章分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "select id, category_name, category_no from category where id = ?"
	err = DB.Get(category, sqlStr, id)
	return
}

// 获取多个文章分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	// 构建查询 sql
	sqlStr, args, err := sqlx.In("select id, category_name, category_no from category where id in (?)", categoryIds)
	if err != nil {
		return
	}
	// 执行查询语句
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlStr)
	return
}
