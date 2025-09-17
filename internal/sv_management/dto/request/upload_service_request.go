package request

type UploadServiceRequest struct {
	Title   string `json:"service_name" binding:"required"`
	Url     string `json:"url" binding:"required"`
	Order   int    `json:"order" binding:"required"`
	GroupID string `json:"group_id" binding:"required"`
}
