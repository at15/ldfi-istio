# ldfi-istio

A istio mixer adapter for LDFI (Lineage-driven Fault Injection)

## Overview

This is a work in progress (i.e. README only) repo for [ldfi-adapter](https://github.com/disorderlylabs/ldfi-adapter) and aims to address the following (engineering) problems

- Dynamic configuration of fault injection rule for the mixer adapter, it is [relatively static](https://github.com/disorderlylabs/ldfi-adapter/blob/master/mygrpcadapter.go#L102-L145) and can't be used in a feedback loop
- Apply to a more complex service, e.g. [Hipster Shop](https://github.com/GoogleCloudPlatform/microservices-demo) and runs on cloud providers.

## TODO

Env

- [ ] instruction for minikube
- [ ] instruction for GKE

Adapter

- [ ] provide API for updating failure conditions in adapter, CRD might be the easiest way.
- [ ] see if it is possible to run the mixer as an actual k8s operator, this allows the adapter to have the actual view of entire cluster and even fail node directly, similar to [pingcap/chaos-mesh](https://github.com/pingcap/chaos-mesh)

LDFI

- [ ] see how LDFI works w/ trace from jaeger and modeling the system using datalog

## Related

- [ldfi-adapter](https://github.com/disorderlylabs/ldfi-adapter)
- [at15/bop-bpf](https://github.com/at15/bop-bpf)

## Reference

### Fault injection on k8s

- [pingcap/chaos-mesh](https://github.com/pingcap/chaos-mesh) A Chaos Engineering Platform for Kubernetes