# SECTION 6: CLUSTER MAINTENANCE

### OS Upgrades
* Pod Eviction Timeout: The maximum amount of time kubernetes will wait for a missing node to come back online.
* A safer way to take a node down for maintanence you can safely drain the pods off of the node you are upgrading which then moves them to other nodes in the cluster. Also node is marked as unscheduled.
  ```
  kubectl drain <node-name> --ignore
  ```
* To remove the unschedulable flag you can run the following command.
  ```
  kubectl uncordon <node-name>
  ```
* To block all pods from being scheduled on a node without draining that node can be done with the following cordon command.
  ```
  kubectl cordon <node-name>
  ```
### Kubernetes software versions

### Cluster Upgrade Process

### Backup and Restore Methods

### Working ETCDCTL