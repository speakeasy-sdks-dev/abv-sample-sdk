<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	abvsamplesdk "github.com/speakeasy-sdks-dev/abv-sample-sdk"
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/components"
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/operations"
	"log"
)

func main() {
	s := abvsamplesdk.New(
		abvsamplesdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
		Metadata: map[string][]string{
			"key": []string{
				"string",
			},
		},
		Op: &operations.QueryParamOp{
			And: false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->