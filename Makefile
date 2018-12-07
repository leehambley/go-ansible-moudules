test-units:
	docker run --rm --network none --volume ${PWD}:/go/src/github.com/leehambley/go-ansible-modules golang:1.11.2-alpine3.7 go test -v github.com/leehambley/go-ansible-modules/...

test-integration:
	docker run --rm --network none --volume ${PWD}:/go/src/github.com/leehambley/go-ansible-modules golang:1.11.2-alpine3.7 go test -v -tags 'integration test' github.com/leehambley/go-ansible-modules/...

.PHONY: test-units test-integration