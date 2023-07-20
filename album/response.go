package album

type AlbumResponse struct {
	ID string `json:"album_id"`
	Title string `json:"album_title"`
	Artist string `json:"album_artist"`
	Price float64 `json:"album_price"`	
}