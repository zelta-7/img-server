package repository

import "github.com/jinzhu/gorm"

type Repo interface {
	PostImg(newImg ImageRecord) error
	GetImg(url string) (string, error)
}

type imgRepo struct {
	db *gorm.DB
}

func NewImgRepo(db *gorm.DB) Repo {
	return &imgRepo{db: db}
}

func (r *imgRepo) GetImg(url string) (string, error) {
	var images ImageResponse
	if err := r.db.Where("url = ?", url).Find(&images).Error; err != nil {
		return "", err
	}
	return images.CompressedPath, nil
}

func (r *imgRepo) PostImg(newImg ImageRecord) error {
	return r.db.Create(&newImg).Error
}
