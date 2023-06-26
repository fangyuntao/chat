package database

import (
	"context"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/db/tx"
	"github.com/OpenIMSDK/chat/pkg/common/db/model/organization"
	table "github.com/OpenIMSDK/chat/pkg/common/db/table/organization"
	"gorm.io/gorm"
)

type OrganizationDatabaseInterface interface {
	//department
	GetDepartmentByID(ctx context.Context, departmentID string) (*table.Department, error)
	CreateDepartment(ctx context.Context, department ...*table.Department) error
	UpdateDepartment(ctx context.Context, department *table.Department) error
	GetParent(ctx context.Context, parentID string) ([]*table.Department, error)
	FindDepartmentMember(ctx context.Context, list []string) ([]*table.DepartmentMember, error)
	GetDepartment(ctx context.Context, departmentID string) ([]*table.DepartmentMember, error)
	GetList(ctx context.Context, departmentIDList []string) ([]*table.Department, error)
	DeleteDepartment(ctx context.Context, departmentIDList []string) error
	UpdateParentID(ctx context.Context, oldParentID, newParentID string) error
	//departmentMember
	FindDepartmentMember(ctx context.Context, list []string) ([]*table.DepartmentMember, error)
	GetDepartmentMember(ctx context.Context, userID string) ([]*table.DepartmentMember, error)
	CreateDepartmentMember(ctx context.Context, DepartmentMember *table.DepartmentMember) error
	DeleteDepartmentIDList(ctx context.Context, departmentIDList []string) error
	DeleteDepartmentMemberByUserID(ctx context.Context, userID string) error
	DeleteDepartmentMemberByKey(ctx context.Context, userID string, departmentID string) error
	//organizationUser
	CreateOrganizationUser(ctx context.Context, OrganizationUser *table.OrganizationUser) error
	UpdateOrganizationUser(ctx context.Context, OrganizationUser *table.OrganizationUser) error
	DeleteOrganizationUser(ctx context.Context, userID string) error
	GetOrganizationUser(ctx context.Context, userID string) (*table.OrganizationUser, error)
}

func NewOrganizationDatabase(db *gorm.DB) OrganizationDatabaseInterface {
	return &OrganizationDatabase{
		tx:         tx.NewGorm(db),
		Department: organization.NewDepartment(db),
	}
}

type OrganizationDatabase struct {
	tx               tx.Tx
	Department       table.DepartmentInterface
	DepartmentMember table.DepartmentMemberInterface
	OrganizationUser table.OrganizationUserInterface
}

func (o *OrganizationDatabase) DeleteDepartmentMember(ctx context.Context, userID string) error {
	return o.OrganizationUser
}

func (o *OrganizationDatabase) DeleteOrganizationUser(ctx context.Context, userID string) error {
	return o.OrganizationUser.Delete(ctx, userID)
}

func (o *OrganizationDatabase) UpdateOrganizationUser(ctx context.Context, OrganizationUser *table.OrganizationUser) error {
	return o.OrganizationUser.Update(ctx, OrganizationUser)
}

func (o *OrganizationDatabase) CreateOrganizationUser(ctx context.Context, OrganizationUser *table.OrganizationUser) error {
	return o.OrganizationUser.Create(ctx, OrganizationUser)
}

func (o *OrganizationDatabase) DeleteDepartmentIDList(ctx context.Context, departmentIDList []string) error {
	return o.DepartmentMember.DeleteDepartmentIDList(ctx, departmentIDList)
}

func (o *OrganizationDatabase) DeleteDepartment(ctx context.Context, departmentIDList []string) error {
	return o.Department.Delete(ctx, departmentIDList)
}

func (o *OrganizationDatabase) UpdateParentID(ctx context.Context, oldParentID, newParentID string) error {
	return o.Department.UpdateParentID(ctx, oldParentID, newParentID)
}

func (o *OrganizationDatabase) GetList(ctx context.Context, departmentIDList []string) ([]*table.Department, error) {
	return o.Department.GetList(ctx, departmentIDList)
}

func (o *OrganizationDatabase) FindDepartmentMember(ctx context.Context, departmentIDList []string) ([]*table.DepartmentMember, error) {
	return o.DepartmentMember.Find(ctx, departmentIDList)
}

func (o *OrganizationDatabase) GetParent(ctx context.Context, parentID string) ([]*table.Department, error) {
	return o.Department.GetParent(ctx, parentID)
}

func (o *OrganizationDatabase) UpdateDepartment(ctx context.Context, department *table.Department) error {
	return o.Department.Update(ctx, department)
}

func (o *OrganizationDatabase) GetDepartmentByID(ctx context.Context, departmentID string) (*table.Department, error) {
	return o.Department.FindOne(ctx, departmentID)
}

func (o *OrganizationDatabase) CreateDepartment(ctx context.Context, department ...*table.Department) error {
	return o.Department.Create(ctx, department...)
}

func (o *OrganizationDatabase) GetDepartment(ctx context.Context, departmentID string) ([]*table.DepartmentMember, error) {
	return o.DepartmentMember.GetDepartment(ctx, departmentID)
}