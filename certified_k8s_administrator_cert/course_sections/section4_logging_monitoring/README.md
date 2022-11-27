# SECTION 4: LOGGING/MONITORING

## Monitoring Cluster Components:
* Monitoring Solutions:
  * Metrics Server - Basic
    * One per kubernetes cluster
    * Aggregates metrics in memory
    * No historical data
  * Prometheus - Advanced
  * Elastic Stack - Advanced
  * DataDog - Advanced
  * Dynatrace - Advanced
* Kubelet
  * cAdvisor

* To deploy metric server
  * Minikube
    ```
    minikube addons enable metrics-server
    ```
  * All other k8s versions"
    ```
    git clone https://github.com/kubernetes-incubator/metrics-server

    kubectl create -f deploy/1.8+/
    ```
* To view metrics from metrics server for nodes
  ```
  kubectl top node
  ```
* To view metrics from metrics server for pods
  ```
  kubectl top pod
  ```
## Managing Application Logs:
* Viewing logs for all containers in pod
  ```
  kubectl logs -f <pod_name>
  ```
* Viewing logs for a specific container in a pod
  ```
  kubectl logs -f <pod_name> <container_name>
  ```
