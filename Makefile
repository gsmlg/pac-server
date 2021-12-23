default: build

download:
	curl -sSLO https://raw.githubusercontent.com/petronny/gfwlist2pac/master/gfwlist.pac

build: generate
	@CGO_ENABLE=0 go build -o pac-server ./...

generate:
	@patch gfwlist.pac patch
	@go-bindata gfwlist.pac

unbound:
	@python3 gfwlist2unbound.py

