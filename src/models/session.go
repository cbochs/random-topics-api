package models

// Session ...
type Session struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

// SessionInput ...
type SessionInput struct {
	Code string `json:"code" binding:"required"`
}

type SessionClosedResult struct {
	Code   string `json:"code"`
	Status string `json:"status"`
	OK     bool   `json:"ok"`
	Reason string `json:"reason"`
}
