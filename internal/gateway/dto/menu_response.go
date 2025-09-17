package dto

type MenuResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Order  int    `json:"order"`
	IsShow bool   `json:"is_show"`
}
