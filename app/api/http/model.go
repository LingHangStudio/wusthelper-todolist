package http

type TodolistRespItem struct {
	Name       string `json:"name"`
	Time       string `json:"time"`
	Comment    string `json:"comment"`
	Uuid       string `json:"uuid"`
	CreateTime string `json:"createTime"`
}

type TodolistReqItem struct {
	Name       *string `json:"name"`
	Time       *string `json:"time"`
	Comment    *string `json:"comment"`
	Uuid       *string `json:"uuid"`
	CreateTime *string `json:"createTime"`
}
