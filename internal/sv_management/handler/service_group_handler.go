package handler

import (
	"net/http"
	"services-management/helper"
	"services-management/internal/sv_management/dto/request"
	service "services-management/internal/sv_management/services"

	"github.com/gin-gonic/gin"
)

type ServiceGroupHandler struct {
	service service.SVGroupService
}

func NewServiceGroupHandler(service service.SVGroupService) *ServiceGroupHandler {
	return &ServiceGroupHandler{
		service: service,
	}
}

func (s *ServiceGroupHandler) Upload(c *gin.Context) {
	var req request.UploadServiceGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	err := s.service.UploadServiceGroup(c.Request.Context(), req)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInternal)
		return
	}
	helper.SendSuccess(c, http.StatusOK, "Upload service group successfully", nil)
}
