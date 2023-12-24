package service

import "github.com/zelta-7/img-server/pkg/repository"

type ImgService interface {
	GetImg(url string) (string, error)
	PostImg(repository.ImageRecord) error
}

type imgService struct {
	imgRepository repository.Repo
}

func (s *imgService) GetImg(url string) (string, error) {
	return s.imgRepository.GetImg(url)
}

func (s *imgService) PostImg(img repository.ImageRecord) error {
	return s.imgRepository.PostImg(img)
}

func NewImgService(imgRepository repository.Repo) ImgService {
	return &imgService{imgRepository: imgRepository}
}
