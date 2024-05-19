package main

/*
#cgo LDFLAGS: -L./lib -lcalc_lib
#include <stdlib.h>
#include "calc_lib.h"
*/
import "C"

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/add/:a/:b", func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		result := C.add(C.int(a), C.int(b))
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})
	r.GET("/subtract/:a/:b", func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		result := C.subtract(C.int(a), C.int(b))
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})
	r.GET("/multiply/:a/:b", func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		result := C.multiply(C.int(a), C.int(b))
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})
	r.GET("/divide/:a/:b", func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "division by zero",
			})
			return
		}
		result := C.divide(C.int(a), C.int(b))
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})
	r.Run()
}
