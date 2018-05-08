IOS_TREADER_PATH ?= gitlab.com/alexnikita/treader/reader

build-ios:
	gomobile bind -target=ios $(IOS_TREADER_PATH)

install-deps:
	go get -d gitlab.com/alexnikita/gols/... && \
	go get -d github.com/golang/mobile/...

push:
	make build-ios && git push