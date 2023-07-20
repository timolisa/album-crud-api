package album

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Album, error)
	FindByID(ID int) (Album, error)
	Create(album Album) (Album, error)
	Update(album Album) (Album, error)
	Delete(ID int) (error)
}

type repository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Album, error) {
	var albums []Album
	err := r.db.Find(&albums).Error

	return albums, err
}

func (r *repository) FindByID(ID int) (Album, error) {
	var album Album
	err := r.db.Find(&album, ID).Error

	return album, err
}

func (r *repository) Create(album Album) (Album, error) {
	err := r.db.Create(&album).Error

	return album, err
}

func (r *repository) Update(album Album) (Album, error) {
	err := r.db.Save(&album).Error

	return album, err
}

func (r *repository) Delete(ID int) (error) {
	err := r.db.Delete(&ID).Error
	return err
}