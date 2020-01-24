PKGS = ./ldfiadapter/... ./cmd/...
PKGST = ./ldfiadapter ./cmd

.PHONY: install
install:
	go install ./cmd/ldfiadapter

.PHONY: fmt
fmt:
	goimports -d -l -w $(PKGST)