package main

import (
	"csv-reader/handlers"
	"embed"
	"flag"
	"fmt"
	"github.com/rs/cors"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	var port int

	flag.IntVar(&port, "port", 3000, "The port to listen on")
	flag.Parse()

	frontendAssets := getProdFrontendAssets()
	corsMiddleware := cors.New(cors.Options{AllowedOrigins: []string{"http://localhost:8080"}})

	http.Handle("/api/v1/law", corsMiddleware.Handler(http.HandlerFunc(handlers.GetRandomLaw)))
	http.Handle("/", http.FileServer(http.FS(frontendAssets)))

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getProdFrontendAssets() fs.FS {
	f, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		panic(err)
	}

	return f
}

func getDevFrontendAssets() fs.FS {
	return os.DirFS("frontend/dist")
}
