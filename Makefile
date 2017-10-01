image = redirector

compile:
	go build .

compile-arm:
	GOOS=linux GOARCH=arm GOARM=5 go build .

build:
	docker build -t $(image) .

deploy:
	docker run -d -p 80:8080 $(image)
