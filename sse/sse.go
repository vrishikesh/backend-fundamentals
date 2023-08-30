package sse

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func SSE() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello SSE")
	})
	mux.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		f, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "data: Message %d\n\n", i)
			f.Flush()
			time.Sleep(1 * time.Second)
		}
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}
