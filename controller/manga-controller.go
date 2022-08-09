package controller

import (
	"fmt"
	"mini-project-nunu/dto"
	"mini-project-nunu/entity"
	"mini-project-nunu/helper"
	"mini-project-nunu/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// * MangaController is a contract
type MangaController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type mangaController struct {
	mangaService service.MangaService
	jwtService   service.JWTService
}

func NewMangaController(mangaServ service.MangaService, jwtServ service.JWTService) MangaController {
	return &mangaController{mangaService: mangaServ, jwtService: jwtServ}
}

func (c *mangaController) All(context *gin.Context) {
	var allManga []entity.Manga = c.mangaService.AllManga()
	res := helper.BuildResponse(true, "OK", allManga)
	context.JSON(http.StatusOK, res)
}

func (c *mangaController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("no param ID was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var manga entity.Manga = c.mangaService.FindByID(id)
	if (manga == entity.Manga{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", manga)
		context.JSON(http.StatusOK, res)
	}
}
func (c *mangaController) Insert(context *gin.Context) {
	var mangaCreateDTO dto.MangaCreateDTO
	errDTO := context.ShouldBind(&mangaCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			mangaCreateDTO.UserID = convertedUserID
		}
		result := c.mangaService.Insert(mangaCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}
func (c *mangaController) Update(context *gin.Context) {
	var mangaUpdateDTO dto.MangaUpdateDTO
	errDTO := context.ShouldBind(&mangaUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.mangaService.IsAllowedToEdit(userID, mangaUpdateDTO.MangaID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			mangaUpdateDTO.UserID = id
		}
		result := c.mangaService.Update(mangaUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *mangaController) Delete(context *gin.Context) {
	var manga entity.Manga
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	manga.MangaID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.mangaService.IsAllowedToEdit(userID, manga.MangaID) {
		c.mangaService.Delete(manga)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *mangaController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
