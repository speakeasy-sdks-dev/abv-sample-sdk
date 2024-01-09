# Requests
(*Requests*)

## Overview

REST APIs for retrieving request information

### Available Operations

* [GenerateRequestPostmanCollection](#generaterequestpostmancollection) - Generate a Postman collection for a particular request.
* [GetRequestFromEventLog](#getrequestfromeventlog) - Get information about a particular request.
* [QueryEventLog](#queryeventlog) - Query the event log to retrieve a list of requests.

## GenerateRequestPostmanCollection

Generates a Postman collection for a particular request. 
Allowing it to be replayed with the same inputs that were captured by the SDK.

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
    res, err := s.Requests.GenerateRequestPostmanCollection(ctx, operations.GenerateRequestPostmanCollectionRequest{
        RequestID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostmanCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GenerateRequestPostmanCollectionRequest](../../models/operations/generaterequestpostmancollectionrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.GenerateRequestPostmanCollectionResponse](../../models/operations/generaterequestpostmancollectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRequestFromEventLog

Get information about a particular request.

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
    res, err := s.Requests.GetRequestFromEventLog(ctx, operations.GetRequestFromEventLogRequest{
        RequestID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UnboundedRequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetRequestFromEventLogRequest](../../models/operations/getrequestfromeventlogrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetRequestFromEventLogResponse](../../models/operations/getrequestfromeventlogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## QueryEventLog

Supports retrieving a list of request captured by the SDK for this workspace.
Allows the filtering of requests on a number of criteria such as ApiID, VersionID, Path, Method, etc.

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
    res, err := s.Requests.QueryEventLog(ctx, operations.QueryEventLogRequest{
        Filters: &components.Filters{
            Filters: []components.Filter{
                components.Filter{
                    Key: "<key>",
                    Operator: "string",
                    Value: "string",
                },
            },
            Limit: 241978,
            Offset: 451388,
            Operator: "string",
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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.QueryEventLogRequest](../../models/operations/queryeventlogrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.QueryEventLogResponse](../../models/operations/queryeventlogresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
