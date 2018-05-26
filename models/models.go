package models

// User used to generate token payload
type User struct {
	Name     string
	Location string
	Sub      string
	Admin    bool
}

// Events ...
// swagger:model
type Events struct {
	Events []Event `json:"events"`
}

// Event ...
// swagger:model
type Event struct {
	Title          string `json:"Title"`
	Date           string `json:"date"`
	ImageURL       string `json:"image"`
	AvailableSeats int    `json:"availableSeats"`
	Location       string `json:"location"`
}
