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
* Kinds of services
  * ClusterIP: Service gives cluster wide access to all pods in the cluster.
  * NodePort: Exposes application on all nodes in the cluster to external clients as well as pods within the cluster.
* Virtual Object.
* Assigned an IP address from a predefined block of IPs.
* Kube proxy components gets the IP's and creates forwarding rules in iptables. 
* IP/Port Combination.
* Can gather the range of IPs available by running the below command. You will be looking for the `range` option in the command
  ```
  ps aux |grep kube-api-server
  ```
* Gathering the nat rules for services in the cluster.
  ```
  iptables -L -t nat |grep <service_name>
  ```
* You can also trace the rule creation by kube proxy in the kube-proxy.log file.
  ```
  cat /var/log/kube-proxy.log
  ```
* How to find what proxier kube-proxy is using.
  ```
  kubectl logs kube-proxy-dl5pq -n kube-system |grep Proxier
  ```
## DNS in Kubernetes
* K8s deploys a DNS server automatically.
* If installing manually you install yourself.
* DNS record example for a service.
  ```
  |Hostname    |Namespace|Type|Root          |IP Address   |
   web-service  apps      svc  cluster.local  10.107.37.188
  ```
* DNS record example for a pod.
  ```
  |Hostname    |Namespace|Type|Root          |IP Address   |
   10-244-2-5   apps      pod  cluster.local  10.244.2.5
  ```
## CoreDNS in Kubernetes
* Deployed as a pod in the kube-system namespace.
* K8s uses a core file located at /etc/coredns/Corefile which is the configuration for the Core DNS server deployed in the cluster.
* Corefile is passed in using a configmap object.
* /etc/resolve.conf on pods is updated by the kubelet automatically on creation based on the configuration set in /var/lib/kubelet/config.yaml
  ```
  clusterDNS:
  - 10.x.x.x
  clusterDomain: cluster.local
  ```
## Ingress
* Layer 7 load balancer.
* Either provision with a nodPort or cloud based load balancer.
* First deploy a support solution
* Then specify a set of rules.
* Ingress controller nginx, haproxy, traefik
* Ingress resources = configuration
* Not deployed by default.
* GCE and nginx are being maintained by the k8s project.

### Deploying ingress controller
* Deployment definition file.
  ```yaml
  apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    name: nginx-ingress-controller
  spec:
    replicas: 1
    selector:
      matchLabels:
        name: nginx-ingress
    template:
      metadata:
        labels:
          name: nginx-ingress
      spec:
        containers:
          - name: nginx-ingress-controller
            image: quay.io/kubernetes-ingress-controller/nginx-ingresss-controller:0.21.0
        args:
          - /nginx-ingress-controller
          - --configmap=$(POD_NAMESPCE)/nginx-configuration
        env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                fieldPath: metadata.namespace
        ports:
          - name: http
            containerPort: 80
          - name: https
            containerPort: 443
  ```
  * Deploy config map for nginx that includes `err-log-path`, `keep-alive`, and `ssl-protocols`
    ```yaml
    kind: ConfigMap
    apiVersion: v1
    metadata:
      name: nginx-configuration
    ```
  * Create a service for the ingress controller.
    ```yaml
    apiVersion: v1
    kind: Service
    metadata:
      name: nginx-ingress
    spec:
      type: NodePort
      ports:
      - port: 80
        targetPort: 80
        protocol: TCP
        name: http
      - port: 443
        targetPort: 443
        protocol: TCP
        name: https
      selector:
        name: nginx-ingress
    ```
  * Create a service account with the appropriate roles, cluster roles, rolebindings.

### Create ingress resources/configs for services.
* Create an ingress object via manifest.
  * Path directly to serivce (Basic)
    ```yaml
    apiVersion: extensions/v1beta1
    kind: Ingress
    metadata:
      name: ingress-wear
    spec:
      backend:
        serviceName: wear-service
        servicePort: 80
    ```
  * More complex setup with mutiple paths.
    ```yaml
    apiVersion: extensions/v1beta1
    kind: Ingress
    metadata:
      name: ingress-wear-watch
      annotations:
        nginx.ingress.kubernetes.io/rewrite-target: /
        nginx.ingress.kubernetes.io/ssl-redirect: "false"
    spec:
      rules:
      - http:
          paths:
          - path: /wear
            backend:
              serviceName: weare-service
              servicePort: 80
          - path: /watch
            backend:
              serviceName: watch-service
              servicePort: 80
    ```
  * How to configure a default back
  * How to create an ingress with mutiple domains.
    ```yaml
    apiVersion: extensions/v1beta1
    kind: Ingress
    metadata:
      name: ingress-wear-watch
    spec:
      rules:
      - host: wear.my-online-store.com
        http:
          paths:
          - backend:
              serviceName: weare-service
              servicePort: 80
      - host: watch.my-online-store.com
        http:
          paths:
          - backend:
              serviceName: watch-service
              servicePort: 80
    ```

  [Section 10: Design & Install Cluster](https://github.com/LunaticSystem/learning_and_development/tree/main/certified_k8s_administrator_cert/course_sections/section10_design_install_cluster)