```
protoc -I=. --go_out=. point.proto
```

```
protoc --go_out=plugins=grpc:. point.proto
```
