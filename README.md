# Winnow

Winnow is a small yet powerful permissions library for go.

## Install
```
go get https://github.com/UnwrittenFun/winnow
```

## How to use

### Load grants
Load grants from somewhere (e.g. disk, database, etc.)
```go
var win Winnow
if err := json.Unmarshal(dataFromStorage, &win); err != nil {
    panic(err)
}
```

Example permissions file:
```json
{
    "grants": [
        {
            "collection": "posts",
            "match": {
                "public": true
            },
            "operations": ["read"]
        },
        {
            "collection": "posts",
            "match": {
                "authorId": 12345
            },
            "operations": ["create", "read", "update", "delete"]
        }
    ]
}
```

Alternatively you can define grants in code:
```go
win := winnow.Winnow{
    Grants: []winnow.Grant {
        {
            Collection: "posts",
            Match: map[string]interface{}{
                "public": true,
            },
            Operations: []string{"read"},
        },
    },
}
```

### Check for permission

```go
canRead := win.Can("read", "posts", map[string]interface{"public": true})
```
