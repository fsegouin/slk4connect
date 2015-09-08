release:
	rsync -r src workspace
	cd workspace
	go get slk4connect
	go install slk4connect
