/*
 * @Author: matiastang
 * @Date: 2022-04-19 17:20:54
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-19 18:35:11
 * @FilePath: /golang-scaffold/src/app/test/test.go
 * @Description: 测试
 */
package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

type UserInfo struct {
	Id   string `json:"userId"`
	Name string
}

func GetUserInfo(c *gin.Context) {
	userId := c.Param("id")
	info := UserInfo{
		userId,
		"name_" + userId,
	}
	fmt.Println(info)
	c.JSON(http.StatusOK, info)
}
