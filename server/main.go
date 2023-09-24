package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var tpl *template.Template

func main() {

	// Load templates
	tmpl, err := template.ParseGlob("../template/*.html")
	if err != nil {
		log.Fatal(err)
	}
	tpl = tmpl

	// Creeting a server with graceful shutdown
	server := &http.Server{
		Addr: ":8081", // listen on port 9095 for incoming requests
		//Handle routes
	}

	// Serve static files from the "/static/" URL path
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	// Handle routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/solution00", solutionHandler("solution00.html"))
	http.HandleFunc("/solution02", solutionHandler("solution02.html"))
	http.HandleFunc("/solution03", solutionHandler("solution03.html"))
	http.HandleFunc("/solution04", solutionHandler("solution04.html"))
	http.HandleFunc("/solution05", solutionHandler("solution05.html"))
	http.HandleFunc("/cycle", solutionHandler("cycle.html"))

	// Start the server in a goroutine
	go func() {

		log.Printf("\n Server started at \n 👉 http://localhost:8081")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Listen for termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Shutdown the servert gracefully
	log.Println("\n ❌ Goodbye Gophers!!")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func solutionHandler(templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		tpl.ExecuteTemplate(w, templateName, nil)
	}

}
