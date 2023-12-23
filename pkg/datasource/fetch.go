package datasource

import "context"

func Fetch(ctx context.Context, manifest DataSource, results chan<- Result) {
	switch manifest.Provider {
	case "chrono":
		FetchChrono(ctx, manifest, results)
	case "memory":
		FetchMemory(ctx, manifest, results)
	}
}
