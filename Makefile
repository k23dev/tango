BINARY_NAME=app
BINARY_NAME_WIN=app.exe
BUILD_DIR=./build
BUILD_DIR_WIN=./build/windows
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
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -gccgoflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR}

build-win:
	# create directories
	mkdir -p ${BUILD_DIR_WIN}
	mkdir ${BUILD_DIR_WIN}/_db
	mkdir ${BUILD_DIR_WIN}/logs
	mkdir ${BUILD_DIR_WIN}/public

	# copy the config files and seeds
	cp -r ./config ./build/config
	cp -r ./seeds ./build/seeds
	
	# templates generator
	make templates

	# compile into binary file
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME_WIN} -gccgoflags "-w -s"
	chmod +x ${BINARY_NAME_WIN}
	mv ${BINARY_NAME_WIN} ${BUILD_DIR_WIN}

templates:
	templ generate

templates-clean:
	rm -f ${VIEWS_DIR}components/*.go
	rm -f ${VIEWS_DIR}layouts/*.go
	rm -f ${VIEWS_DIR}forms/*.go
	rm -f ${VIEWS_DIR}tables/*.go
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
	rm -f ${HTMLX_DIR}/htmx.min.js
	wget https://unpkg.com/htmx.org/dist/htmx.min.js -P ${HTMLX_DIR}

htmlx-update:
	rm -f ${HTMLX_DIR}/htmx.min.js
	wget https://unpkg.com/htmx.org/dist/htmx.min.js -P ${HTMLX_DIR}

tailwind-install:
	npm install -D tailwindcss
	npx tailwindcss init

tailwind-dev:
	npx tailwindcss -i ./app/views/tailwind.css -o ./public/assets/css/app.css --watch

tailwind-build:
	npx tailwindcss -i ./app/views/tailwind.css -o ./public/assets/css/app.css --minify

cli-install:
	git clone https://github.com/k23dev/tango_cli
	cd tango_cli
	GOOS=linux GOARCH=amd64 go build -o tango_cli -gccgoflags "-w -s"
	chmod +x tango_cli
	cp tango_cli ../
