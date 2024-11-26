package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki-client-go/pkg/urlutil"
	"github.com/prometheus/common/model"
)

func main() {
	u, err := url.Parse("http://localhost:3100/loki/api/v1/push")
	handleError(err)
	cfg := loki.Config{
		URL: urlutil.URLValue{
			URL: u,
		},
		BatchWait: 5 * time.Second,
		BatchSize: 1024 * 1024,
		Timeout:   2 * time.Second,
	}

	client, err := loki.New(cfg)
	handleError(err)
	defer client.Stop()

	labels := model.LabelSet{
		"job":   "direct-from-go",
		"app":   "example-go",
		"level": "info",
	}

	i := 0
	for {
		msg := fmt.Sprintf("Hello World from Go! i=%d", i)

		// sent to Loki
		err = client.Handle(labels, time.Now(), msg)
		handleError(err)

		// print to stdout
		log.Default().Println(msg)

		i++
		time.Sleep(1 * time.Second)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
