package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/zhangliying0113/gin/blog/controller"
	"github.com/zhangliying0113/gin/blog/dao/db"
)

func main() {
	// 注册路由
	router := gin.Default()
	// 定义数据库信息，用于初始化
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 嵌入性能监控工具
	ginpprof.Wrapper(router)
	// 加载静态资源
	router.Static("/static/", "./static")
	// 加载 html 模板
	router.LoadHTMLGlob("/views/*")
	// 加载首页
	router.GET("/", controller.IndexHandle)

	//router := gin.Default()
	//dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	//err := db.Init(dns)
	//if err != nil {
	//	panic(err)
	//}
	//ginpprof.Wrapper(router)
	//router.Static("/static/", "./static")
	//router.LoadHTMLGlob("views/*")
	//router.GET("/", controller.IndexHandle)
	//router.GET("/category/", controller.CategoryList)
	//router.GET("/article/new/", controller.NewArticle)
	//router.POST("/article/submit/", controller.ArticleSubmit)
	//router.GET("/article/detail/", controller.ArticleDetail)
	//router.POST("/upload/file/", controller.UploadFile)
	//router.GET("/leave/new/", controller.LeaveNew)
	//router.GET("/about/me/", controller.AboutMe)
	//router.POST("/comment/submit/", controller.CommentSubmit)
	//router.POST("/leave/submit/", controller.LeaveSubmit)
	//_ = router.Run(":8000")
}
