dump:
	go build -gcflags "-N -l" -o app
	# go tool objdump -s "main\.GenHeapProfile" app
	go tool objdump -s "pprof\.Lookup" app

test:
	GOOS=linux GOARCH=amd64 go build -gcflags "-N -l" -o app
	go test -v $^
