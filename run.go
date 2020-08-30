package wwwIEDev

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func Run() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.LoggingHandler(log.Writer(), http.DefaultServeMux))
}
