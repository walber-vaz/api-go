.PHONY: run clean test build

C_RED=\033[0;31m
C_GREEN=\033[0;32m
C_YELLOW=\033[0;33m
C_BLUE=\033[0;34m
C_PURPLE=\033[0;35m
C_CYAN=\033[0;36m
C_WHITE=\033[0;37m
C_RESET=\033[0m

default: run

run:
	@echo "$(C_CYAN)[RUN-DEV]$(C_RESET)$(C_WHITE)Running server dev in go...$(C_RESET)"
	@go run main.go

clean:
	@echo "$(C_RED)[CLEAR]$(C_RESET)$(C_WHITE)Cleaning up...$(C_RESET)"
	@rm -rf *.so *.test *.out *.log *.db *.zip apijobs

test:
	@echo "$(C_YELLOW)[TEST]$(C_RESET)$(C_WHITE)Running tests...$(C_RESET)"
	@go test -v ./...

build:
	@echo "$(C_PURPLE)[BUILD]$(C_RESET)$(C_WHITE)Building dev air...$(C_RESET)"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -installsuffix cgo -o apijobs ./main.go