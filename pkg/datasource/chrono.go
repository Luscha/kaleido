package datasource

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gitlab.com/technity/go-x/pkg/logger"
	"gitlab.com/technity/go-x/pkg/xhttp"
)

const chronoEndpoint = "http://172.30.224.1:4445"
const tenant = "tenant-9fe806d5-28d8-45a3-a272-f577db014f67"

func FetchChrono(ctx context.Context, manifest DataSource, results chan<- Result) {
	startTime := time.Now()

	logger.GetLogger(ctx).WithFields(map[string]any{
		"type": "data",
		"name": manifest.Name,
	}).Info("fetching")

	url := fmt.Sprintf("%s/%s/metrics/product-financials/time-partitioned", chronoEndpoint, tenant)
	jsonManifest, err := json.Marshal(manifest.Manifest)
	if err != nil {
		results <- Result{ID: manifest.Name, Err: err}
		return
	}

	client := xhttp.GetHttpClient(ctx)
	response, err := client.Post(url, "application/json", bytes.NewBuffer(jsonManifest))

	if response.StatusCode != http.StatusOK {
		results <- Result{ID: manifest.Name, Status: response.StatusCode, Err: err}
		return
	}

	defer response.Body.Close()
	result := Result{ID: manifest.Name, Status: response.StatusCode}
	result.Body, err = io.ReadAll(response.Body)
	if err != nil {
		results <- Result{ID: manifest.Name, Status: response.StatusCode, Err: err}
		return
	}

	fmt.Printf("fetch '%s' took: %v\n", manifest.Name, time.Now().Sub(startTime))
	results <- result
}
