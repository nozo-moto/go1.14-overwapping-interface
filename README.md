# Go1.14-変更

Overlapping interafeceについて

https://github.com/golang/proposal/blob/master/design/6977-overlapping-interfaces.md#apendix-a-typical-example



# うごかしかた　
go1.14のダウンロード
```bash
GO111MODULE=off go get golang.org/dl/go1.14rc1
go1.14rc1 download
```

動かす
```bash
$ go1.14rc1 run main.go
a a impl b bimpl get
```


これがgo 1.13だと
```bash
$ go run main.go
# command-line-arguments
./main.go:22:5: duplicate method Get
```



