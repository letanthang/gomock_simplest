add mockgen to go.mod
```
go get github.com/golang/mock/mockgen/model
```
Generate code
```
mockgen -destination=mocks/mock_foo.go -package=mocks . Foo
```
Run the test
```test
go test 
```
