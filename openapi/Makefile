default:

client:
	@echo "Generating OpenAPI Go Client"
	@docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.7.0 generate \
		-i /local/easydns.yaml \
		-g go \
		-o /local/client/go
