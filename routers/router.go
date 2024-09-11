package routers

import (
	"go-rest-prisma/controllers"
	"go-rest-prisma/data/response"
	"go-rest-prisma/helpers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(postController *controllers.PostController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// fmt.Println(w, "Welcome Home")
		webResponse := response.WebResponse{
			Code:   200,
			Status: "ok",
			Data:   "WELCOME HOME",
		}
		helpers.WriteResponseBody(w, webResponse)
	})

	router.GET("/api/v1/posts", postController.FindAll)
	router.GET("/api/v1/post/:postId", postController.FindById)
	router.POST("/api/v1/post", postController.Create)
	router.PATCH("/api/v1/post/:postId", postController.Update)
	router.DELETE("/api/v1/post/:postId", postController.Delete)

	return router
}
