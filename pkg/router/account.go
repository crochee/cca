// Date: 2021/9/22

// Package router
package router

import (
	"github.com/gin-gonic/gin"

	"caty/api/v1/account"
)

func registerAccount(v1Router *gin.RouterGroup) {
	v1Router.POST("/account", account.Register)
	v1Router.GET("/account", account.List)
	v1Router.PATCH("/account/:id", account.Update)
	v1Router.GET("/account/:id", account.Retrieve)
	v1Router.DELETE("/account/:id", account.Delete)
	v1Router.POST("/account/login", account.Login)
}
