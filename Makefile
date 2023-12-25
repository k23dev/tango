BINARY_NAME=app
BUILD_DIR=./build
HTMLX_DIR=./public/assets/js
VIEWS_DIR=./app/views/

tango-install:
	mkdir _db
	make go-install-deps
	make tailwind-install
	make htmlx-install
	make templates
	go mod tidy

go-install-deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go mod tidy

build:
	# create directories
	mkdir -p ${BUILD_DIR}
	mkdir ${BUILD_DIR}/_db
	mkdir ${BUILD_DIR}/logs
	mkdir ${BUILD_DIR}/public

	# copy the config files and seeds
	cp -r ./config ./build/config
	cp -r ./seeds ./build/seeds
	
	# templates generator
	make templates

	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -ldflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR}

templates:
	templ generate

templates-clean:
	rm -f ${VIEWS_DIR}components/*.go
	rm -f ${VIEWS_DIR}layouts/*.go
	rm -f ${VIEWS_DIR}menus/*.go
	rm -f ${VIEWS_DIR}*.go

dev:
	# make templates-clean
	make templates
	go run .

dev-air:
	make templates-clean
	make templates
	air

run:
	make templates
	go run .

test:
	go test ./tests

clean:
	go clean
	rm -rf ${BUILD_DIR}

htmlx-install:
	wget https://unpkg.com/htmx.org/dist/htmx.min.js -P ${HTMLX_DIR}

htmlx-update:
	rm -rf ${HTMLX_DIR}/*
	wget https://unpkg.com/htmx.org/dist/htmx.min.js -P ${HTMLX_DIR}

tailwind-install:
	npm install -D tailwindcss
	npx tailwindcss init

tailwind-dev:
	npx tailwindcss -i ./app/views/tailwind.css -o ./public/assets/css/app.css --watch

tailwind-build:
	npx tailwindcss -i ./app/views/tailwind.css -o ./public/assets/css/app.css --minify