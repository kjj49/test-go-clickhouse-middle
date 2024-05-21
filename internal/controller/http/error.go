package http

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error" example:"message"`
}

func newErrorResponse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, errorResponse{msg})
}
