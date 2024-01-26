# Makefile

# Specify the path to the proto file
PROTO_FILE = proto/products/products.proto

# Specify the output directory for generated Go files
OUTPUT_DIR = ./gen/go/

# Define the protoc command with necessary options
PROTOC_COMMAND = protoc -I proto $(PROTO_FILE) \
    --go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
    --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative

# Default target, which will run the protoc command
all:
	$(PROTOC_COMMAND)
