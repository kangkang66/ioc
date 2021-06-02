package api

import (
	"github.com/gin-gonic/gin"
	"gomod/internal/module/luckywheel"
	"gomod/internal/module/sign_in"
	"gomod/internal/module/task"
)

func ApiProvider(
	taskServ *task.Service,
	signInServ *sign_in.Service,
	luckywheelServ *luckywheel.Service,
	) (*Api,error) {
	return &Api{
		taskServ: taskServ,
		signInServ: signInServ,
		luckywheelServ: luckywheelServ,
	},nil
}

type Api struct {
	taskServ *task.Service
	signInServ *sign_in.Service
	luckywheelServ *luckywheel.Service
}

func (a *Api) TaskInfo(c *gin.Context) {
	taskId := c.Query("task_id")
	data,err := a.taskServ.Info(c.Request.Context(), taskId)
	if err != nil {
		c.JSON(500,map[string]interface{}{"err":err})
	}else{
		c.JSON(200,data)
	}
}

func (a *Api) SignInInfo(c *gin.Context) {
	taskId := c.Query("task_id")
	data,err := a.taskServ.Info(c.Request.Context(), taskId)
	if err != nil {
		c.JSON(500,map[string]interface{}{"err":err})
	}else{
		c.JSON(200,data)
	}
}

func (a *Api) LuckywheelInfo(c *gin.Context) {
	taskId := c.Query("task_id")
	data,err := a.taskServ.Info(c.Request.Context(), taskId)
	if err != nil {
		c.JSON(500,map[string]interface{}{"err":err})
	}else{
		c.JSON(200,data)
	}
}
