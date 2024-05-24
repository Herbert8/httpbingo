
release: all_doc compile

release_rebuild: all_doc build

prepare_doc: gen_redoc gen_swagger gen_rapidoc

# 将 OpenAPI 的 JSON 转为 YAML
convert_json_yaml:
	yq eval -P < api/openapi-spec/HTTPBinGo.openapi.json > api/openapi-spec/HTTPBinGo.openapi.yaml

# 生成 ReDoc
gen_redoc: convert_json_yaml
	npx @redocly/cli build-docs api/openapi-spec/HTTPBinGo.openapi.yaml -o cmd/httpbingo/doc/html/redoc.html
#	redocly build-docs api/openapi-spec/HTTPBinGo.openapi.yaml -o cmd/httpbingo/doc/html/redoc.html
#	open -R cmd/httpbingo/doc/html/redoc.html

# 生成 Swagger 使用手册
gen_swagger:
	cat api/HTTPBinGo.html | sed 's|<title>Apifox 接口文档</title>|<title>HTTPBinGo 使用说明</title>|g' > cmd/httpbingo/doc/html/swagger.html

# 生成 RapiDoc
gen_rapidoc:
	specs_dir=cmd/httpbingo/doc/html/rapidoc/specs ; mkdir -p "$$specs_dir" ; \
		cp api/openapi-spec/HTTPBinGo.openapi.yaml "$$specs_dir/temp.yaml"

# 生成所有格式文档
all_doc: prepare_doc convert_json_yaml gen_redoc gen_swagger gen_rapidoc

gen_version_info:
	bash scripts/version_info.sh

# 构建
build: gen_version_info
	bash scripts/build.sh httpbingo ./cmd/httpbingo -a -v

# 编译
compile: gen_version_info
	bash scripts/build.sh httpbingo ./cmd/httpbingo -v

.PHONY: build compile
