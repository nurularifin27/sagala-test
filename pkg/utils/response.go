package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PaginationMeta struct {
	TotalData   int `json:"total_data"`
	TotalPage   int `json:"total_page"`
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
}

type Meta struct {
	RequestID  string          `json:"request_id,omitempty"`
	Timestamp  string          `json:"timestamp"`
	Pagination *PaginationMeta `json:"pagination,omitempty"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	response := Response{
		Status:  "success",
		Message: "Operation successful",
		Data:    data,
		Meta:    buildMeta(c, nil),
	}
	c.JSON(http.StatusOK, response)
}

func SuccessWithPagination(c *gin.Context, data interface{}, pagination *PaginationMeta) {
	response := Response{
		Status:  "success",
		Message: "Operation successful",
		Data:    data,
		Meta:    buildMeta(c, pagination),
	}
	c.JSON(http.StatusOK, response)
}

func Created(c *gin.Context, data interface{}) {
	response := Response{
		Status:  "success",
		Message: "Resource created successfully",
		Data:    data,
		Meta:    buildMeta(c, nil),
	}
	c.JSON(http.StatusCreated, response)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Error(c *gin.Context, statusCode int, message string) {
	response := Response{
		Status:  "error",
		Message: message,
		Meta:    buildMeta(c, nil),
	}
	c.JSON(statusCode, response)
}

func ErrorWithDetails(c *gin.Context, statusCode int, message string, errorDetail interface{}) {
	response := Response{
		Status:  "error",
		Message: message,
		Meta:    buildMeta(c, nil),
		Error:   errorDetail,
	}
	c.JSON(statusCode, response)
}

func NotFound(c *gin.Context, message string) {
	response := Response{
		Status:  "error",
		Message: message,
		Meta:    buildMeta(c, nil),
	}
	c.JSON(http.StatusNotFound, response)
}

func buildMeta(c *gin.Context, pagination *PaginationMeta) *Meta {
	meta := &Meta{
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if requestID, exists := c.Get("X-Request-ID"); exists {
		meta.RequestID = requestID.(string)
	}

	// Add pagination if provided
	if pagination != nil {
		meta.Pagination = pagination
	}

	return meta
}
