
release: all_doc compile

release_rebuild: all_doc build

prepare_doc: gen_redoc gen_swagger gen_rapidoc

convert_json_yaml:
	yq eval -P < api/openapi-spec/HTTPBinGo.openapi.json > api/openapi-spec/HTTPBinGo.openapi.yaml

gen_redoc: convert_json_yaml
	npx @redocly/cli build-docs api/openapi-spec/HTTPBinGo.openapi.yaml -o cmd/httpbingo/doc/html/redoc.html
#	redocly build-docs api/openapi-spec/HTTPBinGo.openapi.yaml -o cmd/httpbingo/doc/html/redoc.html
#	open -R cmd/httpbingo/doc/html/redoc.html

gen_swagger:
	cat api/HTTPBinGo.html | sed 's|<title>Apifox 接口文档</title>|<title>HTTPBinGo 使用说明</title>|g' > cmd/httpbingo/doc/html/swagger.html

gen_rapidoc:
	cp api/openapi-spec/HTTPBinGo.openapi.yaml cmd/httpbingo/doc/html/rapidoc/specs/temp.yaml

all_doc: prepare_doc convert_json_yaml gen_redoc gen_swagger gen_rapidoc

build:
	bash scripts/build.sh httpbingo ./cmd/httpbingo -a -v

compile:
	bash scripts/build.sh httpbingo ./cmd/httpbingo -v

.PHONY: build compile