# GetAllAPIVersionsRequest


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `APIID`                                         | *string*                                        | :heavy_check_mark:                              | The ID of the Api to retrieve.                  |
| `Metadata`                                      | map[string][]*string*                           | :heavy_minus_sign:                              | Metadata to filter Apis on                      |
| `Op`                                            | [*operations.Op](../../models/operations/op.md) | :heavy_minus_sign:                              | Configuration for filter operations             |