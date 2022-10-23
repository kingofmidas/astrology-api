package dto

type ApodResponse struct {
	Title     string `json:"title"`
	Date      string `json:"date"`
	URL       string `json:"url"`
	MediaType string `json:"media_type"`
}

type ApodQueryParams struct {
	Date string
}
