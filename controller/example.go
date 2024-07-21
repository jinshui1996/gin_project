package controller

import (
	"gin_project/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var agent = service.NewAgent(service.NewGetterFunc(), 1000)
var consistentHash = service.NewHashService()

type ExampleController struct{}

func (exampleController *ExampleController) TestCache(ctx *gin.Context) {
	var name = ctx.Query("name")
	log.Default().Println(name)
	val, err := agent.Get(name)
	if err != nil {
		// log.Fatal(err)
		log.Default().Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	log.Println(val)
	// 生成一个gin的返回值
	ctx.JSON(http.StatusOK, val)
}

func (exampleController *ExampleController) TestConsistentHash(ctx *gin.Context) {
	consistentHash.AddHashKey("test")
	if consistentHash.GetHashKey("test1") != "test1" {
		log.Println("consistent not have ", "test1")
	}
	if consistentHash.GetHashKey("test") != "test" {
		log.Println("consistent not have ", "test")
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "consistent hash error"})
		return

	}
	log.Println("consistent have ", "test")
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "api ok"})
}
