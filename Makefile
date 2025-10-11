.PHONY: generate
generate:
	go generate ./...

.PHONY: oapi-gen
oapi-gen:
	oapi-codegen --config=config.yaml ./shared/api/order/api.yaml