package models

type ServiceInputs struct {
	ID        int64    `json:"id,omitempty"`
	ServiceID int64    `json:"service_id"`
	Required  bool     `json:"required"`
	Services  Services `json:"services"`
	InputIDs  []int64  `json:"Input_ids"`
	Inputs    []Inputs `json:"inputs"`
}
