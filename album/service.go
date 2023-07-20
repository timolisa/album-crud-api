package album

import (
	"errors"
	"strconv"
)

type Service interface {
	FindAll() ([]AlbumResponse, error)
	FindByID(ID string) (AlbumResponse, error)
	Create(req AlbumRequest) (AlbumResponse, error)
	Update(albumID string, req AlbumRequest) (AlbumResponse, error)
	Delete(ID string) (string)
}

type service struct {
	repository Repository
}

func NewAlbumService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByID(ID string) (AlbumResponse, error) {
	albumID, _ := strconv.Atoi(ID)
	album, err := s.repository.FindByID(albumID)
	if err != nil {
		return AlbumResponse{}, err
	}
	return album.convertAlbumToResponse(), err
}

func (s *service) FindAll() ([]AlbumResponse, error) {
	albums, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var albumsResponse []AlbumResponse

	for _, album := range albums {
		response := album.convertAlbumToResponse()

		albumsResponse = append(albumsResponse, response)
	}

	return albumsResponse, err
}

func (s *service) Create(req AlbumRequest) (AlbumResponse, error) {
	price, err := strconv.ParseFloat(req.Price, 64)

	if err != nil {
		return AlbumResponse{}, errors.New("invalid price format")
	}

	if price <= 0 {
		return AlbumResponse{}, errors.New("price cannot be negative")
	}

	album := Album{  
		Title: req.Title,
		Artist: req.Artist,
		Price: price,
	}

	createdAlbum, err := s.repository.Create(album)
	if err != nil {
		return AlbumResponse{}, err
	}

	response := createdAlbum.convertAlbumToResponse()

	return response, nil
}

func (s *service) Update(albumID string, req AlbumRequest) (AlbumResponse, error) {
	price, err := strconv.ParseFloat(req.Price, 64)
	id, _ := strconv.Atoi(albumID)
	if err != nil {
		return AlbumResponse{}, errors.New("invalid price format")
	}

	if price <= 0 {
		return AlbumResponse{}, errors.New("price cannot be negative")
	}

	album := Album{
		ID: id,
		Title: req.Title,
		Artist: req.Artist,
		Price: price,
	}

	updatedAlbum, err := s.repository.Update(album)

	if err != nil {
		return AlbumResponse{}, err
	}

	response := updatedAlbum.convertAlbumToResponse()

	return response, nil
}

func (s *service) Delete(ID string) (string) {
	albumID, _ := strconv.Atoi(ID)
	err := s.repository.Delete(albumID)

	if err != nil {
		return "error: failed to delete album"
	}
	return "Album deleted successfully"
}

func (album Album) convertAlbumToResponse() (AlbumResponse) {
	response := AlbumResponse{
		ID: strconv.Itoa(album.ID),
		Title: album.Title,
		Artist: album.Artist,
		Price: album.Price,
	}
	return response
}