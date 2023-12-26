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

const chronoEndpoint = "http://172.18.48.1:4445"
const tenant = "tenant-e0a5a0e3-df10-4059-a538-39033f21c4ff"

func FetchChrono(ctx context.Context, manifest DataSource, results chan<- Result) {
	startTime := time.Now()

	logger.GetLogger(ctx).WithFields(map[string]any{
		"type": "data",
		"name": manifest.Name,
	}).Info("fetching")

	url := fmt.Sprintf("%s/%s/metrics/%s/time-partitioned", chronoEndpoint, tenant, manifest.Repository)
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
