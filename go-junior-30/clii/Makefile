BIN=clii

.PHONY:inst build-win build-mac-intel build-mac-arm
.SILENT: inst
inst:
	go build .
	sudo mv $(BIN) /usr/local/bin 