package poll

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func LongPolling() {
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", Acknowledge)
	mux.HandleFunc("/checkstatus", func(w http.ResponseWriter, r *http.Request) {
		jobName := r.URL.Query().Get("jobId")
		if jobName == "" {
			fmt.Fprintf(w, "jobId is required")
			return
		}
		done := make(chan struct{})
		log.Printf("job: %s", jobName)
		status := jobs[jobName]
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		go func() {
			defer close(done)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					status = jobs[jobName]
					log.Printf("status: %d", status)
					if status == 100 {
						return
					}
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()
		<-done
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
