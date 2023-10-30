package global

import (
	"time"
)

type Model struct {
	ID        int       `json:"id" form:"id" gorm:"primaryKey;AUTO_INCREMENT"`                 // Primary key ID
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"column:create_at"`            // Creation time
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at"`           // Update time
	IsDeleted int       `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;default:0"` // Soft delete: 0-active 1-deleted
}
