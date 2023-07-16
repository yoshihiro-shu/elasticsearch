package main

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	logger := NewLogger()

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch.draft.com",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Error("failed at set es client. err is %s", err.Error())
		return
	}
	res, err := es.Ping()
	if err != nil {
		logger.Error("failed at es ping. err is %s", err.Error())
		return
	}
	defer res.Body.Close()
	logger.Infof("res is %s", res.String())
	logger.Infof("status is %s", res.Status())
}

func NewLogger() *slog.Logger {
	logger := slog.New()
	logger.SetName("BACK-END")
	setLevelHandler := handler.NewConsoleHandler(slog.AllLevels)
	logger.AddHandler(setLevelHandler)
	return logger
}
