package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tangren1998/ddd-cheanArch/user/adapter"
)

func GetUserById(c *gin.Context) {
	id := c.Param("id") // 提取id
	// todo: 参数验证
	controller := adapter.UserController{C: c}
	controller.GetUserByID(id)
}
