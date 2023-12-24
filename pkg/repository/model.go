package repository

import "github.com/jinzhu/gorm"

type ImageRecord struct {
	gorm.Model
	Id             uint   `gorm:"unique; not null"`
	Url            string `gorm:"type varchar(100); not null"`
	CompressedPath string `gorm:"type varchar(100); not null"`
}

type ImageResponse struct {
	Url            string `json:"url"`
	CompressedPath string `json:"compressedpath"`
}
