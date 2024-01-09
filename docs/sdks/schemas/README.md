# Schemas
(*Schemas*)

## Overview

REST APIs for managing Schema entities

### Available Operations

* [DeleteSchema](#deleteschema) - Delete a particular schema revision for an Api.
* [DownloadSchema](#downloadschema) - Download the latest schema for a particular apiID.
* [DownloadSchemaRevision](#downloadschemarevision) - Download a particular schema revision for an Api.
* [GetSchema](#getschema) - Get information about the latest schema.
* [GetSchemaDiff](#getschemadiff) - Get a diff of two schema revisions for an Api.
* [GetSchemaRevision](#getschemarevision) - Get information about a particular schema revision for an Api.
* [GetSchemas](#getschemas) - Get information about all schemas associated with a particular apiID.
* [RegisterSchema](#registerschema) - Register a schema.

## DeleteSchema

Delete a particular schema revision for an Api.

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
    res, err := s.Schemas.DeleteSchema(ctx, operations.DeleteSchemaRequest{
        APIID: "string",
        RevisionID: "string",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteSchemaRequest](../../models/operations/deleteschemarequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteSchemaResponse](../../models/operations/deleteschemaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DownloadSchema

Download the latest schema for a particular apiID.

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
    res, err := s.Schemas.DownloadSchema(ctx, operations.DownloadSchemaRequest{
        APIID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DownloadSchemaRequest](../../models/operations/downloadschemarequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DownloadSchemaResponse](../../models/operations/downloadschemaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DownloadSchemaRevision

Download a particular schema revision for an Api.

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
    res, err := s.Schemas.DownloadSchemaRevision(ctx, operations.DownloadSchemaRevisionRequest{
        APIID: "string",
        RevisionID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DownloadSchemaRevisionRequest](../../models/operations/downloadschemarevisionrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DownloadSchemaRevisionResponse](../../models/operations/downloadschemarevisionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSchema

Returns information about the last uploaded schema for a particular API version. 
This won't include the schema itself, that can be retrieved via the downloadSchema operation.

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
    res, err := s.Schemas.GetSchema(ctx, operations.GetSchemaRequest{
        APIID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetSchemaRequest](../../models/operations/getschemarequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetSchemaResponse](../../models/operations/getschemaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSchemaDiff

Get a diff of two schema revisions for an Api.

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
    res, err := s.Schemas.GetSchemaDiff(ctx, operations.GetSchemaDiffRequest{
        APIID: "string",
        BaseRevisionID: "string",
        TargetRevisionID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaDiff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetSchemaDiffRequest](../../models/operations/getschemadiffrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetSchemaDiffResponse](../../models/operations/getschemadiffresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSchemaRevision

Returns information about the last uploaded schema for a particular schema revision. 
This won't include the schema itself, that can be retrieved via the downloadSchema operation.

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
    res, err := s.Schemas.GetSchemaRevision(ctx, operations.GetSchemaRevisionRequest{
        APIID: "string",
        RevisionID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSchemaRevisionRequest](../../models/operations/getschemarevisionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetSchemaRevisionResponse](../../models/operations/getschemarevisionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSchemas

Returns information the schemas associated with a particular apiID. 
This won't include the schemas themselves, they can be retrieved via the downloadSchema operation.

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
    res, err := s.Schemas.GetSchemas(ctx, operations.GetSchemasRequest{
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetSchemasRequest](../../models/operations/getschemasrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetSchemasResponse](../../models/operations/getschemasresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RegisterSchema

Allows uploading a schema for a particular API version.
This will be used to populate ApiEndpoints and used as a base for any schema generation if present.

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
    res, err := s.Schemas.RegisterSchema(ctx, operations.RegisterSchemaRequest{
        RequestBody: operations.RegisterSchemaRequestBody{
            File: operations.File{
                Content: []byte("0xCAFCA03e0e"),
                FileName: "bronze_table_blues.m2a",
            },
        },
        APIID: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.RegisterSchemaRequest](../../models/operations/registerschemarequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.RegisterSchemaResponse](../../models/operations/registerschemaresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
