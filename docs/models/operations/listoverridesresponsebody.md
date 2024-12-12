# ListOverridesResponseBody

List of overrides for the given namespace.


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Overrides`                                                                                         | [][operations.Overrides](../../models/operations/overrides.md)                                      | :heavy_check_mark:                                                                                  | N/A                                                                                                 |                                                                                                     |
| `Cursor`                                                                                            | **string*                                                                                           | :heavy_minus_sign:                                                                                  | The cursor to use for the next page of results, if no cursor is returned, there are no more results | eyJrZXkiOiJrZXlfMTIzNCJ9                                                                            |
| `Total`                                                                                             | *int64*                                                                                             | :heavy_check_mark:                                                                                  | The total number of overrides for the namespace                                                     |                                                                                                     |