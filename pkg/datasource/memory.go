package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

func FetchMemory(ctx context.Context, manifest DataSource, results chan<- Result) {
	file, err := os.Open(filepath.Join("pkg", "datasource", manifest.Repository))
	if err != nil {
		fmt.Println("Error opening file:", err)
		results <- Result{ID: manifest.Name, Err: err}
		return
	}
	defer file.Close()

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		results <- Result{ID: manifest.Name, Err: err}
		return
	}

	// Dynamic selector
	selector := "features.#.properties"

	// Extract properties using the dynamic selector
	propertiesResults := gjson.Get(string(data), selector)

	// Store results in an interface slice
	var res []interface{}

	// Iterate over results and append to the slice
	propertiesResults.ForEach(func(key, value gjson.Result) bool {
		res = append(res, value.Value())
		return true // keep iterating
	})

	// Marshal the slice back to []byte
	resultBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := Result{ID: manifest.Name, Body: resultBytes}
	results <- result
}
