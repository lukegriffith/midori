.PHONY: test
test:
	go test ./...


.PHONY: install
install:
	go install cmd/journal/journal.go
