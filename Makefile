APP_NAME=chrono-cli

build: clean
	go build -o dist/${APP_NAME} .

clean:
	rm -f dist/${APP_NAME}