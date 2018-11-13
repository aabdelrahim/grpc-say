build:
	GOOS=linux go build -o app
	docker build -t gcr.io/grpc-say/say .
	rm -f app

push:
	gcloud docker -- push gcr.io/grpc-say/say