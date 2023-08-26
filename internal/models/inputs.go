package models

type Inputs struct {
	ID         int64  `json:"id,omitempty"`
	Name       string `json:"name"`
	TypeInputs int64  `json:"type_inputs"`
}
