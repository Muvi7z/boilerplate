.PHONY: generate
generate:
	go generate ./...

.PHONY: oapi-gen
oapi-gen:
	oapi-codegen --config=config.yaml ./shared/api/order/api.yaml

docker:
	docker build -f docker/inventory/Dockerfile -t boilerplate-service .