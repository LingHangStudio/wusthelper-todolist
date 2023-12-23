package dao

import (
	"go.uber.org/zap"
	"wusthelper-todolist-service/app/model"
	"wusthelper-todolist-service/library/ecode"
	"wusthelper-todolist-service/library/log"
)

const (
	_GetUserTodolistSql    = "select * from todolist where owner = ? and deleted = false"
	_GetTodolistItemSql    = "select * from todolist where id = ? and deleted = false"
	_DeleteTodolistItemSql = "update todolist set deleted = true where id = ? and owner = ?"
)

// SaveTodolistItem 保存一个todo项
func (d *Dao) SaveTodolistItem(basic *model.TodolistItem) (int64, error) {
	result, err := d.db.InsertOne(basic)
	if err != nil {
		log.Warn("Error happened when insert one todolist record.", zap.Error(err))
		return 0, ecode.DaoOperationErr
	}

	return result, nil
}

// GetUserTodolist 获取指定用户的todolist
func (d *Dao) GetUserTodolist(uid string) (*[]model.TodolistItem, error) {
	result := make([]model.TodolistItem, 0)
	err := d.db.SQL(_GetUserTodolistSql, uid).Find(&result)
	if err != nil {
		log.Warn("Error happened when select one todolist record.", zap.Error(err))
		return nil, ecode.DaoOperationErr
	}

	return &result, nil
}

// GetTodolistItemById 按照id获取某个todo项
func (d *Dao) GetTodolistItemById(id uint64) (*model.TodolistItem, error) {
	var result model.TodolistItem
	has, err := d.db.SQL(_GetTodolistItemSql, id).Get(&result)
	if err != nil {
		log.Warn("Error happened when select one todolist record.", zap.Error(err))
		return nil, ecode.DaoOperationErr
	}

	if has {
		return &result, nil
	} else {
		return nil, nil
	}
}

func (d *Dao) UpdateTodolistItem(id uint64, owner string, updatedItem *model.TodolistItem) (int64, error) {
	count, err := d.db.
		Where("id = ?", id).
		Where("owner = ?", owner).
		Update(updatedItem)

	if err != nil {
		log.Warn("Error happened when update one todolist record.", zap.Error(err))
		return 0, ecode.DaoOperationErr
	}

	return count, nil
}

func (d *Dao) DeleteTodolistItem(id uint64, owner string) error {
	_, err := d.db.Exec(_DeleteTodolistItemSql, id, owner)
	if err != nil {
		log.Warn("Error happened when delete one todolist record.", zap.Error(err))
		return ecode.DaoOperationErr
	}

	return nil
}
