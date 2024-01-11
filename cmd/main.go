package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rontivero/entrevistatecnica/pkg/configs"
	"github.com/rontivero/entrevistatecnica/pkg/domain/album"
	"github.com/rontivero/entrevistatecnica/pkg/domain/currencies"
	"github.com/rontivero/entrevistatecnica/pkg/repository/http"
	"github.com/rontivero/entrevistatecnica/pkg/repository/storage"
	"github.com/rontivero/entrevistatecnica/pkg/rest/reader"
	"github.com/rontivero/entrevistatecnica/pkg/rest/writer"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

func main() {
	cfg := getConfig()
	router := gin.Default()

	// Capa de Repositorios
	storageRepo := storage.NewMapStorage()
	//storageRepo := storage.NewSliceStorage() // Elegir entre un almacenamiento tipo relacional o no relacional
	currenciesRepo := http.NewCurrenciesRepository()

	// Capa de Servicios
	currenciesService := currencies.NewService(currenciesRepo)
	albumSrv := album.NewService(storageRepo, currenciesService)

	// Capa de Manejadores de rutas
	readerHandler := reader.NewHandler(albumSrv)
	readerHandler.GetRoutes(router)

	writerHandler := writer.NewHandler(albumSrv)
	writerHandler.GetRoutes(router)

	// Inicio de aplicación
	router.Run(fmt.Sprintf("%s:%s", cfg.Domain, cfg.Port)) // Inicializo la aplicación
}

func getConfig() configs.Config {
	file, err := os.Open("config/local.yml")
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var config configs.Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
