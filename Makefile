swagger:
	swag init -g .\cmd\api.go --output .\cmd\docs

build-api:
	go build -o .\cmd\bin\AthenaAPI.exe .\cmd\api.go