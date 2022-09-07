build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ./tmp/app main.go
	rm -rf ./tmp/app.zip
	zip ./tmp/app.zip ./tmp/app