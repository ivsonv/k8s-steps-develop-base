## Terminal Commands 
* $ `kind` create cluster --config=configs/king.yml --name=clusterXXX
* $ `kubectl` cluster-info --context kind-clusterXXX
* $ `kubectl` get nodes
* $ `kubectl` config get-clusters
* $ `kubectl` config use-context=kind-clusterXXX 
* $ `kubectl` apply -f configs/pod.yml
* $ `kubectl` get pods
* $ `kubectl` port-forward pod/goserver 7000:8000
* $ `kubectl` apply -f configs/replicaset.yml
* $ `kubectl` get replicaset
* $ `kubectl` apply -f configs/deployment.yml
* $ `kubectl` get deployment
* $ `kubectl` delete pod ??
* $ `kubectl` delete replicaset ???
* $ `kubectl` describe pod ??name??
* $ `kubectl` describe deployment goserver
* $ `kubectl` rollout history deployment goserver (revisions)
* $ `kubectl` rollout undo deployment goserver (LastVersion)
* $ `kubectl` rollout undo deployment goserver --to-resision=??code revision??
* $ `kubectl` apply -f configs/service.yml 
* $ `kubectl` kubectl get services
* $ `kubectl` port-forward svc/goserver-service 7000:8000

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/7e896790-1ea4-4922-ba21-7fa638c68d7a)

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/967f1d12-8170-458d-a16e-8ba5fc235906)

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/16698637-24a0-4b4f-b441-fecbf20fe72c)

![image](https://github.com/ivsonv/k8s-steps-initials/assets/63156114/44ca2759-b34d-4464-b0c8-dc5a0215aca9)


