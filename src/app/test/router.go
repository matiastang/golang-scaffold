/*
 * @Author: matiastang
 * @Date: 2022-04-18 17:01:55
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-19 18:35:34
 * @FilePath: /golang-scaffold/src/app/test/router.go
 * @Description: router
 */
package test

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/", Index)
	e.GET("/userinfo/:id", GetUserInfo)
}
