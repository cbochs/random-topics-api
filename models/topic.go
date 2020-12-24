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

// SubmitTopicInput ...
type SubmitTopicInput struct {
	Code  string `json:"code" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}

// UpdateTopicInput ...
type UpdateTopicInput struct {
	Topic string `json:"topic" binding:"required"`
}

// GetTopicResult ...
type GetTopicResult struct {
	Code        string `json:"code"`
	SessionCode string `json:"session_code"`
	Topic       string `json:"topic"`
	Type        string `json:"type"`
	OK          bool   `json:"ok"`
	Reason      string `json:"reason"`
}
