package responses

import (
	"github.com/hwangseonu/gin-backend-example/server/models"
	"time"
)

type PostResponse struct {
	Id       int               `json:"id" bson:"_id"`
	Title    string            `json:"title"`
	Content  string            `json:"content"`
	Writer   UserResponse      `json:"writer"`
	Comments []CommentResponse `json:"comments"`
	CreateAt time.Time         `json:"create_at"`
	UpdateAt time.Time         `json:"update_at"`
}

type CommentResponse struct {
	Id       int          `json:"id"`
	Content  string       `json:"content"`
	Writer   UserResponse `json:"writer"`
	CreateAt time.Time    `json:"create_at"`
	UpdateAt time.Time    `json:"update_at"`
}

func NewPostResponse(post *models.Post) PostResponse {
	user := models.FindUserById(post.Writer)
	res := PostResponse{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
		Writer: UserResponse{
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
		},
		Comments: make([]CommentResponse, 0),
		CreateAt: post.CreateAt,
		UpdateAt: post.UpdateAt,
	}
	for _, com:= range post.Comments {
		u := models.FindUserById(com.Writer)
		res.Comments = append(res.Comments, CommentResponse{
			Id: com.Id,
			Content: com.Content,
			Writer: UserResponse{
				Username: u.Username,
				Nickname: u.Nickname,
				Email: u.Email,
			},
			CreateAt: com.CreateAt,
			UpdateAt: com.UpdateAt,
		})
	}
	return res

}