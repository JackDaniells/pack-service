package entity

type Pack struct {
	Size     int `json:"size"`
	Quantity int `json:"quantity,omitempty"`
}
