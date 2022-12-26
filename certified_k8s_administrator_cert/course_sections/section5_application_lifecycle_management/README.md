# SECTION 5: APPLICATION LIFECYCLE MANAGEMENT

## Rolling Updates and Rollbacks
* Rollout
  * Creates a new deployment revision any time a change has been made to the deployment.
  * Enables us to rollback to an older version
  * You can get the rollout status for a deployment by running the following command.
    ```
    kubectl rollout status deployment/myapp-deployment
    ```
  * You can get the rollout reversions by running the following command.
    ```
    kubectl rollout history deployment/myapp-deployment
    ```
* Deployment Strategies
  * Recreate Strategy - Not default deployment strategry as it destroys pods before other pods are created.
  * Rolling Update - Default deployment strategy as it creates new application pods before the old pods are destroyed. i.e zero downtime.
* Rollback
  * Command to rollback a deployment update
    ```
    kubectl rollout undo deployment/myapp-deployment
    ```
* Versioning
## Configure Applications

### Commands
* CMD - Command line arguments that get overwritten by arguments you pass at the command line.
* ENTRYPOINT - Sets a command to be run with specific flags but cannot 
### Commands And Arguments

### Configure Environment Variables in Applications
* Setting an ENV var.
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: simple-webapp-color
  spec:
    containers:
    - name: simple-webapp-color
      image: simple-webapp-color
      ports:
        - containerPort: 8080
      env:                     <----------- Env variable block
        - name: APP_COLOR      <----------- Name of env var
          value: pink          <----------- Value of env var
  ```
* Setting an ENV var from a configMap.
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: simple-webapp-color
  spec:
    containers:
    - name: simple-webapp-color
      image: simple-webapp-color
      ports:
        - containerPort: 8080
      env:                     <----------- Env variable block
        - name: APP_COLOR      <----------- Name of env var
          valueFrom:
              configMapKeyRef: <----------- Value from configMap
  ```
* Setting an ENV var from a Secret.
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: simple-webapp-color
  spec:
    containers:
    - name: simple-webapp-color
      image: simple-webapp-color
      ports:
        - containerPort: 8080
      env:                     <----------- Env variable block
        - name: APP_COLOR      <----------- Name of env var
          valueFrom:
              secretKeyRef: <----------- Value from secret
  ```
### Configuring ConfigMaps in Applications
* ConfigMaps are used to pass configuration data in the form of Key: Value pairs in kubernetes.
* Imperative creation of configMaps.
  ```
  kubectl create configmap \
     <config-name> --from-literal=<key>=<value>

  kubectl create configmap <config-name> --from-file=<path-to-file>
  ```
* Declaarative creation of configMaps.
  * Config Map Yaml File
    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: app-config
    data:
      APP_COLOR: blue
      APP_MODE: prod
    ```
  * Add config map to pod definition.
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: simple-webapp-color
      labels:
        name: simple-webapp-color
    spec:
      containers:
      - name: simple-webapp-color
        image: simple-webapp-color
        ports:
          - containerPort: 8080
        envFrom:
          - configMapRef:
                name: app-config
    ```
### Configured Secrets in Applications
* Secrets are used to store sensitive information. Stored in an encoded format.
* Process to create and use secrets
  * Create Secret
    * Imperative
      ```
      kubectl create secret generic <secret-name> --from-literal=<key>=<value> --from-literal=<key2>=<value2>

      kubectl create secret generic <secret-name> --from-file=<path-to-file>
      ```
    * Declarative
      * Encode all of your values for the secret in base64
        ```
        echo -n "mysql" |base64
        bXlzcWw=
        echo -n "root" |base64
        cm9vdA==
        echo -n "paswrd" |base64
        cGFzd3Jk
        ```
      * Create Definition File
        ```yaml
        apiVersion: v1
        kind: Secret
        metadata:
          name: app-secret
        data:
          DB_HOST: bXlzcWw=
          DB_USER: cm9vdA==
          DB_Password: cGFzd3Jk
        ```
     * Get Full Secret Information
       ```
       kubectl get secret app-secret -o yaml
       ```
  * Configure Pod to use Secret for Single ENV Vars
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: simple-webapp-color
      labels:
        name: simple-webapp-color
    spec:
      containers:
      - name: simple-webapp-color
        image: simple-webapp-color
        ports:
          - containerPort: 8080
        envFrom:
          - secretRef:
                name: app-secret
                key: DB_Password
    ```
  * Configure Pod to mount secret as a volume
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: simple-webapp-color
      labels:
        name: simple-webapp-color
    spec:
      containers:
      - name: simple-webapp-color
        image: simple-webapp-color
        ports:
          - containerPort: 8080
        envFrom:
    volumes:
    - name: app-secret-volume
      secret:
        secretName: app-secret
    ```
* Notes on secrets
  * Secrets are not Encrypted. Only encoded
  * Do not check-in secret objects to SCM along with code.
  * Secrets are not encrypted in ETCD - Consider enabling encrypting secrets at rest.
    * https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/
  * Anyone able to create pods/deployments in the same namespace can access the secrets
    * Configure least-privilege access to secrets - RBAC
  * Consider third-party secrets store providers AWS Provider, Azure Provider, GCP Provider, Vault Provider
### Scale Applications
* Was already explained in the deployments section
## Multi Container PODS
* Creating a multi container pod definition
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: simple-webapp-color
    labels:
      name: simple-webapp-color
  spec:
    containers:
    - name: simple-webapp-color
      image: simple-webapp-color
      ports:
        - containerPort: 8080
    - name: log-agent
      image: log-agent
  ```
## Multi-container PODS Design Patterns
* Side Car Pattern
  ```yaml
  metadata:
  name: simple-webapp
  labels:
    app: webapp
  spec:
    containers:
      - name: main-application
        image: nginx
        volumeMounts:
          - name: shared-logs
            mountPath: /var/log/nginx
      - name: sidecar-container
        image: busybox
        command: ["sh","-c","while true; do cat /var/log/nginx/access.log; sleep 30; done"]
        volumeMounts:
          - name: shared-logs
            mountPath: /var/log/nginx
    volumes:
      - name: shared-logs
        emptyDir: {}
  ```
* Adapter Pattern
* Ambassador Pattern
## InitContainers
* Is configured in a pod like all other containers, except that it is specified inside a `initContainers` section, like below
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: my-app
    labels:
      name: myapp
  spec:
    containers:
    - name: myapp-container
      image: busybox:1.28
      command: ['sh', '-c', 'echo The app is running! && sleep 3600']
    initContainers:
    - name: init-myservice
      image: busybox
      command: ['sh', '-c', 'git clone <some-repo-that-will-be-used-by-application> ; done;']
  ```
* Init containers must run to completion before the real container hosting the application starts.
* If any init containers fail to complete, kubernetes restarts the pod repeatedly until the init container succeeds.
* Read more about initContainers in https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
## Self Healing Applications
* Kubernetes supports self-healing applications through ReplicaSets and Replication Controllers.
  * Replication controller helps in ensuring that a POD is re-created automatically when the application within the POD crashes.
  * It also helps in ensuring enough replicas of the application are running at all times.
* Kubernetes provides additional support to check the health of applications running within PODs and take necessary actions through Liveness and Readiness Probes.  <-----------Not Required for CKA but is require for CKAD

[Section 6: Cluster Maintenence](https://github.com/LunaticSystem/learning_and_development/tree/main/certified_k8s_administrator_cert/course_sections/section6_cluster_maintenance)