import (
	"database-example/model"

	"github.com/google/uuid"
)

type BlogPostDTO struct{
	Id uuid.UUID `json:id`
	UserId uuid.UUID `json:id`
	Title string `json:"username" gorm:"not null"`
	Description string `json:"username" gorm:"not null"`
	Date time.Time `json:"date" gorm:"not null"`
	CommentIds []uuid.UUID `json:"commentIds" gorm:"type:uuid[]"`
}