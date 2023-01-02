BIBX_PKG=github.com/blazejsewera/bibx/cmd/bibx
BIBX_EXE=bibx

.PHONY: build test clean

build:
	@go build -o $(BIBX_EXE) $(BIBX_PKG)
	@echo "> binary ./$(BIBX_EXE) built"

test:
	@go test ./...

clean:
	@rm -f $(BIBX_EXE)
