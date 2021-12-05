# Examples for the latest rpcx-plus

A lot of examples for [rpcx-plus](https://github.com/halokid/rpcx-plus)


## How to run
you should build rpcx-plus with necessary tags, otherwise only need to install rpcx-plus:

```sh
go get -u -v github.com/halokid/rpcx-plus/...
```

If you install succeefullly, you can run examples in this repository.

Enter one sub directory in this repository,  `go run server.go` in one terminal and `cd client; go run client.go` in another ternimal, and you can watch the run result.

For example,

```sh
cd 101basic
go run server.go
```

And

```sh
cd 101basic/client
go run client.go
```
