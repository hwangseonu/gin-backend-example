package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/models"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func NewPost(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)

	var payload struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	_, err := models.CreatePost(user, payload.Title, payload.Content)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.Status(200)
}

func GetAllPost(c *gin.Context) {
	if posts, err := models.GetPosts(); err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
	} else {
		result := make([]gin.H, len(posts))
		for i, v := range posts {
			result[i] = gin.H{
				"post_id": v.ID,
				"author": getNicknameById(v.Author),
				"title": v.Title,
				"content": v.Content,
				"create_at": v.CreateAt,
				"update_at": v.UpdateAt,
				"comments": []models.Comment{},
			}
		}
		c.JSON(200, result)
	}
}

func GetPost(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("pid"))

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if post, err := models.GetPostById(uint32(pid)); err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"post_id": post.ID,
			"author": getNicknameById(post.Author),
			"title": post.Title,
			"content": post.Content,
			"create_at": post.CreateAt,
			"update_at": post.UpdateAt,
			"comments": getCommentsResponse(post.Comments),
		})
	}
}

func DeletePost(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)
	pid, err1 := strconv.Atoi(c.Param("pid"))

	if ErrorHandle(c, 400, err1) {
		return
	}
	post, err := models.GetPostById(uint32(pid))
	if ErrorHandle(c, 404, err) {
		return
	}
	if post.Author != user.ID {
		c.JSON(403, gin.H{"message": "cannot access this resource"})
		return
	}
	if ErrorHandle(c, 500, models.RemovePost(uint32(pid))) {
		return
	}
	c.Status(200)

}

func AddComment(c *gin.Context) {
	var payload struct {
		Content string `json:"content" binding:"required"`
	}
	u, _ := c.Get("user")
	user := u.(*models.User)
	pid, err1 := strconv.Atoi(c.Param("pid"))
	err2 := c.ShouldBindJSON(&payload)
	if ErrorHandle(c, 400, err1) || ErrorHandle(c, 400, err2){
		return
	}
	post, err := models.GetPostById(uint32(pid))
	if ErrorHandle(c, 404, err) {
		return
	}
	if ErrorHandle(c, 500, post.AddComment(user, payload.Content)) {
		return
	}
	c.Status(201)
}

func DeleteComment(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)
	pid, err1 := strconv.Atoi(c.Param("pid"))
	cid, err2 := strconv.Atoi(c.Param("cid"))

	if ErrorHandle(c, 400, err1) || ErrorHandle(c, 400, err2) {
		return
	}

	post, err3 := models.GetPostById(uint32(pid))
	if ErrorHandle(c, 404, err3) {
		return
	}

	comment := post.GetComment(uint32(cid))
	if comment == nil {
		c.JSON(404, gin.H{"message": "Could not find comment"})
		return
	}
	if comment.Author != user.ID {
		c.JSON(403, gin.H{"message": "cannot access this resource"})
		return
	}
	post.RemoveComment(comment.ID)
	c.Status(200)
}

func ErrorHandle(c *gin.Context, code int, e error) bool {
	if e != nil {
		c.AbortWithStatusJSON(code, gin.H{"message": e.Error()})
		return true
	}
	return false
}

func getNicknameById(id bson.ObjectId) string {
	u, err := models.GetUserById(id)
	if err != nil {
		return "탈퇴한 사용자"
	} else {
		return u.Nickname
	}
}

func getCommentsResponse(comments []models.Comment) []gin.H {
	result := make([]gin.H, len(comments))
	for i, v := range comments {
		result[i] = gin.H{
			"comment_id": v.ID,
			"author": getNicknameById(v.Author),
			"content": v.Content,
			"create_at": v.CreateAt,
			"update_at": v.UpdateAt,
		}
	}
	return result
}
