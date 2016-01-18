.PHONY: build-arm6 build-arm7

GO_BUILD_FLAGS=-v -ldflags '-w -s' -a -tags netgo -installsuffix netgo

build-arm6::
	@GOOS=linux GOARCH=arm GOARM=6 go build -v $(GO_BUILD_FLAGS) -o rpi-board-info main.go

build-arm7::
	@GOOS=linux GOARCH=arm GOARM=7 go build -v $(GO_BUILD_FLAGS) -o rpi-board-info main.go
