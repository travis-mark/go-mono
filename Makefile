hello: cmd/hello.go
	go run cmd/hello.go

xmas: cmd/xmas.go
	go run cmd/xmas.go

install: cmd/*
	go install cmd/hello.go
	go install cmd/xmas.go