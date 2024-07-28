P ?= array_string

test:
	go test ./${P}/${F}.go ./${P}/${F}_test.go