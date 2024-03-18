package repository

import (
	"context"
	"goland-crud/model"
)

type GrowBoxRepository interface {
	Save(ctx context.Context, GrowBox model.GrowBox)
	Update(ctx context.Context, GrowBox model.GrowBox)
	Delete(ctx context.Context, GrowBoxId int)
	FindById(ctx context.Context, GrowBoxId int) (model.GrowBox, error)
	FindAll(ctx context.Context) []model.GrowBox
}
