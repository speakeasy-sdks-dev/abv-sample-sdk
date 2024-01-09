# Metadata
(*Metadata*)

## Overview

REST APIs for managing Version Metadata entities

### Available Operations

* [DeleteVersionMetadata](#deleteversionmetadata) - Delete metadata for a particular apiID and versionID.
* [GetVersionMetadata](#getversionmetadata) - Get all metadata for a particular apiID and versionID.
* [InsertVersionMetadata](#insertversionmetadata) - Insert metadata for a particular apiID and versionID.

## DeleteVersionMetadata

Delete metadata for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/components"
	abvsamplesdk "github.com/speakeasy-sdks-dev/abv-sample-sdk"
	"context"
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/operations"
	"log"
	"net/http"
)

func main() {
    s := abvsamplesdk.New(
        abvsamplesdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Metadata.DeleteVersionMetadata(ctx, operations.DeleteVersionMetadataRequest{
        APIID: "string",
        MetaKey: "string",
        MetaValue: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteVersionMetadataRequest](../../models/operations/deleteversionmetadatarequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.DeleteVersionMetadataResponse](../../models/operations/deleteversionmetadataresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetVersionMetadata

Get all metadata for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/components"
	abvsamplesdk "github.com/speakeasy-sdks-dev/abv-sample-sdk"
	"context"
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/operations"
	"log"
)

func main() {
    s := abvsamplesdk.New(
        abvsamplesdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Metadata.GetVersionMetadata(ctx, operations.GetVersionMetadataRequest{
        APIID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetVersionMetadataRequest](../../models/operations/getversionmetadatarequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetVersionMetadataResponse](../../models/operations/getversionmetadataresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## InsertVersionMetadata

Insert metadata for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/components"
	abvsamplesdk "github.com/speakeasy-sdks-dev/abv-sample-sdk"
	"context"
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/operations"
	"log"
)

func main() {
    s := abvsamplesdk.New(
        abvsamplesdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Metadata.InsertVersionMetadata(ctx, operations.InsertVersionMetadataRequest{
        VersionMetadata: components.VersionMetadataInput{
            MetaKey: "string",
            MetaValue: "string",
        },
        APIID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.InsertVersionMetadataRequest](../../models/operations/insertversionmetadatarequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.InsertVersionMetadataResponse](../../models/operations/insertversionmetadataresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
