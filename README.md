<div align="center">
  <br>
  <img alt="Open Sauced" src="https://i.ibb.co/7jPXt0Z/logo1-92f1a87f.png" width="300px">
  <h1>🍕 Go API Client 🍕</h1>
  <strong>The generated Go OpenSauced API client</strong>
  <br>
</div>
<br>
<p align="center">
  <img src="https://img.shields.io/github/languages/code-size/open-sauced/go-api" alt="GitHub code size in bytes">
  <a href="https://github.com/open-sauced/go-api/issues">
    <img src="https://img.shields.io/github/issues/open-sauced/go-api" alt="GitHub issues">
  </a>
  <a href="https://github.com/open-sauced/go-api/releases">
    <img src="https://img.shields.io/github/v/release/open-sauced/go-api.svg?style=flat" alt="GitHub Release">
  </a>
  <a href="https://discord.gg/U2peSNf23P">
    <img src="https://img.shields.io/discord/714698561081704529.svg?label=&logo=discord&logoColor=ffffff&color=7389D8&labelColor=6A7EC2" alt="Discord">
  </a>
  <a href="https://twitter.com/saucedopen">
    <img src="https://img.shields.io/twitter/follow/saucedopen?label=Follow&style=social" alt="Twitter">
  </a>
</p>

This is a [generated OpenAPI client](https://openapi-generator.tech/)
for the [OpenSauced API](https://github.com/open-sauced/api)
which uses the [Swagger doc for the defined API](https://github.com/open-sauced/api/blob/beta/swagger.yaml).
It can be used as a canonical and versioned source to interface with the OpenSauced API using Go.

### 🔧 Versioning

The version of the client corresponds _directly_ withe the version of the API.
See the [OpenSauced API release notes](https://github.com/open-sauced/api/releases) for breaking changes, new features, etc.

### 🏗️ Usage

Add the client as a dependency in your project:

```
go get github.com/open-sauced/go-api/client
```

Here's a sample Go program that uses the client to get all pull requests

```go
package main

import (
        "context"
        "fmt"
        "os"

        client "github.com/open-sauced/go-api/client"
)

func main() {
    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)

    // The prod server is at index 1
    ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)

    resp, r, err := apiClient.PullRequestsServiceAPI.ListAllPullRequests(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsServiceAPI.ListAllPullRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    fmt.Printf("Responses: %v", resp)
}
```

Explore the generated docs within `client/docs` and `generated/README.md` to learn more about how this client works.

### 🚜 Generating the client

```bash
make generate-client
```

There are some rough edges when generating the OpenAPI client that you'll want to keep in mind:

- Duplicate `type` structs may be defined (especially for endpoints that utilize dependent Dtos)
  - Remove duplicate `type` structs that don't correspond to the right API endpoint.
    Always refer to the API's swagger doc for the canonical source of truth.
  - Also remove the corresponding documentation bits.
- The tests may not have correctly generated struct member func names.
  - Either remove those generated tests or find the correct struct member func and make the correct change.
- In general, it's a good idea to view the diff and run `go test ./...` before committing new changes.
- The "contributions" service, at the time of this writing, is not fully complete.
  So, it will have missing endpoints and no actual services. Its files can be safely deleted.
