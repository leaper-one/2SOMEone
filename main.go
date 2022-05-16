package main

import (
	"2SOMEone/core"
	_ "2SOMEone/docs"
	"2SOMEone/middlewares"
	"2SOMEone/service"
	"2SOMEone/util"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	pb "2SOMEone/grpc/user"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

const (
	SUCCESS = 200
	FAIL = 500
)

var (
	dbc         = util.OpenDB("./2-some-one.db")
	userService = service.NewUserService(dbc)
	noteService = service.NewNoteService(dbc)
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}
	
// Sent phone message code.
func (s *UserService) SentMessageCode(ctx context.Context, in *pb.SentMessageCodeRequest) (*pb.SentMessageCodeResponse, error) {
	fmt.Printf("SentMessageCode: %v\n", in)
	code, err := userService.SendPhoneCode(ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	fmt.Printf("code: %v\n", code)
	return &pb.SentMessageCodeResponse{Code: SUCCESS, Msg: "success."}, nil
}

// @title BubbleBox
// @version 1.0
// @description NULL
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://api.2some.one
// @BasePath 
func main() {
	// dbc := util.OpenDB("./2-some-one.db")
	// gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// r := gin.Default()
	// r.Use(middlewares.Cors())

	// initRoutes(r)
	// r.Run(":3002")


	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// pb.RegisterGreeterServer(grpcServer, newServer())
	pb.RegisterUserServiceServer(grpcServer, &UserService{})
	err = grpcServer.Serve(lis)
	if err != nil {
		return 
	}
}

func initRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.GET("/phonecode", phoneCode)
	r.POST("/signup", signUp)
	r.POST("/setinfo", middlewares.JWTAuthMiddleware(), setInfo)
	r.POST("/auth", auth)
	r.GET("/to/:rame", middlewares.JWTAuthMiddleware(), to)
	r.GET("/me", middlewares.JWTAuthMiddleware(), me)
	r.GET("/user/:user_id", user)
	r.GET("/note/:note_id", middlewares.JWTAuthMiddleware(),note)
	r.GET("/notes/received",middlewares.JWTAuthMiddleware(), notesReceived)
	r.GET("/notes/sent",middlewares.JWTAuthMiddleware(), notesSent)
}

// @Summary 获取手机验证码
// @Description 向用户手机发送验证码
// @Tags user
// @Produce json
// @Param phone query string true "手机号"
// @Success 200 {string} string "{"code": 2000, "msg": ""}"
// @Router /phonecode [get]
func phoneCode(c *gin.Context) {
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
}

// @Summary 用户注册
// @Description 用户注册
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code": 2000, "msg": "注册成功"}"
// @Router /signup [post]
func signUp(c *gin.Context) {
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
}

// @Summary 设置用户信息
// @Description 设置用户信息
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code": 2000, "msg": "成功"}"
// @Router /setinfo [post]
func setInfo(c *gin.Context) {
	ctx := context.Background()
	user_id := c.MustGet("user_id").(string)
	var setinfo core.UserForMe
	err := c.ShouldBind(&setinfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	err = userService.SetInfo(ctx, user_id, &setinfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "成功",
	})
}

// @Summary 用户授权
// @Description 用户授权
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code": 2000, "msg": "success", "data": {"token": ""}}"
// @Router /auth [post]
func auth(c *gin.Context) {
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
}

// @Summary 创建 Note
// @Description 创建 Note
// @Tags note
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code": 2000, "msg": ""}"
// @Router /note [post]
func to(c *gin.Context) {
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
}

// @Summary 获取本用户信息
// @Description 获取本用户信息
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code": 2000, "msg": "success", "data": {"user": {}}}"
// @Router /user [get]
func me(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"user": user},
	}) // 返回信息
}

// @Summary 访问用户信息
// @Description 访问用户信息
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code": 2000, "msg": "success", "data": {"user": {}}}"
// @Router /user/{user_id} [get]
func user(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"user": user},
	}) // 返回信息
}

// @Summary 获取用户的Note
// @Description 获取用户的Note
// @Tags note
// @Accept json
// @Produce json
// @Success 200 {string} string "{"code": 2000, "msg": "success", "data": {"notes": []}}"
// @Router /note/{note_id} [get]
func note(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"note": note},
	})
}

// @Summary 获取收到的Note
// @Description 获取收到的Note
// @Tags note
// @Accept json
// @Produce json
// @Router /note/received [get]
func notesReceived(c *gin.Context) {
	ctx := context.Background()
	user_id := c.MustGet("user_id")
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	notes, count, err := noteService.RecipientGet(ctx, offset, limit, user_id.(string))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   2000,
		"msg":    "成功",
		"count":  count,
		"notes":  notes,
		"offset": offset,
		"limit":  limit,
	})
}

// @Summary 获取发出的Note
// @Description 获取发出的Note
// @Tags note
// @Accept json
// @Produce json
// @Router /note/sent [get]
func notesSent(c *gin.Context) {
	ctx := context.Background()
	user_id := c.MustGet("user_id")
	offset_str := c.DefaultQuery("offset", "1")
	offset, _ := strconv.Atoi(offset_str)
	limit_str := c.DefaultQuery("limit", "20")
	limit, _ := strconv.Atoi(limit_str)
	notes, count, err := noteService.SenderGet(ctx, offset, limit, user_id.(string))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   2000,
		"msg":    "成功",
		"count":  count,
		"notes":  notes,
		"offset": offset,
		"limit":  limit,
	})
}
