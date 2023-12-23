package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wusthelper-todolist-service/app/conf"
	"wusthelper-todolist-service/library/ecode"
	_token "wusthelper-todolist-service/library/token"
)

var (
	jwt *_token.Token
	dev bool
)

func Init(c *conf.Config) {
	jwt = _token.New(c.Server.TokenSecret, c.Server.TokenTimeout)
	dev = c.Server.Env == conf.DevEnv
}

func UserTokenCheck(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		rejectRequest(c)
		return
	}

	claims, valid := jwt.GetClaimVerify(token)
	if (!dev && !valid) || claims == nil {
		rejectRequest(c)
		return
	}

	_student := (*claims)["StuNum"]
	student, ok := (_student).(string)
	if !ok {
		rejectRequest(c)
		return
	}
	c.Set("student", student)

	uid, ok := ((*claims)["StuId"]).(string)
	if !ok {
		rejectRequest(c)
		return
	}
	uint64uid, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		rejectRequest(c)
		return
	}

	c.Set("uid", uint64uid)

	c.Next()
}

func rejectRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": ecode.TokenInvalid.Code(),
		"msg":  "token invalid",
	})

	return
}
