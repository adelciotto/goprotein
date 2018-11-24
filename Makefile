.PHONY: build clean

BUILD_FLAGS := -v
PACK_MAIN_PKG := github.com/adelciotto/goprotein/cmd/pack
TRANSLATE_MAIN_PKG := github.com/adelciotto/goprotein/cmd/translate

default: build

build: build_pack build_translate

build_pack:
	go build $(BUILD_FLAGS) $(PACK_MAIN_PKG)

build_translate:
	go build $(BUILD_FLAGS) $(TRANSLATE_MAIN_PKG)

clean:
	go clean

