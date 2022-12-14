# SECTION 9: Networking

## Prerequisite Learning
#### Switching Routing:
#### DNS:
#### CoreDNS:
#### Network Namepaces:
#### Docker Networking:
#### CNI:

## Cluster Networking
## Pod Networking
## CNI in Kubernetes
* Installing a CNI in Kubernetes: https://v1-22.docs.kubernetes.io/docs/setup/production-environment/tools/kubeadm/high-availability/#steps-for-the-first-control-plane-node
## CNI Weave
* Deployed as a daemonset
* Deploy weave cni plugin
  ```
  kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version |base64 | tr -d '\n')"
  ```
* Check logs for weave pods(in kubeadm setup)
  ```
  kubectl logs weave-net-5gmcb weave -n kube-system
  ```
## IP Address Management - Weave

## Service Networking
## DNS in Kubernetes
## CoreDNS in Kubernetes
## Ingress