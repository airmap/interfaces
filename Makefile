PROTOS_DIR := ./grpc
PROTO_FILES := $(shell find $(PROTOS_DIR) -name '*.proto')
SOURCE_DIR := ./src

all: golang

golang:
	@rm -rf $(SOURCE_DIR)/go
	@mkdir -p $(SOURCE_DIR)/go
	@for x in $(PROTO_FILES); do \
		protoc -I$(PROTOS_DIR) --go_out=plugins=grpc,paths=source_relative:$(SOURCE_DIR)/go $$x; \
	 done
