package poll

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var jobs = make(map[string]int)

func Poll() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		src := rand.NewSource(time.Now().UnixNano())
		name := fmt.Sprintf("job:%d", src.Int63())
		jobs[name] = 0
		fmt.Fprintf(w, "acknowledged: %s", name)
	})
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

func updateStatus() {
	for {
		for k, v := range jobs {
			if v == 100 {
				// delete(jobs, k)
				continue
			}
			jobs[k] = v + 10
		}
		time.Sleep(2 * time.Second)
	}
}
