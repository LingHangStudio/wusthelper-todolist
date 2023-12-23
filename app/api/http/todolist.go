package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
	"wusthelper-todolist-service/app/service"
	"wusthelper-todolist-service/library/ecode"
	"wusthelper-todolist-service/library/log"
)

const (
	_timeLayoutNoSecond = "2006-01-02 15:04"
	_timeLayout         = "2006-01-02 15:04:05"
)

// addTodolistItem 添加一条todo
// @Summary 添加一条todo
// @Description 添加一条todo
// @Tags
// @Accept application/json
// @Produce application/json
// @Param token header string false "token"
// @Param object body http.TodolistReqItem false "todo信息"
// @Success 200 {uint64} id
// @Router /add-countdown [post]
func addTodolistItem(c *gin.Context) {
	var request = new(TodolistReqItem)
	err := c.BindJSON(request)
	if err != nil {
		log.Debug("参数校验错误", zap.Error(err))
		responseEcode(c, ecode.ParamWrong)
	}

	todoTime := time.Now()
	if request.Time != nil {
		// 这里有坑，直接使用time.Parse()时区会错误，会直接偏移8小时，因此要指定时区来解析
		todoTime, err = time.ParseInLocation(_timeLayoutNoSecond, *request.Time, timezone)
		if err != nil {
			responseEcode(c, ecode.ParamWrong)
			return
		}
	}

	data := service.TodolistItemData{
		Title:   request.Name,
		Time:    &todoTime,
		Comment: request.Comment,
	}

	student := _getStudent(c)
	id, err := srv.AddTodolistItem(student, data)
	if err != nil {
		log.Error("Save todolist item failed.", zap.Error(err))
		responseEcode(c, ecode.ServerErr)
		return
	}

	response(c, ecode.OK.Code(), "ok", strconv.FormatUint(id, 10))
}

// getUserTodolist 获取用户当前的todolist
// @Summary 获取用户当前的todolist
// @Description 获取用户当前的todolist
// @Tags
// @Accept */*
// @Produce application/json
// @Param token header string false "token"
// @Router /list-countdown [get]
func getUserTodolist(c *gin.Context) {
	student := _getStudent(c)

	resultList, err := srv.GetUserTodolist(student)
	if err != nil || resultList == nil {
		log.Error("Getting todolist item failed.", zap.Error(err))
		responseEcode(c, ecode.ServerErr)
		return
	}

	respItemList := make([]TodolistRespItem, len(*resultList))
	for i, item := range *resultList {
		if item.Comment != nil {
			respItemList[i].Comment = *item.Comment
		}

		if item.Title != nil {
			respItemList[i].Name = *item.Title
		}

		if item.Time != nil {
			respItemList[i].Time = item.Time.Format(_timeLayoutNoSecond)
		} else {
			respItemList[i].Time = time.Now().Format(_timeLayoutNoSecond)
		}

		// 注意这里的creatTime格式是yyyy-mm-dd hh:mm:ss，是带有秒的
		respItemList[i].CreateTime = item.CreateTime.Format(_timeLayout)
		respItemList[i].Uuid = strconv.FormatUint(item.Id, 10)
	}

	response(c, ecode.OK.Code(), "ok", respItemList)
}

// deleteTodolistItem 删除一个todo项
// @Summary 删除一个todo项
// @Description 删除一个todo项
// @Tags
// @Accept */*
// @Produce application/json
// @Param token header string false "token"
// @Param uuid query string false "欲删除的id"
// @Router /del-countdown [get]
func deleteTodolistItem(c *gin.Context) {
	student := _getStudent(c)
	id := _getTodolistItemQueryId(c)
	if id == 0 {
		responseEcode(c, ecode.ParamWrong)
		return
	}

	err := srv.DeleteTodolistItem(student, id)
	if err != nil {
		log.Error("Delete todolist item failed.", zap.Error(err))
		responseEcode(c, ecode.ServerErr)
		return
	}

	response(c, ecode.OK.Code(), "ok", nil)
}

// modifyTodolistItem 修改一个todo项
// @Summary 修改一个todo项
// @Description 修改一个todo项
// @Tags
// @Accept */*
// @Produce application/json
// @Param token header string false "token"
// @Param uuid query string false "欲修改的id"
// @Param object body http.TodolistReqItem false "todo信息"
// @Router /modify-countdown [get]
func modifyTodolistItem(c *gin.Context) {
	student := _getStudent(c)
	request := new(TodolistReqItem)
	err := c.BindJSON(request)
	if err != nil {
		log.Debug("参数校验错误", zap.Error(err))
		responseEcode(c, ecode.ParamWrong)
		return
	}

	if request.Uuid == nil {
		log.Debug("参数校验错误", zap.Error(err))
		responseEcode(c, ecode.ParamWrong)
		return
	}

	data := service.TodolistItemData{
		Title: request.Name,
		//Time:    request.Time,
		Comment: request.Comment,
	}
	if request.Time != nil {
		t, err := time.ParseInLocation(_timeLayoutNoSecond, *request.Time, timezone)
		if err != nil {
			log.Debug("参数校验错误", zap.Error(err))
			responseEcode(c, ecode.ParamWrong)
			return
		}

		data.Time = &t
	}

	id, err := strconv.ParseUint(*request.Uuid, 10, 64)
	if err != nil {
		log.Debug("参数校验错误", zap.Error(err))
		responseEcode(c, ecode.ParamWrong)
		return
	}

	err = srv.ModifyTodolistItem(student, id, data)
	if err != nil {
		log.Error("Update todolist item failed.", zap.Error(err))
		responseEcode(c, ecode.ServerErr)
		return
	}

	result, err := srv.GetTodolistItem(id)
	if err != nil {
		log.Error("Getting todolist item failed.", zap.Error(err))
		responseEcode(c, ecode.ServerErr)
		return
	}

	resp := TodolistRespItem{
		Name:       *result.Title,
		Time:       result.Time.Format(_timeLayoutNoSecond),
		Comment:    *result.Comment,
		Uuid:       strconv.FormatUint(result.Id, 10),
		CreateTime: result.CreateTime.Format(_timeLayoutNoSecond),
	}

	response(c, ecode.OK.Code(), "ok", resp)
}

// copyTodolistItem 复制一个todo项（添加分享的倒计时）
// @Summary 复制一个todo项（添加分享的倒计时）
// @Description 复制一个todo项（添加分享的倒计时）
// @Tags
// @Accept */*
// @Produce application/json
// @Param token header string false "token"
// @Param uuid query string false "欲添加的id"
// @Router /modify-countdown [get]
func copyTodolistItem(c *gin.Context) {
	student := _getStudent(c)
	id := _getTodolistItemQueryId(c)

	err := srv.CopyTodolistItem(student, id)
	if err != nil {
		log.Error("Copy todolist item failed.", zap.Error(err))
		responseEcode(c, ecode.ServerErr)
		return
	}

	response(c, ecode.OK.Code(), "ok", nil)
}

func _getTodolistItemQueryId(c *gin.Context) uint64 {
	_id, has := c.GetQuery("uuid")
	if !has {
		return 0
	}

	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil {
		return 0
	}

	return id
}
