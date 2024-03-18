package service

import (
	"context"
	"goland-crud/data/request"
	"goland-crud/data/response"
	"goland-crud/helper"
	"goland-crud/model"
	"goland-crud/repository"
)

type GrowBoxServiceImpl struct {
	GrowBoxRepository repository.GrowBoxRepository
}

func NewGrowBoxServiceImpl(GrowBoxRepository repository.GrowBoxRepository) GrowBoxService {
	return &GrowBoxServiceImpl{GrowBoxRepository: GrowBoxRepository}
}

// Create implements GrowBoxService.
func (b *GrowBoxServiceImpl) Create(ctx context.Context, request request.GrowBoxCreateRequest) {
	GrowBox := model.GrowBox{
		Name:        request.Name,
		Description: request.Description,
		Count_fans:  request.Count_fans,
		Filtration:  request.Filtration,
		Dimensions:  request.Dimensions,
		Automation:  request.Automation,
	}
	b.GrowBoxRepository.Save(ctx, GrowBox)
}

// Delete implements GrowBoxService.
func (b *GrowBoxServiceImpl) Delete(ctx context.Context, GrowBoxId int) {
	GrowBox, err := b.GrowBoxRepository.FindById(ctx, GrowBoxId)
	helper.PanicIfError(err)
	b.GrowBoxRepository.Delete(ctx, GrowBox.Id)
}

// FindAll implements GrowBoxService.
func (b *GrowBoxServiceImpl) FindAll(ctx context.Context) []response.GrowBoxResponse {
	GrowBoxs := b.GrowBoxRepository.FindAll(ctx)

	var GrowBoxResp []response.GrowBoxResponse

	for _, value := range GrowBoxs {
		GrowBox := response.GrowBoxResponse{
			Id:          value.Id,
			Name:        value.Name,
			Description: value.Description,
			Count_fans:  value.Count_fans,
			Filtration:  value.Filtration,
			Dimensions:  value.Dimensions,
			Automation:  value.Automation,
		}

		GrowBoxResp = append(GrowBoxResp, GrowBox)
	}

	return GrowBoxResp
}

// FindById implements GrowBoxService.
func (b *GrowBoxServiceImpl) FindById(ctx context.Context, GrowBoxId int) response.GrowBoxResponse {
	GrowBox, err := b.GrowBoxRepository.FindById(ctx, GrowBoxId)
	helper.PanicIfError(err)
	return response.GrowBoxResponse(GrowBox)
}

// Upadate implements GrowBoxService.
func (b *GrowBoxServiceImpl) Update(ctx context.Context, request request.GrowBoxUpdateRequest) {
	GrowBox, err := b.GrowBoxRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	GrowBox.Name = request.Name
	GrowBox.Description = request.Description
	GrowBox.Count_fans = request.Count_fans
	GrowBox.Filtration = request.Filtration
	GrowBox.Dimensions = request.Dimensions
	GrowBox.Automation = request.Automation

	b.GrowBoxRepository.Update(ctx, GrowBox)
}
