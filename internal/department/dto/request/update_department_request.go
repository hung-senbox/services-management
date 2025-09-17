package request

type UpdateDepartmentRequest struct {
	ID          string `json:"id" binding:"required"`
	LocationID  string `json:"location_id"`
	ComponentID string `json:"component_id"`
	RegionID    string `json:"region_id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Message     string `json:"message"`
	Icon        string `json:"icon"`
}
