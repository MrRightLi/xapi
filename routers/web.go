package routers

import (
	"github.com/gin-gonic/gin"
	"xapi/app/controllers"
	"os"
	"io"
	"xapi/tools"
	"xapi/config"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var DB = make(map[string]string)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	gin.DisableConsoleColor()
	f, _ := os.Create("storage/log/gin.log")
	gin.DefaultErrorWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(gin.Logger())

	admin := router.Group("/admin")
	{
		admin.GET("/getWorkOrders", controllers.GetOrders)

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

	test := router.Group("test")
	{
		test.GET("/ymal", func(context *gin.Context) {
			logger := new(tools.Logger)
			logger.InitLogger()
			logger.Info("hello logger-info")

			conf := new(config.Yaml)
			yamlFile, err := ioutil.ReadFile("config/test.yaml")
			if err != nil {
				logger.Error("yamlFile.Get err")
			}
			err = yaml.Unmarshal(yamlFile, conf)
			if err != nil {
				logger.Error("Unmarshal: %v")
			}
			context.JSON(200, gin.H{"User": conf.Mysql.User})
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
