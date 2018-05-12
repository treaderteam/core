IOS_TREADER_PATH ?= gitlab.com/alexnikita/treader/reader/Reader.framework

build-ios:
	gomobile bind -target=ios -o $(IOS_TREADER_PATH)

get-ios:
	make saveremote && make untarize && make clean

save-ios:
	make build-ios && make tarize && make saveremote && make clean

install-deps:
	go get -d gitlab.com/alexnikita/gols/... && \
	go get -d github.com/golang/mobile/...

push:
	make save-ios && git push

tarize:
	tar -czvf 1.tar.gz Reader.framework

untarize:
	tar -xzvf 2.tar.gz -C $(IOS_TREADER_PATH)

saveremote:
	curl -X POST --data-binary @1.tar.gz http://188.166.94.143:8100/save

getremote:
	curl -X GET http://188.166.94.143:8100/get >> 2.tar.gz

clean:
	rm *.gz