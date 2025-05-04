package types

type EventResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Capacity    int     `json:"capacity"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}

type PaginatedEventsResponse struct {
	Events      []EventResponse `json:"events"`
	TotalItems  int64           `json:"total_items"`
	TotalPages  int             `json:"total_pages"`
	CurrentPage int             `json:"current_page"`
	PageSize    int             `json:"page_size"`
}
