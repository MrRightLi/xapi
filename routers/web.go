package routers

import (
	"github.com/gin-gonic/gin"
	"xapi/app/controllers"
	"os"
	"io"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
	"xapi/tools"
)

var DB = make(map[string]string)

func InitRouter() *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery())
	gin.DisableConsoleColor()
	f, _ := os.Create("storage/log/gin.log")
	gin.DefaultErrorWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(gin.Logger())

	router.GET("/test", func(context *gin.Context) {
		logger := new(tools.Logger)
		logger.InitLogger()
		logger.Info("hello logger-info")

		var hosts = []string{"127.0.0.1:2181"} //server端host
		conn, _, err := zk.Connect(hosts, time.Second*5)
		if err != nil {
			fmt.Println(err)
			return
		}
		var path = "/xapi"
		var data = []byte("hello zk")
		var flags = 0
		//flags有4种取值：
		//0:永久，除非手动删除
		//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
		//zk.FlagSequence  = 2:会自动在节点后面添加序号
		//3:Ephemeral和Sequence，即，短暂且自动添加序号
		var acls = zk.WorldACL(zk.PermAll) //控制访问权限模式

		p, err_create := conn.Create(path, data, int32(flags), acls)
		if err_create != nil {
			fmt.Println(err_create)
			return
		}
		logger.Info("hello logger-info" + p)

		defer conn.Close()
	})

	admin := router.Group("/admin")
	{
		router.GET("getWorkOrders", controllers.GetOrders)

		admin.POST("/post", func(context *gin.Context) {
			message := context.PostForm("message")
			nick := context.DefaultPostForm("nick", "anonymous")
			context.JSON(200, map[string]string{
				"status":  "posted",
				"message": message,
				"nick":    nick,
			})
		})
	}
	return router
}

func routerdemo() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		c.String(200, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return r
}
