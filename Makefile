DIR ?= array_string

test:
	go test -v ./${DIR}/${FILE}.go ./${DIR}/${FILE}_test.go
