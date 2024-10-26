package models

type ReqCreateTinyURL struct {
	URL  string `json:"url"`                         // 原始URL
	Time int    `json:"time" binding:"gte=1,lte=30"` // 有效期 单位 天
}
