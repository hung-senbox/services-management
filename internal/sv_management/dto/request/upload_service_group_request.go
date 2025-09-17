package request

type UploadServiceGroupRequest struct {
	Title string `json:"title" binding:"required"`
	Order int    `json:"order" binding:"required"`
}
