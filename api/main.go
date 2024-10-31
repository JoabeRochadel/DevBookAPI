package main

import (
	"DevBookAPI/src/config"
	"DevBookAPI/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Generate()
	fmt.Printf("Rodando a API na porta: %d\n", config.Port)
	fmt.Println(config.ConnectionString)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
