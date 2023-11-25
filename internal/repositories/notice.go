package repositories

import (
	"gorm.io/gorm"
	"new-mall/internal/models"
)

type NoticeRepository struct {
	*gorm.DB
}

func NewNoticeRepository(db *gorm.DB) *NoticeRepository {
	return &NoticeRepository{
		DB: db,
	}
}

func (r *NoticeRepository) GetNoticeById(id uint) (notice *models.Notice, err error) {
	err = r.DB.Model(&models.Notice{}).Where("id=?", id).First(&notice).Error
	return
}

func (r *NoticeRepository) CreateNotice(notice *models.Notice) error {
	return r.DB.Model(&models.Notice{}).Create(&notice).Error
}
