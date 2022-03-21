package main

import (
	"2SOMEone/core"
	"2SOMEone/middlewares"
	"2SOMEone/service"
	"2SOMEone/util"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// go checkImgByDate(".jpg")
	// go checkImgByDate(".png")
	// baseapi := "https://imgapi.leaper.one"
	dbc := util.OpenDB("./2-some-one.db")

	userServer := service.NewUserService(dbc)

	r := gin.Default()
	r.Use(middlewares.Cors())

	r.GET("/phonecode", func(c *gin.Context) {
		ctx := context.Background()
		// phone := c.Request.FormValue("phone")
		phone := c.Query("phone")
		// code, err := util.SendMail(email)
		code, err := userServer.SendPhoneCode(ctx, phone)
		if err != nil {
			c.String(http.StatusBadGateway, fmt.Sprintf("参数错误:%v", phone))
			return
		}
		fmt.Printf("code: %v\n", code)
		fmt.Printf("err: %v\n", err)
		// accounter.CreateByEmail(ctx, email, code)
		c.String(http.StatusOK, fmt.Sprintf("已向 %s 发送验证码 ", phone))
	})

	r.POST("/signup", func(c *gin.Context) {
		ctx := context.Background()
		// remail := c.Request.FormValue("email")
		rphone := c.Request.FormValue("phone")
		rcode := c.Request.FormValue("code")
		password := c.Request.FormValue("password")

		sign_user := &core.SignUpUser{Phone: rphone, Code: rcode, Password: password}
		// user, err := userServer.SendPhoneCode.SignUpByEmail(ctx, remail, rcode, password)
		user, _, err := userServer.SignUp(ctx, sign_user)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("注册失败：%v", user.Phone))
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.POST("/auth", func(c *gin.Context) {
		ctx := context.Background()
		var login_user core.LoginUser
		err := c.ShouldBind(&login_user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  "无效的参数",
			})
			return
		}
		// user, _ := .FindByEmail(ctx, login_user.Email)

		// 校验用户名和密码是否正确
		token, err := userServer.Auth(ctx, &login_user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  "鉴权失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": token},
		})
		})

	r.Run(":3002")
}
