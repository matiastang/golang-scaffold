/*
 * @Author: matiastang
 * @Date: 2022-04-19 17:14:35
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-24 17:03:52
 * @FilePath: /golang-scaffold/src/main.go
 * @Description:
 */
package main

import (
	"fmt"

	"scaffold/app/test"
	"scaffold/routers"
)

func init() {
	fmt.Println("learn main init")
}

func main() {
	fmt.Println("learn main main")
	// 加载多个APP的路由配置
	routers.Include(test.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		// TODO: - 服务启动失败
	}
}
