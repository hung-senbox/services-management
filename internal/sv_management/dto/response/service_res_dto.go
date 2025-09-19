package response

type ServiceResDto struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Order int    `json:"order"`
	Url   string `json:"url"`
}
