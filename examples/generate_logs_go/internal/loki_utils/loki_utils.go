package loki_utils

import (
	"net/url"
	"time"

	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki-client-go/pkg/urlutil"
	"github.com/prometheus/common/model"
)

type Loki struct {
	client *loki.Client
}

func NewLokiClient(lokiURL string) (*Loki, error) {
	u, err := url.Parse(lokiURL)
	if err != nil {
		return nil, err
	}
	cfg := loki.Config{
		URL: urlutil.URLValue{
			URL: u,
		},
		BatchWait: 5 * time.Second,
		BatchSize: 1024 * 1024,
		Timeout:   2 * time.Second,
	}

	client, err := loki.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Loki{client: client}, nil
}

func (l *Loki) Log(labels model.LabelSet, t time.Time, msg string) error {
	return l.client.Handle(labels, t, msg)
}

func (l *Loki) Stop() {
	l.client.Stop()
}
