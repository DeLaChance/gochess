package adapter

import (
	"config"
	"fmt"
	"net/http"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChessController struct {
	service *service.ChessGameService
}

func StartChessController(applicationConfig *config.Config, service *service.ChessGameService) {
	controller := ChessController{service: service}

	// TODO: can this be done in a better way?
	router := gin.Default()
	router.GET("/api/game/:id", controller.GetGameById)
	router.POST("/api/game/", controller.CreateNewGame)
	router.PUT("/api/game/:id/start", controller.StartNewGame)

	hostName := fmt.Sprintf("%s:%d", applicationConfig.HttpHost, applicationConfig.HttpPort)
	router.Run(hostName)
	config.Info.Printf("Started HTTP server at %s", hostName)
}

func (controller *ChessController) GetGameById(context *gin.Context) {

	gameId, err := strconv.Atoi(context.Param("id"))
	if err == nil {
		chessGame, error := controller.service.FetchById(gameId)
		if error == nil {
			chessGameDto := GenerateChessGameDto(chessGame)
			context.IndentedJSON(http.StatusOK, chessGameDto)
		} else {
			context.Status(http.StatusNotFound)
		}
	} else {
		context.Status(http.StatusBadRequest)
	}
}

func (controller *ChessController) CreateNewGame(context *gin.Context) {
	game, error := controller.service.CreateNewGame()
	if error == nil {
		context.IndentedJSON(http.StatusCreated, GameCreatedDto{GameID: int(game.ID)})
	} else {
		context.Status(http.StatusInternalServerError)
	}
}

func (controller *ChessController) StartNewGame(context *gin.Context) {

	gameId, numberConversionError := strconv.Atoi(context.Param("id"))
	if numberConversionError == nil {
		game, error := controller.service.StartNewGame(gameId)
		if error == nil {
			context.IndentedJSON(http.StatusCreated, GameCreatedDto{GameID: int(game.ID)})
		} else {
			context.Status(http.StatusInternalServerError)
		}
	} else {
		context.Status(http.StatusBadRequest)
	}

}
