package poll

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var jobs = make(map[string]int)

func Acknowledge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "only POST method is allowed")
		return
	}

	src := rand.NewSource(time.Now().UnixNano())
	name := fmt.Sprintf("job:%d", src.Int63())
	jobs[name] = 0
	log.Printf("job: %s", name)
	fmt.Fprintf(w, "acknowledged: %s", name)
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
