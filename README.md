# kubernetes - K8S - API
```
* $`kubectl` proxy --port=8080
* http://localhost:8080/
* http://localhost:8080/api/v1/namespaces/default/services/XX

```

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/bacf85c1-0726-4a03-8289-4644a9cc8539)

# FORTIO - TEST STRESS & AUTO SCALE
- `kubectl` run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/healthz"
- `kubectl` get hpa
- `watch` -n1 `kubectl` get pods
- port default -> 80

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/a372bc7c-b595-4771-94fc-8d0b948cfef7)

# Terminal Commands
- $ `kubectl` api-resources (list recurses)
- $ `kind` create cluster --config=configs/king.yaml --name=clusterXXX
- $ `kubectl` cluster-info --context kind-clusterXXX
- $ `kubectl` get nodes, pods, replicaset, deployment, services, hpa, pvc, ns, serviceaccounts,
- $ `kubectl` config get-clusters, view, current-context, use-context
- $ `kubectl` config use-context=kind-clusterXXX 
- $ `kubectl` apply -f configs/XX.yaml (pod.yaml, replicaset.yaml, deployment.yaml, service.yaml, metrics.yaml, hpa.yaml)
- $ `kubectl` apply -f configs/XX.yaml -n=homolog (namespace)
- $ `kubectl` port-forward pod/goserver 7000:8000
- $ `kubectl` delete pod XX
- $ `kubectl` delete replicaset ???
- $ `kubectl` describe pod ??name??
- $ `kubectl` describe deployment goserver
- $ `kubectl` rollout history deployment goserver (revisions) -roolback version
- $ `kubectl` rollout undo deployment goserver (LastVersion) -roolback version
- $ `kubectl` rollout undo deployment goserver --to-resision=??code revision?? -roolback version
- $ `kubectl` port-forward svc/goserver-service 7000:8000
- $ `kubectl` top pod XXX
- $ `kubectl` create ns homolog
- $ `kubectl` exec -it #PODNAME# -- bash

# Certicate manager
Configuring TLS in the application, see the documentation here

- https://cert-manager.io/docs/installation/kubectl/

```
- kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.5/cert-manager.yaml
```
```
kubectl pods -n cert-manager
```
![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/7a596d37-6553-450c-a224-690430e9a136)

# screen terminal commands

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/7e896790-1ea4-4922-ba21-7fa638c68d7a)

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/967f1d12-8170-458d-a16e-8ba5fc235906)

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/16698637-24a0-4b4f-b441-fecbf20fe72c)

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/44ca2759-b34d-4464-b0c8-dc5a0215aca9)