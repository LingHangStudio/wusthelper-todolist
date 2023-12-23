package model

import "time"

type TodolistItem struct {
	Id         uint64     `json:"id"`
	Owner      string     `json:"owner"`   // 拥有者uid
	Title      *string    `json:"title"`   // 标题
	Time       *time.Time `json:"time"`    // 时间
	Comment    *string    `json:"comment"` // 备注
	CreateTime time.Time  `json:"create_time"`
	UpdateTime time.Time  `json:"update_time" xorm:"updated"`
	Deleted    bool       `json:"deleted"`
}

func (TodolistItem) TableName() string {
	return "todolist"
}
