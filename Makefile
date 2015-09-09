release:
	rsync -r src workspace
	cd workspace
	go get connect4
	go test connect4
