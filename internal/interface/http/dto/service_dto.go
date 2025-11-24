package dto

type CreateServiceRequest struct {
	ServiceGroupID string `json:"service_group_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	IsActive       bool   `json:"is_active"`
	Description    string `json:"description"`
}

type UpdateServiceRequest struct {
	ServiceGroupID string `json:"service_group_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	IsActive       bool   `json:"is_active"`
	Description    string `json:"description"`
}

type ServiceResponse struct {
	ID             string `json:"id"`
	ServiceGroupID string `json:"service_group_id"`
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
