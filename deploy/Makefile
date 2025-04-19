BINARY_NAME=dummy
ADMIN_BINARY_NAME=dummy-admin
COMPLETION_DIR=/usr/local/share/dummy-completion/completions

build:
	go build -o bin/$(BINARY_NAME) ../cmd/$(BINARY_NAME)
	go build -o bin/$(ADMIN_BINARY_NAME) ../cmd/$(ADMIN_BINARY_NAME)

lint:
	bash ./scripts/lint.sh

install:
	# Ensure the bin directory exists
	mkdir -p bin
	# Build the binaries
	make build
	# Install the binaries to /usr/local/bin
	cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	cp bin/$(ADMIN_BINARY_NAME) /usr/local/bin/$(ADMIN_BINARY_NAME)
	echo "✅ Installed $(BINARY_NAME) and $(ADMIN_BINARY_NAME) to /usr/local/bin."

completion:
	# Generate shell completion scripts for both binaries
	./bin/$(BINARY_NAME) completion > $(COMPLETION_DIR)/$(BINARY_NAME)
	./bin/$(ADMIN_BINARY_NAME) completion > $(COMPLETION_DIR)/$(ADMIN_BINARY_NAME)

install-completion:
	# Ensure the completion directory exists
	mkdir -p $(COMPLETION_DIR)
	# Generate and install completion scripts
	make completion
	echo "✅ Completion scripts installed to $(COMPLETION_DIR)."