package service

import (
	"github.com/yitter/idgenerator-go/idgen"
	"time"
	"wusthelper-todolist-service/app/model"
	"wusthelper-todolist-service/library/ecode"
)

func (s *Service) AddTodolistItem(owner string, item TodolistItemData) (uint64, error) {
	id := uint64(idgen.NextId())
	now := time.Now()
	entity := model.TodolistItem{
		Id:         id,
		Owner:      owner,
		Title:      item.Title,
		Time:       item.Time,
		Comment:    item.Comment,
		CreateTime: now,
		UpdateTime: now,
		Deleted:    false,
	}

	_, err := s.dao.SaveTodolistItem(&entity)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) GetUserTodolist(uid string) (*[]model.TodolistItem, error) {
	result, err := s.dao.GetUserTodolist(uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetTodolistItem(id uint64) (*model.TodolistItem, error) {
	result, err := s.dao.GetTodolistItemById(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) CopyTodolistItem(newOwner string, originId uint64) error {
	item, err := s.dao.GetTodolistItemById(originId)
	if err != nil {
		return err
	}

	if item == nil {
		return ecode.DataNotExists
	}

	item.Id = uint64(idgen.NextId())
	item.Owner = newOwner

	now := time.Now()
	item.CreateTime = now
	item.UpdateTime = now

	_, err = s.dao.SaveTodolistItem(item)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ModifyTodolistItem(owner string, id uint64, item TodolistItemData) error {
	entity := model.TodolistItem{
		Title:      item.Title,
		Time:       item.Time,
		Comment:    item.Comment,
		UpdateTime: time.Now(),
	}

	_, err := s.dao.UpdateTodolistItem(id, owner, &entity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTodolistItem(owner string, id uint64) error {
	err := s.dao.DeleteTodolistItem(id, owner)
	if err != nil {
		return err
	}

	return nil
}
