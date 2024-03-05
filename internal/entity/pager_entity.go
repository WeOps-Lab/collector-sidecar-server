package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil"
)

type PagerEntity struct {
	Current int   `json:"current"`
	Size    int   `json:"size"`
	Total   int64 `json:"total"`
}

func ExtractPageParam(c *gin.Context) (int, int) {
	current := c.DefaultQuery("current", "0")
	size := c.DefaultQuery("size", "10")
	return goutil.Int(current), goutil.Int(size)
}
