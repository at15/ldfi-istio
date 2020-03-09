# 2020-03-09 Envoy

https://github.com/proxy-wasm/proxy-wasm-cpp-sdk

## TODO

- [ ] update the upstream doc, it seems `PROXY_WASM_CPP_SDK` should be replaced with `DOCKER_SDK` in `Makefile.base_lite`

```bash
docker build -t wasmsdk:v2 -f Dockerfile-sdk .
```

didn't work so I just manually copy content out from `Makefile.base_lite` and `build_wasm.sh`

- run the bash script manually, then `make myproject.wasm` and I got 

```bash
source /root/emsdk/emsdk_env.sh
export PATH=/usr/local/bin:$PATH
```
