# Go Bin Dump

Dump a go binaries's buildinfo to JSON format.

Same info as `go version -m <binary>`

## Usage

Install:
```sh
go install github.com/dcaravel/gobin-dump@main
```

Run:
```sh
gobin-dump <binary path>
```

Example:
```sh
gobin-dump $(which go)
```

```json
{
  "GoVersion": "go1.22.1",
  "Path": "cmd/go",
  "Main": {
    "Path": "",
    "Version": "",
    "Sum": "",
    "Replace": null
  },
  "Deps": null,
  "Settings": [
    {
      "Key": "-buildmode",
      "Value": "exe"
    },
    {
      "Key": "-compiler",
      "Value": "gc"
    },
    {
      "Key": "-gcflags",
      "Value": "cmd/...=-dwarf=false"
    },
    {
      "Key": "-trimpath",
      "Value": "true"
    },
    {
      "Key": "DefaultGODEBUG",
      "Value": "httplaxcontentlength=1,httpmuxgo121=1,panicnil=1,tls10server=1,tlsrsakex=1,tlsunsafeekm=1"
    },
    {
      "Key": "CGO_ENABLED",
      "Value": "0"
    },
    {
      "Key": "GOARCH",
      "Value": "arm64"
    },
    {
      "Key": "GOOS",
      "Value": "darwin"
    }
  ]
}
```