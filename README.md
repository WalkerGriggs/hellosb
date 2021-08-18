# hellosb

`hellosb` is a (near) feature complete reference implementation of the [Open Service Broker](https://github.com/openservicebrokerapi/servicebroker) [specification](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md), generated using [go-swagger](https://github.com/go-swagger/go-swagger).

Hellosb is designed as a target for integration tests for OSB clients and libraries. Catalogs are injected as a yaml file and the http server can either be embedded within your test framework or run solo.

Feature include:

- custom catalogs injected at runtime
- embeddable http server for automated testing
- global resource parameters to change runtime behaviors

## Getting started

### Stand-alone

Build the project like any Go project. Dependencies are vendored.

```
$ go build -o hellosb
```

Check flags with `./hellosb --help`, but a reasonable place to start is:

```
$ ./hellosb server --host localhost --port 3000
```

### Embedded

The generated server can be embedded and served from any go project. You may, for example, want to run a local server embedded in your test framework.

For example:

``` go
package main

import (
    "github.com/go-openapi/loads"

    "github.com/walkergriggs/hellosb/api"
    "github.com/walkergriggs/hellosb/server"
)

func main() {
    // Load the OSB swagger specification
    swaggerSpec, err := loads.Embedded(api.SwaggerJSON, api.FlatSwaggerJSON)
    if err != nil {
        panic(err)
    }

    // Create a new API
    restapi := api.New(swaggerSpec)

    // Configure the catalog path
    restapi.Configure(&api.Config{
        CatalogPath: "/Users/w.griggs/Documents/catalog.json",
    })

    // Define server options
    opts := server.Options{
        Host: "localhost",
        Port: 3000,
        API:  restapi,
    }

    // Run the server
    if err := server.New(opts).Serve(); err != nil {
        panic(err)
    }
}
```

## Usage

The default `mocks/catalog.json` file covers basic-use cases, but you can update it to better meet your needs.

The json matches the [catalog specification](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#service-offering-object), and there's a blank catalog template in the `mocks` directory.

### Instance / Binding Parameters

Provisioning behaviors of service instances and binding can be adjusted with global parameters.

| Name      | Description | Example |
| --------- | ----------- | ------- |
| `opts_provision_duration` | Alters the duration of time before a resource is successfully provisioned. A duration string is a decimal number with a unix suffix ("h", "m", "s", etc). Only applicable for async operations. Defaults to 0m. | "30s", "1h15m" |

```
curl -vv -u user:pass http://127.0.0.1:3000/v2/service_instances/123-4?accepts_incomplete=true -d '{
     "service_id": "123-4",
     "plan_id": "123-4",
     "organization_guid": "org-guid",
     "space_guid": "space-guid",
     "parameters": {
                   "opts_provision_duration": "3m"
     }
}' -X PUT -H "X-Broker-API-Version: 2.16" -H "Content-Type: application/json"
```

### Basic Auth

For the time being, basic auth has been disabled so that any `username:password` pair is valid.

## Contributing

### Generating API from spec

With the exception of `configure_hellosb.go` everything in the `restapi/` directory is generated with [go-swagger](https://github.com/go-swagger/go-swagger).

To generate new API resources:

```
swagger generate server -f /path/to/swagger.yaml -A hellosb --exclude-main --default-scheme http
```

## Versioning

`hellosb` tightly couples to Open Service Broker's [versioning scheme](https://github.com/openservicebrokerapi/servicebroker/blob/master/release-notes.md) for major and minor versions. Each tagged `hellosb` version is generated from the matching `servicebroker` swagger specification.

Patch versions correspond to bug fixes and dev iterations on `hellosb` itself.

For example: `2.16.4` will correlate to `hellosb`'s fourth revision implementing `servicebroker`'s tagged `2.16` specification.
