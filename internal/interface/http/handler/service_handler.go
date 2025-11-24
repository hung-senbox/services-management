package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senbox/services-management/internal/domain/usecase"
	"github.com/senbox/services-management/internal/interface/http/dto"
	"github.com/senbox/services-management/internal/interface/http/mapper"
)

type ServiceHandler struct {
	serviceUseCase usecase.ServiceUseCase
}

func NewServiceHandler(serviceUseCase usecase.ServiceUseCase) *ServiceHandler {
	return &ServiceHandler{
		serviceUseCase: serviceUseCase,
	}
}

func (h *ServiceHandler) CreateService(c *fiber.Ctx) error {
	var req dto.CreateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	service, err := mapper.ToServiceEntity(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid service group ID",
		})
	}

	if err := h.serviceUseCase.CreateService(c.Context(), service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create service",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Service created successfully",
		"data":    mapper.ToServiceResponse(service),
	})
}

func (h *ServiceHandler) GetServiceByID(c *fiber.Ctx) error {
	id := c.Params("id")

	service, err := h.serviceUseCase.GetServiceByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	if service == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Service not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": mapper.ToServiceResponse(service),
	})
}

func (h *ServiceHandler) GetAllServices(c *fiber.Ctx) error {
	services, err := h.serviceUseCase.GetAllServices(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch services",
		})
	}

	var response []*dto.ServiceResponse
	for _, s := range services {
		response = append(response, mapper.ToServiceResponse(s))
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}

func (h *ServiceHandler) GetServicesByGroupID(c *fiber.Ctx) error {
	serviceGroupID := c.Params("groupId")

	services, err := h.serviceUseCase.GetServicesByGroupID(c.Context(), serviceGroupID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid service group ID",
		})
	}

	var response []*dto.ServiceResponse
	for _, s := range services {
		response = append(response, mapper.ToServiceResponse(s))
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}

func (h *ServiceHandler) UpdateService(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	service, err := mapper.ToServiceEntityFromUpdate(id, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	if err := h.serviceUseCase.UpdateService(c.Context(), service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update service",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Service updated successfully",
		"data":    mapper.ToServiceResponse(service),
	})
}

func (h *ServiceHandler) DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.serviceUseCase.DeleteService(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete service",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Service deleted successfully",
	})
}
