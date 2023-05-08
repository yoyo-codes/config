# Credits: https://github.com/ardanlabs/service/blob/dad2b38f35da5a1383c2690f6a37de91384e5047/makefile#L385

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all
