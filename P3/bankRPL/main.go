package main

import (
	"os"

	"github.com/arvianta/BackendOPRECRPL/P3/bankRPL/config"
	"github.com/arvianta/BackendOPRECRPL/P3/bankRPL/handler"
	"github.com/arvianta/BackendOPRECRPL/P3/bankRPL/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(database)

	server := gin.Default()

	server.Use(
		middleware.CORSMiddleware(),
	)

	nasabahHandler := handler.NasabahHandler{DB: database}
	rekeningHandler := handler.RekeningHandler{DB: database}

	// Fungsi Nasabah
	server.POST("/addNasabah", nasabahHandler.HandleInsertNasabah)
	server.PUT("/editNasabah/:id", nasabahHandler.HandleUpdateNasabah)
	server.DELETE("/deleteNasabah/:id", nasabahHandler.HandleDeleteNasabah)

	// Fungsi Rekening
	server.POST("/addRekening", rekeningHandler.HandleInsertRekening)
	server.GET("/seeRekening/:nama", rekeningHandler.HandleGetAllRekeningByNama)
	server.DELETE("/deleteRekening/:nama/:rekeningNumber", rekeningHandler.HandleDeleteNasabahRekening)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
