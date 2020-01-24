# istio


```bash
istioctl manifest apply --set profile=demo
kubectl get svc -n istio-system
kubectl label namespace default istio-injection=enabled
```

## Pods

- istio-tracing is `docker.io/jaegertracing/all-in-one:1.14`