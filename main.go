package main

import (
	"github.com/bearnadbear/mrt-schedules/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	initiateRouter()
}

func initiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	station.Initiate(api)

	router.Run(":8080")
}
