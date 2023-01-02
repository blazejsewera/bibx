BIBX_PKG=github.com/blazejsewera/bibx/cmd/bibx
BIBX_EXE=bibx

.PHONY: build build-linux test clean

build:
	@go build -o $(BIBX_EXE) $(BIBX_PKG)
	@echo "> binary ./$(BIBX_EXE) built"

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o $(BIBX_EXE) $(BIBX_PKG)
	@echo "> binary for linux ./$(BIBX_EXE) built"

test:
	@go test ./...

clean:
	@rm -f $(BIBX_EXE)
