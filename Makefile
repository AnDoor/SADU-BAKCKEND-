run:
	go run ./cmd/main.go

mseed:
	go run ./seed/

nuke:
	rm -rf database.db
