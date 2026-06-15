gen-update-package:
	go get github.com/magic-lib/go-plat-trace@master
	go get github.com/magic-lib/go-plat-cache@master
	go get github.com/magic-lib/go-plat-curl@master
	go get github.com/magic-lib/go-plat-mysql@master
	go get github.com/magic-lib/go-plat-retry@master
	go get github.com/magic-lib/go-plat-utils@master
	go mod tidy

mac:
	GOOS=darwin go build -ldflags="-s -w" -o ${GOPATH}/bin/create-secret ./main/create_secret.go
	$(if $(shell command -v upx || which upx), upx embed-darwin)

linux:
	GOOS=linux go build -ldflags="-s -w" -o ${GOPATH}/bin/create-secret-linux ./main/create_secret.go
	if command -v upx >/dev/null 2>&1; then \
    	upx $(GOPATH)/bin/create-secret-linux; \
    fi