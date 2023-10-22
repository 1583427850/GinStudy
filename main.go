package main

import (
	"GinCode/note"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(c *gin.Context) {
	//返回json格式
	c.JSON(http.StatusOK, "hello gin")
	c.JSON(http.StatusOK, gin.H{"name": "tom"})
}

func main() {
	//使用gin的默认http引擎
	//禁用输出颜色
	gin.DisableConsoleColor()
	//创建一个日志文件
	f, _ := os.Create("F:/GoCode/GinCode/gin.log")
	//设置将日志写入对应的文件
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	//result风格配置路由规则----------------------------------
	r.GET("/hello", hello)

	//配置rustful风格请求格式
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, "GET")
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POST")
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, "PUT")
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, "DELETE")
	})
	//-----------------------------------------------------------

	r.GET("/asciiJson", note.AsciiJSON)
	//将目录下的所有html页面扫描进来
	r.LoadHTMLGlob("Html/*")
	r.GET("/html", note.Html)

	r.Use(func(e *gin.Context) {
		fmt.Println("1")

	}, func(e *gin.Context) {
		fmt.Println("2")
	})

	push(r)

	r.GET("/jsonp", note.JSONP)
	r.POST("/bind", note.BindObject)

	r.POST("query", note.Query)

	r.POST("/uploadfile", note.RecveiveFile)
	//启动服务
	//这样子启动会启动所有中间件，可以用r := gin.New()启动，就不会启动Logger 和 Recovery 中间件

	//创建一个新的路由组，可以用来统一设置函数
	//accouts := gin.H{"root": "123456", "admin": "123456"}

	//如果一个路由下面，有多个子路由，可以用路由组来判断,
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/get", func(c *gin.Context) { c.JSON(http.StatusOK, "查询所有商品") })
		shopGroup.POST("/post", func(c *gin.Context) { c.JSON(http.StatusOK, "新增商品") })
		shopGroup.PUT("/put", func(c *gin.Context) { c.JSON(http.StatusOK, "修改商品") })
		shopGroup.DELETE("/delete", func(c *gin.Context) { c.JSON(http.StatusOK, "删除商品") })
	}
	//authentication := r.Group("/admin", gin.BasicAuth(gin.Accounts{"root": "123456", "admin": "123456"}))
	//
	//authentication.GET("/select", func(c *gin.Context) {
	//	username := c.Query("username")
	//	password := c.Query("password")
	//})

	//中间件,可以定义在用户和接口之间,用来统一鉴权和统一的日志管理等等功能,就传入一个.HandlerFunc()函数
	authentication := r.Group("/admin", func(c *gin.Context) { fmt.Println("先经过统一认证处理(可以添加认证处理).......") })
	{
		authentication.GET("/get", func(c *gin.Context) {

			fmt.Println("获取用户中.....")
			c.JSON(http.StatusOK, "获取用户")
		})
	}

	//如果是没有匹配到的路由，那么会执行到这个方法
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "没有存在这个网页")
	})
	gin.DisableConsoleColor()

	//传递cookie
	r.GET("/cookie", note.GetCookie)

	r.Run(":8080")

}

func push(r *gin.Engine) {
	var html = template.Must(template.New("https").Parse(`
	<html>
	<head>
	  <title>Https Test</title>
	  <script src="/assets/app.js"></script>
	</head>
	<body>
	  <h1 style="color:red;">Welcome, Ginner!</h1>
	</body>
	</html>
	`))
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	//利用push可以在刚刚打开网站的时候，就将所有js代码发送到前端，然后缓存到本地
	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}

		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
}

func ReiveMultipart() {

}
func Spend(c *gin.Context) {
	fmt.Println("中间件处理逻辑前.....")
	//继续执行接下来的中间件
	c.Next()
	fmt.Println("执行完成后续所有逻辑后回来")
}
