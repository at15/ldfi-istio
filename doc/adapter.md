# Mixer Adapter

## TODO

- [ ] update istio version

## Write a new adapter

```bash
mkdir -p $GOPATH/src/istio.io/ && \
cd $GOPATH/src/istio.io/  && \
git clone https://github.com/istio/istio

export MIXER_REPO=$GOPATH/src/istio.io/istio/mixer
export ISTIO=$GOPATH/src/istio.io
cd $ISTIO/istio && make mixs
```

## Reference

- https://github.com/istio/istio/wiki/Mixer-Out-Of-Process-Adapter-Dev-Guide
- https://github.com/istio/istio/wiki/Mixer-Out-Of-Process-Adapter-Walkthrough