# rock-paper-scissors

> ref: https://dev.classmethod.jp/articles/golang-grpc-sample-project/

## execute
```go run ./cmd/api```

# fixed
in ```service/server.go```, `xtamp` is not declared.
So import `ts "google.golang.org/protobuf/types/known/timestamppb"` and  

```go
CreateTime: &ts.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
```