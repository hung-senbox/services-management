package request

type UploadDepartmentRequest struct {
	LocationID  string `json:"location_id"`
	ComponentID string `json:"component_id"`
	RegionID    string `json:"region_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Message     string `json:"message"`
	Icon        string `json:"icon" binding:"required"`
}
