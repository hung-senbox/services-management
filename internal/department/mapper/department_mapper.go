package mapper

import (
	"department-service/internal/department/dto/response"
	"department-service/internal/department/model"
	"department-service/internal/gateway/dto"
)

func MapDepartmentToResponse(dept *model.Department, menus []dto.MenuResponse, iconUrl string) *response.DepartmentResponseDTO {
	staffs := dept.Staffs
	if staffs == nil {
		staffs = []model.Staff{}
	}
	if menus == nil {
		menus = []dto.MenuResponse{}
	}
	return &response.DepartmentResponseDTO{
		ID:             dept.ID.Hex(),
		LocationID:     dept.LocationID,
		OrganizationID: dept.OrganizationID,
		RegionID:       dept.RegionID,
		Name:           dept.Name,
		Description:    dept.Description,
		Message:        dept.Message,
		Icon:           dept.Icon,
		IconUrl:        iconUrl,
		Leader:         dept.Leader,
		Staffs:         staffs,
		Menus:          menus,
	}
}

func MapDepartmentsToResponses(depts []*model.Department) []*response.DepartmentResponseDTO {
	responses := make([]*response.DepartmentResponseDTO, len(depts))
	for i, dept := range depts {
		responses[i] = MapDepartmentToResponse(dept, nil, "")
	}
	return responses
}

func MapDepartmentToResponse4Web(dept *model.Department, homeMenus []dto.MenuResponse, iconUrl string, leader response.LeaderResponseDTO, staffs []response.StaffResponseDTO, organizationMenus []dto.MenuResponse) *response.GetDepartment4Web {

	if homeMenus == nil {
		homeMenus = []dto.MenuResponse{}
	}
	if organizationMenus == nil {
		organizationMenus = []dto.MenuResponse{}
	}
	if staffs == nil {
		staffs = []response.StaffResponseDTO{}
	}
	return &response.GetDepartment4Web{
		ID:                dept.ID.Hex(),
		LocationID:        dept.LocationID,
		OrganizationID:    dept.OrganizationID,
		RegionID:          dept.RegionID,
		Name:              dept.Name,
		Description:       dept.Description,
		Message:           dept.Message,
		Icon:              dept.Icon,
		IconUrl:           iconUrl,
		Leader:            leader,
		Staffs:            staffs,
		HomeMenus:         homeMenus,
		OrganizationMenus: organizationMenus,
	}
}

func MapDepartmentsToGroupedResponses4Web(
	depts []*model.Department,
	homeMenusMap map[string][]dto.MenuResponse,
	iconUrls map[string]string,
	leaders map[string]response.LeaderResponseDTO,
	staffsMap map[string][]response.StaffResponseDTO,
	organizationMenusMap map[string][]dto.MenuResponse,
) []*response.DepartmentGroupResponse {
	groupMap := make(map[string][]*response.GetDepartment4Web)

	for _, dept := range depts {
		homeMenus := homeMenusMap[dept.ID.Hex()]
		iconUrl := iconUrls[dept.ID.Hex()]
		leader := leaders[dept.ID.Hex()]
		staffs := staffsMap[dept.ID.Hex()]
		organizationMenusMap := organizationMenusMap[dept.ID.Hex()]

		resp := MapDepartmentToResponse4Web(dept, homeMenus, iconUrl, leader, staffs, organizationMenusMap)
		groupMap[dept.RegionID] = append(groupMap[dept.RegionID], resp)
	}

	// convert map -> slice
	var result []*response.DepartmentGroupResponse
	for regjonIdx, depts := range groupMap {
		result = append(result, &response.DepartmentGroupResponse{
			RegionID:    regjonIdx,
			Departments: depts,
		})
	}

	return result
}

func MapDepartmentToResponse4App(dept *model.Department) *response.GetDepartment4App {
	return &response.GetDepartment4App{
		ID:             dept.ID.Hex(),
		LocationID:     dept.LocationID,
		OrganizationID: dept.OrganizationID,
		Name:           dept.Name,
		Description:    dept.Description,
		Message:        dept.Message,
		Icon:           dept.Icon,
	}
}

func MapDepartmentsToResponses4App(depts []*model.Department) []*response.GetDepartment4App {
	responses := make([]*response.GetDepartment4App, len(depts))
	for i, dept := range depts {
		responses[i] = MapDepartmentToResponse4App(dept)
	}
	return responses
}

func MapDepartmentToResponse4Gateway(dept *model.Department) *response.GetDepartment4Gateway {
	return &response.GetDepartment4Gateway{
		ID:   dept.ID.Hex(),
		Icon: dept.Icon,
		Name: dept.Name,
	}
}

func MapDepartmentsToResponses4Gateway(depts []*model.Department) []*response.GetDepartment4Gateway {
	responses := make([]*response.GetDepartment4Gateway, len(depts))
	for i, dept := range depts {
		responses[i] = MapDepartmentToResponse4Gateway(dept)
	}
	return responses
}

func MapDepartmentsToResponses4Web(
	depts []*model.Department,
	homeMenusMap map[string][]dto.MenuResponse,
	iconUrls map[string]string,
	leaders map[string]response.LeaderResponseDTO,
	staffsMap map[string][]response.StaffResponseDTO,
	organizationMenusMap map[string][]dto.MenuResponse,
) []*response.GetDepartment4Web {
	result := make([]*response.GetDepartment4Web, 0, len(depts))
	for _, dept := range depts {
		homeMenus := homeMenusMap[dept.ID.Hex()]
		iconUrl := iconUrls[dept.ID.Hex()]
		leader := leaders[dept.ID.Hex()]
		staffs := staffsMap[dept.ID.Hex()]
		orgMenus := organizationMenusMap[dept.ID.Hex()]

		resp := MapDepartmentToResponse4Web(dept, homeMenus, iconUrl, leader, staffs, orgMenus)
		result = append(result, resp)
	}
	return result
}
