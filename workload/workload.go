package workload

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/antithesishq/antithesis-sdk-go/lifecycle"
)

type Workload struct {
}

type Details map[string]any

func (w *Workload) Execute() {
	lifecycle.SetupComplete(Details{"Sandbox": "Available"})

	client := http.Client{}

	for i := 0; i < 60; i++ {
		log.Println(fmt.Sprintf("Executing workload request %d", i))
		req, err := http.NewRequest("POST", "http://server:8080/tests/1", nil)

		if err != nil {
			panic(err)
		}

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			panic("expected status code 200")
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		if string(b) != "1" {
			panic("expected body to be 1")
		}

		time.Sleep(1 * time.Second)
	}
}
