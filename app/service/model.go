package service

import "time"

type TodolistItemData struct {
	Title   *string
	Time    *time.Time
	Comment *string
}
