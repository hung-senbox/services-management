package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senbox/services-management/internal/domain/usecase"
	"github.com/senbox/services-management/internal/interface/http/dto"
	"github.com/senbox/services-management/internal/interface/http/mapper"
)

type ServiceGroupHandler struct {
	serviceGroupUseCase usecase.ServiceGroupUseCase
}

func NewServiceGroupHandler(serviceGroupUseCase usecase.ServiceGroupUseCase) *ServiceGroupHandler {
	return &ServiceGroupHandler{
		serviceGroupUseCase: serviceGroupUseCase,
	}
}

func (h *ServiceGroupHandler) CreateServiceGroup(c *fiber.Ctx) error {
	var req dto.CreateServiceGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	serviceGroup := mapper.ToServiceGroupEntity(&req)
	if err := h.serviceGroupUseCase.CreateServiceGroup(c.Context(), serviceGroup); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create service group",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Service group created successfully",
		"data":    mapper.ToServiceGroupResponse(serviceGroup),
	})
}

func (h *ServiceGroupHandler) GetServiceGroupByID(c *fiber.Ctx) error {
	id := c.Params("id")

	serviceGroup, err := h.serviceGroupUseCase.GetServiceGroupByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	if serviceGroup == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Service group not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": mapper.ToServiceGroupResponse(serviceGroup),
	})
}

func (h *ServiceGroupHandler) GetAllServiceGroups(c *fiber.Ctx) error {
	serviceGroups, err := h.serviceGroupUseCase.GetAllServiceGroups(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch service groups",
		})
	}

	var response []*dto.ServiceGroupResponse
	for _, sg := range serviceGroups {
		response = append(response, mapper.ToServiceGroupResponse(sg))
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}

func (h *ServiceGroupHandler) UpdateServiceGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateServiceGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	serviceGroup, err := mapper.ToServiceGroupEntityFromUpdate(id, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	if err := h.serviceGroupUseCase.UpdateServiceGroup(c.Context(), serviceGroup); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update service group",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Service group updated successfully",
		"data":    mapper.ToServiceGroupResponse(serviceGroup),
	})
}

func (h *ServiceGroupHandler) DeleteServiceGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.serviceGroupUseCase.DeleteServiceGroup(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete service group",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Service group deleted successfully",
	})
}
