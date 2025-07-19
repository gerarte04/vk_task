.PHONY: build_service launch_service

build_service:
	docker build -f Dockerfile -t marketplace:v1.0.0 .

launch_service:
	docker run -p 8080:8080 --env-file .env marketplace:v1.0.0