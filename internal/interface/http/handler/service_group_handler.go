package handler

import (
	"errors"
	"services-management/internal/domain/usecase"
	"services-management/internal/interface/http/dto"
	"services-management/internal/interface/http/mapper"
	libs_helper "services-management/pkg/libs/helper"

	"github.com/gofiber/fiber/v2"
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
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	serviceGroup := mapper.ToServiceGroupEntity(&req)
	if err := h.serviceGroupUseCase.CreateServiceGroup(c.Context(), serviceGroup); err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	return libs_helper.SendSuccess(c, fiber.StatusCreated, "Service group created successfully", mapper.ToServiceGroupResponse(serviceGroup))
}

func (h *ServiceGroupHandler) GetServiceGroupByID(c *fiber.Ctx) error {
	id := c.Params("id")

	serviceGroup, err := h.serviceGroupUseCase.GetServiceGroupByID(c.Context(), id)
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	if serviceGroup == nil {
		return libs_helper.SendError(c, fiber.StatusNotFound, errors.New("service group not found"), libs_helper.ErrNotFound)
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Service group fetched successfully", mapper.ToServiceGroupResponse(serviceGroup))
}

func (h *ServiceGroupHandler) GetAllServiceGroups(c *fiber.Ctx) error {
	serviceGroups, err := h.serviceGroupUseCase.GetAllServiceGroups(c.Context())
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	var response []*dto.ServiceGroupResponse
	for _, sg := range serviceGroups {
		response = append(response, mapper.ToServiceGroupResponse(sg))
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Service groups fetched successfully", response)
}

func (h *ServiceGroupHandler) UpdateServiceGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateServiceGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	serviceGroup, err := mapper.ToServiceGroupEntityFromUpdate(id, &req)
	if err != nil {
		return libs_helper.SendError(c, fiber.StatusBadRequest, err, libs_helper.ErrInvalidRequest)
	}

	if err := h.serviceGroupUseCase.UpdateServiceGroup(c.Context(), serviceGroup); err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Service group updated successfully", mapper.ToServiceGroupResponse(serviceGroup))
}

func (h *ServiceGroupHandler) DeleteServiceGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.serviceGroupUseCase.DeleteServiceGroup(c.Context(), id); err != nil {
		return libs_helper.SendError(c, fiber.StatusInternalServerError, err, libs_helper.ErrInternal)
	}

	return libs_helper.SendSuccess(c, fiber.StatusOK, "Service group deleted successfully", nil)
}
