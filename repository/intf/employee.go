package intf

import (
	"context"

	"github.com/employee-management/model"
)

type Employee interface {
	GetByID(ctx context.Context, id string) (*model.Employee, error)
	Upsert(ctx context.Context, emp *model.Employee) error
	Delete(ctx context.Context, id string) error
	GetList(ctx context.Context, limit, offset int) ([]*model.Employee, error)
}
