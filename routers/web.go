package routers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"xapi/app/controllers"
	"xapi/config"
	"xapi/tools"
)

var DB = make(map[string]string)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	gin.DisableConsoleColor()
	f, _ := os.Create("storage/log/gin.log")
	//gin.DefaultErrorWriter = io.MultiWriter(f)
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(gin.Logger())

	admin := router.Group("/admin")
	{
		admin.POST("/test", controllers.Test)
		admin.GET("/getWorkOrders", controllers.GetOrders)
		admin.GET("/getRepairTasks", controllers.GetTasks)

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
		test.GET("/test", func(context *gin.Context) {
			response := gin.H{
				"test": 1,
			}
			context.JSON(200, response)
		})
		test.GET("/ip", func(context *gin.Context) {
			addrs, _ := net.InterfaceAddrs()
			var gInnerIp []string
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						gInnerIp = append(gInnerIp, ipnet.IP.String())
					}
				}
			}
			response := gin.H{
				"ipv4: ": gInnerIp,
			}
			context.JSON(200, response)
		})
		test.POST("/logs", func(context *gin.Context) {
			data, _ := ioutil.ReadAll(context.Request.Body)
			log.Printf("ctx.Request.body: %v", string(data))

			//context.Request.ParseForm()
			//for k, v := range context.Request.PostForm {"
			//	log.Printf("k:%v\n", k)
			//	log.Printf("v:%v\n", v)
			//}

			context.JSON(200, map[string]string{
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
