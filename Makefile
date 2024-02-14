NAME			=	tiny-btc

BUILD_DIR       =	./build

CMD_DIR			=	./cmd/tinyBTC

GO				=	go

GO_BUILD		=	$(GO) build

GO_MOD			=	$(GO) mod

GO_GET			=	$(GO) get

GO_CLEAN		=	$(GO) clean

GO_TEST			=	$(GO) test

MKDIR			=	mkdir -p

CP				=	cp

RM				=	rm -f

SH				=	sh

all				:	build

build			:
					$(MKDIR) $(BUILD_DIR)
					$(GO_BUILD) -o $(BUILD_DIR)/$(NAME) -v $(CMD_DIR)/main.go

run				:
					$(BUILD_DIR)/$(NAME)

test			:
					$(GO_TEST) -v ./...

vet 			:
					$(GO) vet ./...

clean			:
					$(GO_CLEAN)
					$(RM) $(BUILD_DIR)/$(NAME)

install			:
					$(GO_MOD) download

update			:
					$(GO_CLEAN)
					$(GO_GET) -u ./...
					$(GO_MOD) tidy
					$(GO_MOD) vendor

vendor			:
					$(GO_MOD) vendor

rebuild			:	clean build

re				:	rebuild

.PHONY			:	all make build test vet clean install update generate vendor rebuild re