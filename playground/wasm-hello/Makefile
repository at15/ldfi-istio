docker-build:
	docker run -v $(PWD)/code:/work -w /work  wasmsdk:v2 /build_wasm.sh

docker-exec:
	docker run -it -v $(PWD)/code:/work -w /work  wasmsdk:v2 /bin/bash