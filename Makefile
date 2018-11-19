.PHONY: build clean

PACK_BUILD_FLAGS := -v
PACK_MAIN_PKG := github.com/adelciotto/goprotein/cmd/pack

default: build_pack

build_pack:
	go build $(PACK_BUILD_FLAGS) $(PACK_MAIN_PKG)

clean:
	go clean

