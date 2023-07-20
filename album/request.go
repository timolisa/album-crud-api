package album


type AlbumRequest struct {
	Title string `json:"album_title" binding:"required"`
	Artist string `json:"album_artist" binding:"required"`
	Price string `json:"album_price" binding:"required"`	
}