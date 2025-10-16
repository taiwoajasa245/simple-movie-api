package models

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director,omitempty"`
	Year     int    `json:"year,omitempty"`
	Genre    string `json:"genre,omitempty"`
}

type APIResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// type Movies struct {
// 	Movies interface{} `json:"movies"`
// }
