package main

import (
	"fmt"
	"generate_logs_go/internal/loki_utils"
	"generate_logs_go/internal/random_utils"
	"log"
	"time"

	"github.com/prometheus/common/model"
)

func main() {
	loki, err := loki_utils.NewLokiClient("http://localhost:3100/loki/api/v1/push")
	handleError(err)
	defer loki.Stop()

	// 5 days backfill
	totalSeconds := 5 * 24 * 60 * 60             // 432000
	start := time.Now().Add(-5 * 24 * time.Hour) // timestamp 5 days ago

	for i := 0; i < totalSeconds; i++ {
		ts := start.Add(time.Duration(i) * time.Second)

		msg := fmt.Sprintf("%s i=%d", random_utils.GetRandomMessage(), i)

		err = loki.Log(
			model.LabelSet{
				"job":    "generate-logs-go",
				"app":    "example-from-generator",
				"level":  model.LabelValue(random_utils.GetRandomLogLevel()),
				"method": model.LabelValue(random_utils.GetRandomMethod()),
				"path":   model.LabelValue(random_utils.GetRandomPath()),
			},
			ts,
			msg,
		)
		handleError(err)
		time.Sleep(time.Microsecond)
	}

	fmt.Println("Done! Generated 5 days of logs (1 log/second).")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
