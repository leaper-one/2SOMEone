package main

import (
	"2SOMEone/core"
	"2SOMEone/middlewares"
	"2SOMEone/service"
	"2SOMEone/util"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// baseapi := "https://imgapi.leaper.one"
	dbc := util.OpenDB("./2-some-one.db")
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	userService := service.NewUserService(dbc)
	noteService := service.NewNoteService(dbc)

	r := gin.Default()
	r.Use(middlewares.Cors())

	r.GET("/phonecode", func(c *gin.Context) {
		ctx := context.Background()
		phone := c.Query("phone")
		code, err := userService.SendPhoneCode(ctx, phone)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  err.Error(),
			})
			// c.String(http.StatusBadGateway, fmt.Sprintf("参数错误:%v", phone))
			return
		}
		fmt.Printf("code: %v\n", code)
		fmt.Printf("err: %v\n", err)
		// accounter.CreateByEmail(ctx, email, code)
		// c.String(http.StatusOK, fmt.Sprintf("已向 %s 发送验证码 ", phone))
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "验证码已发送",
		})
	})

	r.POST("/signup", func(c *gin.Context) {
		ctx := context.Background()
		var logup_user core.SignUpUser
		err := c.ShouldBind(&logup_user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  "无效的参数",
			})
			return
		}

		user, err := userService.SignUpByPhone(ctx, &logup_user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  2000,
			"phone": user.Phone,
			"msg":   "注册成功",
		})
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

		// 校验用户名和密码是否正确
		token, err := userService.Auth(ctx, &login_user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				// "msg":  "鉴权失败",
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": token},
		})
	})

	r.POST("/to/:rname", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		ctx := context.Background()
		sender_id := c.MustGet("user_id")
		rname := c.Param("rname")
		var tnote core.Note
		err := c.ShouldBind(&tnote)
		if err != nil || rname == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  err.Error(),
			})
			return
		}

		tnote.Sender = sender_id.(string)
		err = noteService.Create(ctx, tnote.ForStore(), rname)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"code": 2002,
				"msg":  "失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "成功",
		})
	})

	r.GET("/me", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		ctx := context.Background()
		user_id := c.MustGet("user_id").(string)
		user, err := userService.GetMe(ctx, user_id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, user) // 返回信息
	})

	r.GET("/user/:user_name", func(c *gin.Context) {
		ctx := context.Background()
		user_name := c.Param("user_name")
		user, err := userService.VisitUser(ctx, user_name)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, user) // 返回信息
	})

	r.GET("/note/:note_id", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		ctx := context.Background()
		note_id := c.Param("note_id")
		note, err := noteService.GetByID(ctx, note_id)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 2002,
				"msg":  err.Error(),
			})
			return
		}
		note, err = note.ForRead()
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 2002,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, note)

	})

	r.GET("/notes/received/", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		ctx := context.Background()
		user_id := c.MustGet("user_id")
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		notes, count, err := noteService.RecipientGet(ctx, offset, limit, user_id.(string))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "成功",
			"count": count,
			"notes": notes,
			"offset": offset,
			"limit": limit,
		})
		// c.String(http.StatusOK,"yes")
	})

	r.GET("/notes/sent", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		ctx := context.Background()
		user_id := c.MustGet("user_id")
		offset_str:=c.DefaultQuery("offset", "1")
		offset, _ := strconv.Atoi(offset_str)
		limit_str:=c.DefaultQuery("limit", "20")
		limit, _ := strconv.Atoi(limit_str)
		notes, count, err := noteService.SenderGet(ctx, offset, limit, user_id.(string))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "成功",
			"count": count,
			"notes": notes,
			"offset": offset,
			"limit": limit,
		})

	})

	r.Run(":3002")
}
