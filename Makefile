dependencies:
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/onsi/ginkgo/ginkgo
	go get

run-locally:
	go run main.go

swagger:
	$(HOME)/go/bin/swag init
