package controller

import (
	"goland-crud/data/request"
	"goland-crud/data/response"
	"goland-crud/helper"
	"goland-crud/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type GrowBoxController struct {
	GrowBoxService service.GrowBoxService
}

func NewGrowBoxController(GrowBoxService service.GrowBoxService) *GrowBoxController {
	return &GrowBoxController{GrowBoxService: GrowBoxService}
}

func (controller *GrowBoxController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	GrowBoxCreateRequest := request.GrowBoxCreateRequest{}
	helper.ReadRequestBody(requests, &GrowBoxCreateRequest)

	controller.GrowBoxService.Create(requests.Context(), GrowBoxCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *GrowBoxController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	GrowBoxUpdateRequest := request.GrowBoxUpdateRequest{}
	helper.ReadRequestBody(requests, &GrowBoxUpdateRequest)

	GrowBoxId := params.ByName("GrowBoxId")
	id, err := strconv.Atoi(GrowBoxId)
	helper.PanicIfError(err)
	GrowBoxUpdateRequest.Id = id

	controller.GrowBoxService.Update(requests.Context(), GrowBoxUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *GrowBoxController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	GrowBoxId := params.ByName("GrowBoxId")
	id, err := strconv.Atoi(GrowBoxId)
	helper.PanicIfError(err)

	controller.GrowBoxService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *GrowBoxController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.GrowBoxService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *GrowBoxController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	GrowBoxId := params.ByName("GrowBoxId")
	id, err := strconv.Atoi(GrowBoxId)
	helper.PanicIfError(err)

	result := controller.GrowBoxService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
