package note

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AsciiJSON 可以用来将字符串转换为有转义的非 ASCII 字符的 ASCII-only JSON。
func AsciiJSON(c *gin.Context) {
	data := map[string]string{
		"lang": "java语言",
		"tar":  "<h1>",
	}
	//{"lang":"java\u8bed\u8a00","tar":"\u003ch1\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

// Html 返回html页面
func Html(c *gin.Context) {
	//因为前面在启动类已经将所有html页面扫描进来了，最后一个参数可以传入模板里面还没传入的值 {{ .name }}
	c.HTML(http.StatusOK, "htmlTest.html", gin.H{
		"name": "user",
	})
}

func Pusher(c *gin.Context) {
	//main里面
}

func JSONP(c *gin.Context) {
	//可以向不同域的服务器请求数据
	c.JSONP(http.StatusOK, "<h1>123</h1>")
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func BindObject(c *gin.Context) {
	var form LoginForm
	//将对应参数绑定到form这个变量中
	//如果绑定没出现错误，那么就可以进入里面认证
	//if c.ShouldBind(&form) == nil {
	//	if form.User == "user" && form.Password == "password" {
	//		c.JSON(200, gin.H{"status": "you are logged in"})
	//	} else {
	//		c.JSON(401, gin.H{"status": "unauthorized"})
	//	}
	//} else {
	//	c.JSON(200, gin.H{"status": "not bind"})
	//
	//}

	if c.ShouldBind(&form) == nil {
		if form.User == "小张" && form.Password == "123456" {
			c.JSON(http.StatusOK, "登录成功")
		} else {
			c.JSON(http.StatusOK, "账号或密码错误")
		}
	} else {
		c.JSON(http.StatusOK, "绑定失败")
	}
}

func Query(c *gin.Context) {
	//获取请求信息
	id := c.Query("id")
	page := c.DefaultQuery("page", "1")
	username := c.PostForm("username")

	password := c.DefaultPostForm("password", "1234")
	fmt.Println(123)
	fmt.Printf("id:%v page: %v username:%v password:%v ", id, page, username, password)
}

// 保存文件到磁盘中
func RecveiveFile(c *gin.Context) {
	file, err := c.FormFile("file")
	//如果是多文本，用MultipartForm接收

	//发送格式
	//curl -X POST http://localhost:8080/upload \
	//-F "upload[]=@/Users/appleboy/test1.zip" \
	//-F "upload[]=@/Users/appleboy/test2.zip" \
	form, err := c.MultipartForm()
	//获取里面的数据数组
	files := form.File["upload[]"]

	dist := "F://abc"
	for _, file = range files {
		c.SaveUploadedFile(file, dist)
	}
	if err != nil {
		panic("文件获取出现错误")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	//将文件保存到磁盘中
	c.SaveUploadedFile(file, dist)

}

func Reader(c *gin.Context) {
	//获取图片
	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	if err != nil {
		panic("获取文件失败")
	}
	contentType := response.Header.Get("Content-Type")

	//传输到前端下载
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	//将图片传输给前端
	c.DataFromReader(http.StatusOK, response.ContentLength, contentType, response.Body, extraHeaders)

}

func GetCookie(c *gin.Context) {
	testCookie, _ := c.Cookie("testCookie")
	fmt.Println(testCookie)
}
