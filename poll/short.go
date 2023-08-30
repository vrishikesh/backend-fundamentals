package poll

import (
	"fmt"
	"log"
	"net/http"
)

func ShortPolling() {
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", Acknowledge)
	mux.HandleFunc("/checkstatus", func(w http.ResponseWriter, r *http.Request) {
		jobName := r.URL.Query().Get("jobId")
		if jobName == "" {
			fmt.Fprintf(w, "jobId is required")
			return
		}
		status, ok := jobs[jobName]
		if !ok {
			fmt.Fprintf(w, "job not found")
			return
		}
		fmt.Fprintf(w, "status: %d", status)
	})

	go updateStatus()

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
