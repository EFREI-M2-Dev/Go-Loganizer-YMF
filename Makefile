PROJECT_NAME = loganizer
SRC_DIR = .
OUT_DIR = out
SRC_FILE = main.go
GOFLAGS = -v

all: build

build: $(OUT_DIR)/$(PROJECT_NAME)

$(OUT_DIR)/$(PROJECT_NAME):
	mkdir -p $(OUT_DIR)
	go build $(GOFLAGS) -o $@ $(SRC_DIR)/$(SRC_FILE)

clean:
	rm -rf $(OUT_DIR)

run: build
	@if [ -z "$(ARGUMENTS)" ]; then \
		echo "Usage: make run ARGUMENTS='<command> [options]'"; \
		exit 1; \
	fi
	$(OUT_DIR)/$(PROJECT_NAME) $(ARGUMENTS)

test:
	go test $(SRC_DIR)/...

.PHONY: all build clean run install test
