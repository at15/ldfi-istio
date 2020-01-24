# Hipster

https://github.com/GoogleCloudPlatform/microservices-demo

```bash
kubectl apply -f ./release/kubernetes-manifests.yaml
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

- https://github.com/GoogleCloudPlatform/microservices-demo/issues/112
  - only some of them enabled open tracing