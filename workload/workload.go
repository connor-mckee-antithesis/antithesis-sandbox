package workload

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	// "strconv"

	"github.com/antithesishq/antithesis-sdk-go/lifecycle"
	"github.com/antithesishq/antithesis-sdk-go/random"
)

type Workload struct {
}

type Details map[string]any

func (w *Workload) Execute() {

	// extra time to ensure ledger is up
	time.Sleep(time.Duration(5000) * time.Millisecond)

	lifecycle.SetupComplete(Details{"Sandbox": "Available"})

	client := http.Client{}

	for i := 0; i < 1e7; i++ {
		log.Println(fmt.Sprintf("Executing workload request %d", i))
		// using IP instead of hostname
		req, err := http.NewRequest("POST", "http://10.0.0.16:8080/tests/1", nil)

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

		log.Println(fmt.Sprintf("Response: %s", string(b)))

		sleep_millis := int(random.GetRandom() % 10)

		// log.Println(fmt.Sprintf("Sleeping for: %s", strconv.Itoa(sleep_millis)))

		time.Sleep(time.Duration(sleep_millis) * time.Millisecond)
	}
}
