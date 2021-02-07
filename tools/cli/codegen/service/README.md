# Code Generation Tooling

This CLI tool reads an api.json file and write the api operations and clients to a specific directory.

The tool can be run using the following command:

```sh
go run . --definition ../../../../definitions/service/api/v2010 --target ../../../../service/api/v2010
```

**NOTE:** This tool will create any missing directory and overwrite the existing content. Clients & Operations that have been deleted from the definition json file `api.json` will not be removed by this tool.
