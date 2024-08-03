P ?= array_string

test:
	go test -v ./${P}/${F}.go ./${P}/${F}_test.go
