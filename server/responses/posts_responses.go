package responses

import (
	"time"
)

type PostResponse struct {
	Id       int          `json:"id" bson:"_id"`
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	Writer   UserResponse `json:"writer"`
	CreateAt time.Time    `json:"create_at"`
	UpdateAt time.Time    `json:"update_at"`
}