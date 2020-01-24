# Hipster

https://github.com/GoogleCloudPlatform/microservices-demo

```bash
kubectl apply -f ./hipster/kubernetes-manifests.yaml
minikube tunnel
```

```bash
kubectl get services | grep frontend
frontend                ClusterIP      10.96.105.34    <none>         80/TCP         12m
frontend-external       LoadBalancer   10.96.54.223    10.96.54.223   80:31609/TCP   12m
```

Clean up

## Services

- https://github.com/GoogleCloudPlatform/microservices-demo/blob/master/kubernetes-manifests/checkoutservice.yaml
  - reach other services using service name

## Tracing

```bash
istioctl dashboard jaeger
```

- https://github.com/GoogleCloudPlatform/microservices-demo/issues/112
  - only some of them enabled open tracing
  - it is using `jaeger-collector:14268` which should be same as default istio demo
  - https://www.jaegertracing.io/docs/1.16/deployment/ explains the ports
- [ ] it has `adservice` and `recommendationservice`