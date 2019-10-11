package main

import (
	"log"
	"monitor_api/metrics_collector"
	"monitor_api/restful_api"
	"net/http"
)

func main() {

	metrics_collector.GetClusters()

	router := restful_api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
