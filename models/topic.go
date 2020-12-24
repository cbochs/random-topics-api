package models

// Topic ...
type Topic struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Code      string  `json:"code"`
	Submitted string  `json:"submitted"`
	Assigned  string  `json:"assigned"`
	SessionID uint    `json:"session_id"`
	Session   Session `gorm:"foreignKey:SessionID"`
}

// TopicInput ... Code will be either a session code (submission) or
// topic code (updating)
type TopicInput struct {
	Code  string `json:"code" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}

type TopicResult struct {
	Code        string `json:"code"`
	SessionCode string `json:"session_code"`
	Topic       string `json:"topic"`
	Type        string `json:"type"`
	OK          bool   `json:"ok"`
	Reason      string `json:"reason"`
}
