package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-element-vue-backend/model"
	"hello-element-vue-backend/util"
	"net/http"
	"time"
)

// 按页获取乐谱
func getScoreListAtPage(ctx *gin.Context) {

}

func getAllScore(ctx *gin.Context) {
	scores, err := model.GetAllGuitarScore()
	if err != nil {
		ctx.JSON(http.StatusOK, model.FailureResponse(-1, err.Error()))
		return
	}
	fmt.Printf("%+v\n", scores)
	ctx.JSON(http.StatusOK, model.SuccessResponse(scores))
}

// 创建乐谱
func createNewScore(ctx *gin.Context) {
	var score model.GuitarScore
	if err := ctx.ShouldBindJSON(&score); err != nil {
		ctx.JSON(http.StatusBadRequest, model.FailureResponse(-1, err.Error()))
		return
	}

	// 生成 UID
	score.CreatedAt = time.Now()
	score.UpdatedAt = score.CreatedAt
	score.UID = util.NewUID().String()
	for i := range score.Pics {
		score.Pics[i].UID = util.NewUID().String()
		score.Pics[i].ScoreUID = score.UID
		score.Pics[i].CreatedAt = score.CreatedAt
	}

	err := model.CreateGuitarScore(&score)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.FailureResponse(-1, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse(score))
}

// DELETE /guitar/score/:uid
func deleteScore(ctx *gin.Context) {
	uid := ctx.Param("uid")

	score, err := model.GetGuitarScore(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.FailureResponse(-1, err.Error()))
		return
	}

	err = model.DeleteGuitarScore(&score)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.FailureResponse(-1, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse(score))
}