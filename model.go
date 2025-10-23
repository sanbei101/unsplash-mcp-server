package main

type SearchPhotosRequest struct {
	Query       string `json:"query"`
	Page        int    `json:"page"`
	PerPage     int    `json:"per_page"`
	OrderBy     string `json:"order_by"`
	Color       string `json:"color"`
	Orientation string `json:"orientation"`
}

// SearchResponse represents the Unsplash search API response
type SearchResponse struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Photo `json:"results"`
}

type SearchResult struct {
	Photos []Photo `json:"photos"`
}

// Photo represents an Unsplash photo
type Photo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Urls        URLs   `json:"urls"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

// URLs contains various size URLs for a photo
type URLs struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}
