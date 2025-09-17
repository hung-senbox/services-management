package handler

import (
	"net/http"
	"services-management/helper"
	"services-management/internal/sv_management/dto/request"
	service "services-management/internal/sv_management/services"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	service service.SvManagementService
}

func NewServiceHandler(service service.SvManagementService) *ServiceHandler {
	return &ServiceHandler{
		service: service,
	}
}

func (s *ServiceHandler) Upload(c *gin.Context) {
	var req request.UploadServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	err := s.service.UploadService(c.Request.Context(), req)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInternal)
		return
	}
	helper.SendSuccess(c, http.StatusOK, "Upload service successfully", nil)
}
