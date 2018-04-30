build-ios:
	gomobile bind -target=ios gitlab.com/alexnikita/treader/reader

install-deps:
	go get -d gitlab.com/alexnikita/gols/...

push:
	make build-ios && git push