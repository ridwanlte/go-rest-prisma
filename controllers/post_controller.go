package controllers

import (
	request "go-rest-prisma/data/request/post"
	"go-rest-prisma/data/response"
	"go-rest-prisma/helpers"
	service "go-rest-prisma/services/post"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{PostService: postService}
}

// create function
func (controllers *PostController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	PostCreateRequest := request.PostCreateRequest{}
	helpers.ReadRequestBody(requests, &PostCreateRequest)

	controllers.PostService.Create(requests.Context(), PostCreateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   writer,
	}
	helpers.WriteResponseBody(writer, webResponse)
}

// update function
func (controllers *PostController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	var PostUpdateRequest = request.PostUpdateRequest{}
	helpers.ReadRequestBody(requests, &PostUpdateRequest)

	postId := params.ByName("postId")
	PostUpdateRequest.Id = postId

	controllers.PostService.Update(requests.Context(), PostUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helpers.WriteResponseBody(writer, webResponse)
}

// findAll function
func (controllers *PostController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controllers.PostService.FindAll(requests.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}
	helpers.WriteResponseBody(writer, webResponse)
}

// findById function
func (controllers *PostController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	result := controllers.PostService.FindById(requests.Context(), postId)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}
	helpers.WriteResponseBody(writer, webResponse)
}

// delete function
func (controllers *PostController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	controllers.PostService.Delete(requests.Context(), postId)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helpers.WriteResponseBody(writer, webResponse)
}
