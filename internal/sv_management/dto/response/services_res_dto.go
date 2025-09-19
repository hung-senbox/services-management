package response

type ServicesResponse struct {
	Group    ServiceGroupResponse `json:"group"`
	Services []ServiceResDto      `json:"services"`
}

type ServiceGroupResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Order int    `json:"order"`
}
