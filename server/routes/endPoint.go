package routes

import (
	"example/service"
	"github.com/labstack/echo"
	"github.com/rs/cors"
)

func Endpoint() {
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	//user endpoint
	e.GET("/posts/readAll", service.ReadAllPosts)
	e.GET("/article/Limit", service.ReadPostsLimit)
	e.POST("/article", service.CreatePosts)
	e.GET("/article", service.ReadPostsById)
	e.PUT("/article", service.UpdatePosts)
	e.DELETE("/article", service.DeletePosts)
	e.GET("/article/publish", service.ReadPostsPublish)
	e.GET("/article/draft", service.ReadPostsDraft)
	e.GET("/article/trash", service.ReadPostsTrash)
	e.PUT("/article/updateToTrash", service.UpdateStatusToTrash)

	e.Logger.Fatal(e.Start(":1323"))
}