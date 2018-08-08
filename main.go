package main

import (
	_ "github.com/go-sql-driver/mysql"
	"xapi/routers"
)

const HTTP_PORT = ":8000"

func main() {
	r := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(HTTP_PORT)
}
