package server

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	"github.com/kironono/browsercheck/internal/model"
	"github.com/kironono/browsercheck/web"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := web.GetTemplates()
	tmpl, _ := template.ParseFS(t, "index.html")

	type Data struct{}

	err := tmpl.Execute(w, &Data{})
	if err != nil {
		log.Println(err)
	}
}

func Run(config *model.Config) error {
	log.Printf("Config: %v\n", config)

	http.HandleFunc("/", index)

	frontend := web.GetFrontendAssets()
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.FS(frontend))))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr: config.ServerSettings.ListenAddress,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	log.Printf("starting server at %s\n", config.ServerSettings.ListenAddress)
	server.ListenAndServe()

	return nil
}
