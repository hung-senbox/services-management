package dto

type CreateServiceGroupRequest struct {
	Name        string `json:"name" validate:"required"`
	Order       int    `json:"order" validate:"required,gte=0"`
	IsActive    bool   `json:"is_active"`
	Description string `json:"description"`
	IconKey     string `json:"icon_key"`
}

type UpdateServiceGroupRequest struct {
	Name        string `json:"name" validate:"required"`
	Order       int    `json:"order" validate:"required,gte=0"`
	IsActive    bool   `json:"is_active"`
	Description string `json:"description"`
	IconKey     string `json:"icon_key"`
}

type ServiceGroupResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	IsActive    bool   `json:"is_active"`
	Description string `json:"description"`
	IconKey     string `json:"icon_key"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

