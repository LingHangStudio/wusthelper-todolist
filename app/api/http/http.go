package http

import (
	"github.com/gin-gonic/gin"
	"time"
	"wusthelper-todolist-service/app/conf"
	"wusthelper-todolist-service/app/middleware/auth"
	"wusthelper-todolist-service/app/service"
	"wusthelper-todolist-service/library/ecode"
	"wusthelper-todolist-service/library/token"
)

var (
	srv      *service.Service
	jwt      *token.Token
	timezone *time.Location
)

func NewEngine(c *conf.Config, baseUrl string) *gin.Engine {
	timezone, _ = time.LoadLocation("Asia/Shanghai")

	engine := gin.Default()
	rootRouter := engine.RouterGroup.Group(baseUrl)
	//rootRouter.Use(gin.LoggerWithWriter(*log.DefaultWriter().))

	setupOuterRouter(rootRouter)

	srv = service.New(c)
	jwt = token.New(c.Server.TokenSecret, c.Server.TokenTimeout)

	return engine
}

func setupOuterRouter(group *gin.RouterGroup) {
	todolist := group.Group("/lh", auth.UserTokenCheck)
	{
		todolist.GET("/list-countdown-new", getUserTodolist)
		todolist.GET("/list-countdown", getUserTodolist)
		todolist.POST("/add-countdown", addTodolistItem)
		//todolist.POST("/add-pub-countdown", addSharedTodolistItemPub)
		todolist.GET("/add-shared-countdown", copyTodolistItem)
		todolist.POST("/modify-countdown", modifyTodolistItem)
		todolist.GET("/del-countdown", deleteTodolistItem)
	}
}

func responseEcode(c *gin.Context, code ecode.Code) {
	response(c, code.Code(), code.Message(), nil)
}

func response(c *gin.Context, code int, msg string, data any) {
	resp := gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	c.JSON(200, resp)
}

func getUid(c *gin.Context) uint64 {
	uid := c.GetUint64("uid")

	return uid
}

func _getStudent(c *gin.Context) string {
	student := c.GetString("student")
	return student
}
