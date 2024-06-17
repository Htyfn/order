package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"order/internal/config"
	"order/internal/sl"
	"strconv"
	"strings"
)

func main() {

	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.LogFormat, cfg.LogLevel)

	log.Info("Starting app: 'Order'", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
	id := 1

	var endpoint strings.Builder
	endpoint.WriteString(cfg.ProductEndpoint)
	endpoint.WriteString("/product/")
	endpoint.WriteString(strconv.Itoa(id))
	log.Info("logloglog", slog.String("endpoint", endpoint.String()))

	resp, err := http.Get(endpoint.String())
	if err != nil {
		log.Error("err while sending request to Product", sl.Err(err))
	}
	defer resp.Body.Close()

	log.Info("logloglog", slog.String("Response status:", resp.Status))

	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Error("err while parsing response by id "+strconv.Itoa(id), sl.Err(err))
	}
	fmt.Println(string(body))

}
