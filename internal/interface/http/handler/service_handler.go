package handler

import (
	"errors"
	"services-management/internal/domain/usecase"
	"services-management/internal/interface/http/dto"
	"services-management/internal/interface/http/mapper"

	libs_helper "services-management/pkg/libs/helper"

	"github.com/gofiber/fiber/v2"
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
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	service, err := mapper.ToServiceEntity(&req)
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	if err := h.serviceUseCase.CreateService(c.Context(), service); err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	return libs_helper.SendSuccess(c, fiber.StatusCreated, "Service created successfully", mapper.ToServiceResponse(service))
}

func (h *ServiceHandler) GetServiceByID(c *fiber.Ctx) error {
	id := c.Params("id")

	service, err := h.serviceUseCase.GetServiceByID(c.Context(), id)
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	if service == nil {
		return libs_helper.SendError(c, fiber.StatusNotFound, errors.New("service not found"), libs_helper.ErrNotFound)
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Service fetched successfully", mapper.ToServiceResponse(service))
}

func (h *ServiceHandler) GetAllServices(c *fiber.Ctx) error {
	services, err := h.serviceUseCase.GetAllServices(c.Context())
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	var response []*dto.ServiceResponse
	for _, s := range services {
		response = append(response, mapper.ToServiceResponse(s))
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Services fetched successfully", response)
}

func (h *ServiceHandler) GetServicesByGroupID(c *fiber.Ctx) error {
	serviceGroupID := c.Params("groupId")

	services, err := h.serviceUseCase.GetServicesByGroupID(c.Context(), serviceGroupID)
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	var response []*dto.ServiceResponse
	for _, s := range services {
		response = append(response, mapper.ToServiceResponse(s))
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Services fetched successfully", response)
}

func (h *ServiceHandler) UpdateService(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	service, err := mapper.ToServiceEntityFromUpdate(id, &req)
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	if err := h.serviceUseCase.UpdateService(c.Context(), service); err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Service updated successfully", mapper.ToServiceResponse(service))
}

func (h *ServiceHandler) DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.serviceUseCase.DeleteService(c.Context(), id); err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Service deleted successfully", nil)
}
