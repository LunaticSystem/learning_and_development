# SECTION 13: Troubleshooting

## Application Failure
* Frontend
  * Check if the application is available on the nodeport
    ```
    curl http://web-service-ip:node-port
    ```
  * Check the service for the web service to make sure its exposing the correct ports.
    ```
    kubectl describe service web-service
    ```
  * If not make sure the selectors are correct between the configured service and its manifest.
  * Check the pod and whether its in a running state.
    ```

    ```
## Control Plane Failure
* Check status of nodes in cluster
* Check status of pods in the cluster
* Check logs of control plane component pods
  ```
  kubectl logs kube-apiserver -n kube-system
  ```
* Check logs fo control plane component services
  ```
  sudo journalctl -u kube-apiserver 
  ```
## Worker Node Failure
* Check status of node in cluster
  ```
  kubectl get nodes
  ```
* If not ready use decribe on the node.
  ```
  kubectl describe node worker-1
  ```
* Check the status of node.
  ```

  ```
* Check the status and kubelet logs for more information.
* Check the certificates on the kubelet.
## Network Troubleshooting