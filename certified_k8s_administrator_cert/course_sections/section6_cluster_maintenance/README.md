# SECTION 6: CLUSTER MAINTENANCE

## OS Upgrades
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
## Kubernetes software versions

## Cluster Upgrade Process
* Not manditory for all components 
* No other component should be at a version higher than the kubeapi server
* Controller manager and scheduler can be at one version lower then kube api server
* Kubelet and kube proxy can be at two versions lower.
* Kubectl can be at one version higher or lower than kube api server.
* Kubernetes supports 3 prior minor version releases. i.e 1.12, 1.11, 1.10
* Recommended to upgrade one minor version at a time.
* If your cluster was deployed by:
  * A cloud provider - Then use the simple upgrade method provided by the cloud provider
  * Kubeadm - Use the `kubeadm upgrade plan` and `kubeadm upgrade apply` commands which will help assist you in the upgrade.
  * Manually - You will need to upgrade each component manually.

## Kubeadm Upgrade process
Upgrades are comprised of two main steps:<br />
* Upgrading the Master
* Upgrading the Worker nodes
  * Upgrade all at once - however in this case the pods are down and no one can access the applications. Requires downtime.
  * Upgrade one node at a time - zero downtime upgrade
  * Add new nodes to the cluster with newer software version. Especially convenient in a cloud env.

Kubeadm does not upgrade kubelets.

Process (https://v1-24.docs.kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/):<br />
[Master]
* Upgrade kubeadm to the same minor build of k8s you are upgrading too.
  ```
  apt-get upgrade -y kubeadm=1.12.0-00
  ```
* Run the kubeadm tool and apply the upgrade.
  ```
  kubeadm upgrade apply v1.12.0
  ```
* Upgrade kubelet on the master node.
  ```
  apt-get upgrade -y kubelet=1.12.0-00
  systemctl restart kubelet
  ```
[Worker]<br />
* Drain the node you are going to upgrade.
  ```
  kubectl drain <node_name>
  ```
* Upgrade the kubeadm and kubelet packages.
  ```
  apt-get upgrade -y kubeadm=1.12.0-00
  apt-get upgrade -y kubelet=1.12.0-00
  ```
* Use kubeadm to upgrade the node configuration.
  ```
  kubeadm upgrade node config --kubelet-version v1.12.0
  ```
* Restart the kubelet
  ```
  systemctl restart kubelet
  ```
* Unmark the node
  ```
  kubectl uncordon <node_name>
  ```

## Backup and Restore Methods
What should you consider backing up in a k8s cluster?<br />
* Resource Configuration
  * Declaritively creating resources is the [preferred] way to save your configurations in a source code repository.
  * Imperatively creating resources can still be backed up and it is advised to perform this type of backup as a fail safe just in case someone didn't create the resources Declaritvely.
* ETCD Cluster
* Persistent Volumes(Optional)

### Backing up Resource Configs
* Backup using the kubectl command to get all resources in all namespaces and output them to a yaml file.
  ```
  kubectl get all --all-namespaces -o yaml > all-deploy-services.yaml
  ```

### Backing up ETCD
* Backup the data directory mounted used by ETCD to store the database.
  * Works best when in a managed environment where you do not have access to the etcd database.
* Take a snapshot of the ETCD database.
  ```
  ETCDCTL_API=3 etcdctl snapshot save snapshot.db \
    --endpoints=https://127.0.0.1:2379 \
    --cacert=/etc/etcd/ca.crt \
    --cert=/etc/etcd/etcd-server.crt \
    --key=/etc/etcd/etc
  ```
  * To view the status of the snapshot use the following command
    ```
    ETCDCTL_API=3 etcdctl snapshot status snapshot.db \
    --endpoints=https://127.0.0.1:2379 \
    --cacert=/etc/etcd/ca.crt \
    --cert=/etc/etcd/etcd-server.crt \
    --key=/etc/etcd/etc
    ```
  * To restore ETCD from a snapshot from etcdctl use the following command.<br />
    1. Stop the kube-apiserver service
       ```
       systemctl stop kube-apiserver
       ```
    2. Run the etcdctl snapshot restore command.
       ```
       ETCDCTL_API=3 etcdctl snapshot restore snapshot.db \
         --endpoints=https://127.0.0.1:2379 \
         --cacert=/etc/etcd/ca.crt \
         --cert=/etc/etcd/etcd-server.crt \
         --key=/etc/etcd/etc \ 
         --data-dir /var/lib/etcd-from-backup
       ```
    3. Configure etcd service file to use the new data directory location.
       ```
       --data-dir=/var/lib/etcd-from-backup
       ```
    4. Reload the ETCD service.
       ```
       systemctl daemon-reload
       ```
    5. Restart the ETCD service.
       ```
       systemctl restart etcd
       ```
    6. Start up the kube-apiserver service.
       ```
       systemctl start kube-apiserver
       ```

### Working ETCDCTL