#Makefile

# Specify the path to the proto file
PROTO_PRODUCTS = proto/products/products.proto
PROTO_IMAGES = proto/images/images.proto
# Specify the output directory for generated Go files
OUTPUT_DIR = ./gen/go/

# Define the protoc command with necessary options
PROTOC_PRODUCTS = protoc -I proto $(PROTO_PRODUCTS) \
    --go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
    --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative

PROTOC_IMAGES = protoc -I proto $(PROTO_IMAGES) \
	--go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative

# Default target, which will run the protoc command

images:
	$(PROTOC_IMAGES)
products:
	$(PROTOC_PRODUCTS)