# go-flagx

this package is flag extentions. for me

## Usage

- `StringSlice`

```go
// command -name taro -name hanako -name jiro
var names = flagx.StringSlice([]string)
flag.Var(&names, "name", "name list")
flag.Parse()

// names = []string{"taro", "hanako", "jiro"}
```

- `EnvToFlagWithPrefix`

```go
// APP_NAME=foo APP_PORT=5000 command
var (
    name string
    port int
)
flag.StringVar(&name, "name", "", "set name") 
flag.IntVar(&port, "port", 8080, "listen port")
flag.VisitAll(flagx.EnvToFlagWithPrefix("APP_"))
flag.Parse()

// name = "foo", port = 5000
```
