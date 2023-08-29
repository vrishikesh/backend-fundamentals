package req_res

import (
	"io"
	"log"
	"net/http"
)

func ReqRes() {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		log.Printf("Error doing request: %s", err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading body: %s", err)
		return
	}

	log.Printf("Status: %s\n", res.Status)
	log.Printf("Body: %s\n", body)
}
