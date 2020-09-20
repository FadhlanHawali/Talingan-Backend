build:
	@echo " >> building service"
	@go build -v -o v2/bin/talingan-backend v2/cmd/service/*.go

run: build
	./v2/bin/talingan-backend