package usecase

import (
	"context"
	"time"

	"services-management/internal/domain/entity"
	"services-management/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceGroupUseCase interface {
	CreateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error
	GetServiceGroupByID(ctx context.Context, id string) (*entity.ServiceGroup, error)
	GetAllServiceGroups(ctx context.Context) ([]*entity.ServiceGroup, error)
	UpdateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error
	DeleteServiceGroup(ctx context.Context, id string) error
	MigrateServiceGroup(ctx context.Context) error
}

type serviceGroupUseCase struct {
	serviceGroupRepo repository.ServiceGroupRepository
}

func NewServiceGroupUseCase(serviceGroupRepo repository.ServiceGroupRepository) ServiceGroupUseCase {
	return &serviceGroupUseCase{
		serviceGroupRepo: serviceGroupRepo,
	}
}

func (uc *serviceGroupUseCase) CreateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error {
	return uc.serviceGroupRepo.Create(ctx, serviceGroup)
}

func (uc *serviceGroupUseCase) GetServiceGroupByID(ctx context.Context, id string) (*entity.ServiceGroup, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return uc.serviceGroupRepo.FindByID(ctx, objectID)
}

func (uc *serviceGroupUseCase) GetAllServiceGroups(ctx context.Context) ([]*entity.ServiceGroup, error) {
	return uc.serviceGroupRepo.FindAll(ctx)
}

func (uc *serviceGroupUseCase) UpdateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error {
	return uc.serviceGroupRepo.Update(ctx, serviceGroup)
}

func (uc *serviceGroupUseCase) DeleteServiceGroup(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return uc.serviceGroupRepo.Delete(ctx, objectID)
}

// migrate service group
//
// service group 1: Organization Dashboard, order 1, icon : https://senbox.vn/images/icon/organization-dashboard.png, url : /management/dashboard
// service group 2: Users Devices, order 2, icon : https://senbox.vn/images/icon/organization-management.png, url : /management/users/default-view?name=&role=all&status=active
// service group 3: Department Arrangement, order 3, icon : https://senbox.vn/images/icon/device-management.png, url : /management/department
// service group 4: Inbox, order 4, icon : https://senbox.vn/images/icon/device-management.png, url : /management/inbox/default-view?name=&role=all&status=active
// service group 5: Chat(Notification), order 5, icon : https://senbox.vn/images/icon/device-management.png, url : /management/chat
// service group 6: Emergencies, order 6, icon : https://senbox.vn/images/icon/device-management.png, url : /management/emergencies
// service group 7: Classroom Arrangement, order 7, icon : https://senbox.vn/images/icon/device-management.png, url : /management/classes-arrangement
// service group 8: Staff Management, order 8, icon : https://senbox.vn/images/icon/device-management.png, url : /management/staff-management
// service group 9: Colored Timetables, order 9, icon : https://senbox.vn/images/icon/device-management.png, url : /management/time-tables
// service group 10: Food Health, order 10, icon : https://senbox.vn/images/icon/device-management.png, url : /management/food-health
// service group 11: Events(Notification), order 11, icon : https://senbox.vn/images/icon/device-management.png, url : /management/events
// service group 12: Data-Graphs(Behaviors), order 12, icon : https://senbox.vn/images/icon/device-management.png, url : /management/charts
// service group 13: OB/IEP Form-Builder, order 13, icon : https://senbox.vn/images/icon/device-management.png, url : /management/ob-iep
// service group 14: Reports, order 14, icon : https://senbox.vn/images/icon/device-management.png, url : /management/reports
// service group 15: PD Certificates, order 15, icon : https://senbox.vn/images/icon/device-management.png, url : /management/feedbacks
// service group 16: To-Dp(Schedules), order 16, icon : https://senbox.vn/images/icon/device-management.png, url : /management/todo
// service group 17: File Manager, order 17, icon : https://senbox.vn/images/icon/device-management.png, url : /management/files
// service group 18: Location Inventory, order 18, icon : https://senbox.vn/images/icon/device-management.png, url : /management/inventories
// service group 19: QR Print, order 19, icon : https://senbox.vn/images/icon/device-management.png, url : /management/qr
// service group 20: Media, order 20, icon : https://senbox.vn/images/icon/device-management.png, url : /management/media
// service group 21: Product Editor, order 21, icon : https://senbox.vn/images/icon/device-management.png, url : /management/product
// service group 22: Store Manager, order 22, icon : https://senbox.vn/images/icon/device-management.png, url : /management/stores
// service group 23: Task Manager, order 23, icon : https://senbox.vn/images/icon/device-management.png, url : /management/tasks
// service group 24: Handshake, order 24, icon : https://senbox.vn/images/icon/device-management.png, url : /management/handshake

func (uc *serviceGroupUseCase) MigrateServiceGroup(ctx context.Context) error {
	serviceGroups := []*entity.ServiceGroup{
		{
			Name:        "Organization Dashboard",
			Order:       1,
			Icon:        "",
			Url:         "/management/dashboard",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Users Devices",
			Order:       2,
			Icon:        "",
			Url:         "/management/users",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Department Arrangement",
			Order:       3,
			Icon:        "",
			Url:         "/management/department",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Inbox",
			Order:       4,
			Icon:        "",
			Url:         "/management/inbox",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Chat(Notification)",
			Order:       5,
			Icon:        "",
			Url:         "/management/chat",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Emergencies",
			Order:       6,
			Icon:        "",
			Url:         "/management/emergencies",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Classroom Arrangement",
			Order:       7,
			Icon:        "",
			Url:         "/management/classes-arrangement",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Staff Management",
			Order:       8,
			Icon:        "",
			Url:         "/management/staff-management",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Colored Timetables",
			Order:       9,
			Icon:        "",
			Url:         "/management/time-tables",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Food Health",
			Order:       10,
			Icon:        "",
			Url:         "/management/food-health",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Events(Notification)",
			Order:       11,
			Icon:        "",
			Url:         "/management/events",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Data-Graphs(Behaviors)",
			Order:       12,
			Icon:        "",
			Url:         "/management/charts",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "OB/IEP Form-Builder",
			Order:       13,
			Icon:        "",
			Url:         "/management/ob-iep",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Reports",
			Order:       14,
			Icon:        "",
			Url:         "/management/reports",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "PD Certificates",
			Order:       15,
			Icon:        "",
			Url:         "/management/feedbacks",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "To-Dp(Schedules)",
			Order:       16,
			Icon:        "",
			Url:         "/management/todo",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "File Manager",
			Order:       17,
			Icon:        "",
			Url:         "/management/files",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Location Inventory",
			Order:       18,
			Icon:        "",
			Url:         "/management/inventories",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "QR Print",
			Order:       19,
			Icon:        "",
			Url:         "/management/qr",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Media",
			Order:       20,
			Icon:        "",
			Url:         "/management/media",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Product Editor",
			Order:       21,
			Icon:        "",
			Url:         "/management/product",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Store Manager",
			Order:       22,
			Icon:        "",
			Url:         "/management/stores",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Task Manager",
			Order:       23,
			Icon:        "",
			Url:         "/management/tasks",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Handshake",
			Order:       24,
			Icon:        "",
			Url:         "/management/handshake",
			IsActive:    true,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, sg := range serviceGroups {
		if err := uc.serviceGroupRepo.Create(ctx, sg); err != nil {
			return err
		}
	}

	return nil
}
