package es

// MyDocument
type MyDocdument struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Category  string `json:"category"`
}

// SearchCriteria
type SearchCriteria struct {
	Category  string `json:"category"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Title     string `json:"title"`
	UpdatedAt int64  `json:"updated_at"`
	Order     string `json:"order"`
	Sort      string `json:"sort"`
}
