# SECTION 8: Storage

## Volumes
* Volumes and mounts
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: random-number-generator
  spec:
    containers:
    - image: alpine
      name: alpine
      commnand: ["/bin/sh", "-c"]
      args: ["shuf -i 0-100 -n 1 >> /opt/numbers.out;"]
      volumeMounts:
      - mountPath: /opt
        name: data-volume
    
    volumes:
    - name: data-volume
      hostPath:
        path: /data
        type: Directory
  ```
* Volume Storage Options
  * hostPath - Not recommended on a multi-node clusters
  * awsElasticBlockStore - AWS block store volume
    ```yaml
    volumes:
    - name: data-volume
      awsElasticBlockStore:
        volumeID: <volume-id>
        fsType: ext4
    ```
  * Any other cloud service or cluster fs
## Persistent Volumes
* Cluster wide pool of storage volumes configured by an administrator to be used by users deploying applications on the cluster.
* Create a PV
  ```yaml
  apiVersion: v1
  kind: PersistentVolume
  metadata:
    name: pv-vol1
  spec:
    accessModes:
      - ReadWriteOnce
    capacity:
      storage: 1Gi
    hostPath:
      path: /tmp/data
  ```
## Persistent Volume Claims
* Every PVC is bound to only ONE PV
* Can use labels and selectores to specify a PV more specifically
* Create Persistent Volume Claim
  ```yaml
  apiVersion: v1
  kind: PersistentVolumeClaim
  metadata:
    name: myclaim
  spec:
    accessModes:
      - ReadWriteOnce
    resources:
      requests:
        storage: 500Mi
  ```
* What happens to a persistent volume when a claim is deleted?
  * Check the `persistentVolumeReclaimPolicy` configuration on the PV. There are three modes PV's can work in.
    * Retain - This retains the PV and the data on it but doesn't allow any other PVC's to bind to it.
    * Delete - Deletes the PV and all its data once the pvc is deleted.
    * Recycle - Scrubs the data from the PV and makes it available to take a new claim.
## Using PVCs in PODS
```yaml
 apiVersion: v1
  kind: Pod
  metadata:
    name: random-number-generator
  spec:
    containers:
    - image: alpine
      name: alpine
      commnand: ["/bin/sh", "-c"]
      args: ["shuf -i 0-100 -n 1 >> /opt/numbers.out;"]
      volumeMounts:
      - mountPath: /opt
        name: data-volume
    
    volumes:
    - name: data-volume
      persistentVolumeClaim:
          claimName: myclaim
```
## Storage Class
* Define a provisioner like google storage that can automatically provision storage on google cloud and attach it to pods as its made.
* Dynamic Provisioning.
* Creates PV automatically
* SC Definiton
  ```yaml
  apiVersion: storage.k8s.io/v1
  king: StorageClass
  metadata:
    name: google-storage
  provisioner: kubernetes.io/gce-pd
  ```
  * Adding storage class to pvc definition to use dynamic provisioning
    ```yaml
    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      name: myclaim
    spec:
      accessModes:
        - ReadWriteOnce
      storageClassName: google-storage
      resources:
        requests:
          storage: 500Mi
    ```
* Provisioners
  * Local
  * AWSElasticBlockStore
  * AzureFile
  * AzureDisk
  * CephFS
  * Cinder
  * FC
  * Glusterfs
  * GCEPersistentDisk
  * etc.