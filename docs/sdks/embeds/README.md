# Embeds
(*Embeds*)

## Overview

REST APIs for managing embeds

### Available Operations

* [GetEmbedAccessToken](#getembedaccesstoken) - Get an embed access token for the current workspace.
* [GetValidEmbedAccessTokens](#getvalidembedaccesstokens) - Get all valid embed access tokens for the current workspace.
* [RevokeEmbedAccessToken](#revokeembedaccesstoken) - Revoke an embed access EmbedToken.

## GetEmbedAccessToken

Returns an embed access token for the current workspace. This can be used to authenticate access to externally embedded content.
Filters can be applied allowing views to be filtered to things like particular customerIds.

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
    res, err := s.Embeds.GetEmbedAccessToken(ctx, operations.GetEmbedAccessTokenRequest{
        Filters: &components.Filters{
            Filters: []components.Filter{
                components.Filter{
                    Key: "<key>",
                    Operator: "string",
                    Value: "string",
                },
            },
            Limit: 964408,
            Offset: 95617,
            Operator: "string",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EmbedAccessTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetEmbedAccessTokenRequest](../../models/operations/getembedaccesstokenrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetEmbedAccessTokenResponse](../../models/operations/getembedaccesstokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetValidEmbedAccessTokens

Get all valid embed access tokens for the current workspace.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks-dev/abv-sample-sdk/models/components"
	abvsamplesdk "github.com/speakeasy-sdks-dev/abv-sample-sdk"
	"context"
	"log"
)

func main() {
    s := abvsamplesdk.New(
        abvsamplesdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Embeds.GetValidEmbedAccessTokens(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetValidEmbedAccessTokensResponse](../../models/operations/getvalidembedaccesstokensresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RevokeEmbedAccessToken

Revoke an embed access EmbedToken.

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
    res, err := s.Embeds.RevokeEmbedAccessToken(ctx, operations.RevokeEmbedAccessTokenRequest{
        TokenID: "string",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RevokeEmbedAccessTokenRequest](../../models/operations/revokeembedaccesstokenrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RevokeEmbedAccessTokenResponse](../../models/operations/revokeembedaccesstokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
