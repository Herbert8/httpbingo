
prepare_doc: gen_redoc gen_swagger gen_rapidoc

convert_json_yaml:
	yq eval -P < api/openapi-spec/HTTPBinGo.openapi.json > api/openapi-spec/HTTPBinGo.openapi.yaml

gen_redoc: convert_json_yaml
	npx @redocly/cli build-docs api/openapi-spec/HTTPBinGo.openapi.yaml -o cmd/httpbingo/doc/html/redoc.html
	open -R cmd/httpbingo/doc/html/redoc.html

gen_swagger:
	cp api/HTTPBinGo.html cmd/httpbingo/doc/html/swagger.html

gen_rapidoc:
	cp api/openapi-spec/HTTPBinGo.openapi.yaml cmd/httpbingo/doc/html/rapidoc/specs/temp.yaml