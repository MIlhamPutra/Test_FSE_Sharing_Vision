package service

import (
	"example/repository"
	"example/models"
	"net/http"
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo"
)

// ResponseModel function
type ResponseModel struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func ReadAllPosts(c echo.Context) error {
	result := repository.ReadAllPosts()
	return c.JSON(http.StatusOK, result)
}

func CreatePosts(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.PostsModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	Res = (*ResponseModel)(repository.CreatePosts(U))
	return c.JSON(http.StatusOK, Res)
}

func ReadPostsById(c echo.Context) error {
	id := c.QueryParam("Id")
	data, _ := strconv.Atoi(id)
	result := repository.ReadPostsById(data)
	return c.JSON(http.StatusOK, result)
}

func UpdatePosts(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.QueryParam("Id")
	data, _ := strconv.Atoi(id)
	U := new(models.PostsModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdatePosts(U, data))
	return c.JSON(http.StatusOK, Res)
}

func DeletePosts(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.QueryParam("Id")
	data, _ := strconv.Atoi(id)
	fmt.Println("id", data)
	Res = (*ResponseModel)(repository.DeletePosts(data))
	return c.JSON(http.StatusOK, Res)
}

func ReadPostsLimit(c echo.Context) error {
	offset := c.QueryParam("Offset")
	limit := c.QueryParam("Limit")
	data, _ := strconv.Atoi(offset)
	data2, _ := strconv.Atoi(limit)
	result := repository.ReadPostsLimit(data, data2)
	return c.JSON(http.StatusOK, result)
}

func ReadPostsPublish(c echo.Context) error {
	result := repository.ReadPostsPublish()
	return c.JSON(http.StatusOK, result)
}

func ReadPostsDraft(c echo.Context) error {
	result := repository.ReadPostsDraft()
	return c.JSON(http.StatusOK, result)
}

func ReadPostsTrash(c echo.Context) error {
	result := repository.ReadPostsTrash()
	return c.JSON(http.StatusOK, result)
}

func UpdateStatusToTrash(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.QueryParam("Id")
	data, _ := strconv.Atoi(id)
	Res = (*ResponseModel)(repository.UpdateStatusToTrash(data))
	return c.JSON(http.StatusOK, Res)
}