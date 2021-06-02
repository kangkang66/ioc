package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	api,err := InitApi()
	if err != nil {
		panic(err)
	}

	g := gin.Default()
	g.GET("/task/info",api.TaskInfo)
	g.GET("/sign_in/info",api.SignInInfo)
	g.GET("/luckywheel/info",api.LuckywheelInfo)

	g.Run(":8080")
}
