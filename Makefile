build: build-server install-web build-web

build-api:
	go build -ldflags "-s"

install-web:
	# cd web && npm install

build-web:
	# cd web && npm run build

