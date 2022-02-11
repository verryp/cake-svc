package payload

type CakeCreateRequest struct {
	Id          string    `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Rating      int       `json:"rating" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	// CreatedAt   time.Time `json:"-"`
	// UpdatedAt   time.Time `json:"-"`
}