package service

import (
	"context"
	"goland-crud/data/request"
	"goland-crud/data/response"
)

type GrowBoxService interface {
	Create(ctx context.Context, request request.GrowBoxCreateRequest)
	Update(ctx context.Context, request request.GrowBoxUpdateRequest)
	Delete(ctx context.Context, GrowBoxId int)
	FindById(ctx context.Context, GrowBoxId int) response.GrowBoxResponse
	FindAll(ctx context.Context) []response.GrowBoxResponse
}
