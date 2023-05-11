package main

import (
	"fmt"
	"net/http"

	"github.com/go-ldap/ldap/v3"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/authenticate", func(c echo.Context) error {
		//获取POST请求的参数
		account := c.FormValue("account")
		password := c.FormValue("password")
		fmt.Println(account)
		fmt.Println(password)
		//连接到LDAP测试服务器
		conn, err := ldap.Dial("tcp", "ldap.forumsys.com:389")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer conn.Close()

		//构建绑定DN
		bindDN := "cn=" + account + ",dc=example,dc=com"

		//验证用户凭据
		err = conn.Bind(bindDN, password)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		//如果验证成功，返回JSON响应
		response := map[string]bool{"authenticate": true}
		return c.JSON(http.StatusOK, response)
	})

	// 启动服务器
	e.Logger.Fatal(e.Start(":8080"))
}
