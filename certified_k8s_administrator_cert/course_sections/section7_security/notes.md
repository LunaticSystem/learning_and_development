# Section 7: Security
## Kubernetes Security Primitives
* Controlling access to the kubeapi server
  * Authentication Mechanisms
    * Username and password - Files
    * Username and Tokens - Files
    * Certificates
    * External Authentication Providers - LDAP
    * Service Accounts
  * Authorization
    * RBAC Authorization
    * ABAC Authorization
    * Node Authorization
    * Webhook Mode
* All communication in the cluster (ETCD, Kubelet, Kube Proxy, Kube scheduler, and Kube controller manager)
  * Authorize and Authenticate through TLS Certificates.
* Communication between pods withing in the cluster
  * Network Policy
## Authentication
* Users
  * All user access is managed by the api server.
  * kube-apiserver
    * Static Password File
      * CSV files - password, username, userid, groupname(optional)
      * Pass the csv file to the api server service
        ```
        --basic-auth-file=user-details.csv
        ```
      * Restart the kube-apiserver service
      * Or you can pass the csv file to the kube-apiserver.yaml file in a kubeadm env. (/etc/kubernetes/manifests/kube-apiserver.yaml)
        ```
        --basic-auth-file=user-details.csv
        ```
      * In order to authenticate with a user to the kube-apiserver pass the username and password via curl.
        ```
        curl -v -k https://master-node-api:6443/api/v1/pods -u "username:password"
        ```
    * Static Token File
      * CSV file - token, username, userid, group name (optional)
      * pass the csv file to the api server service or api server yaml file.
        ```
        --token-auth-file=user-token-details.csv
        ```
      * To authenticate with the api server you will need to use bearer tokens.
        ```
         curl -v -k https://master-node-api:6443/api/v1/pods --header "Authorization: Bearer <token>"
        ```
  * Notes:
    * Using basic user auth through plaintext username and password is not the recommended approach.
    * Consider volume mount while providing the auth file in a kubeadm setup.
    * Setup Role Based Authorization for the new users.

### Basic Auth Setup Process
1. Create a file with user details locally at /tmp/users/user-details.csv
     ```
     # User File Contents
     password123,user1,u0001
     password123,user2,u0002
     password123,user3,u0003
     password123,user4,u0004
     password123,user5,u0005
     ```

2. Edit the kube-apiserver static pod configured by kubeadm to pass in the user details. The file is located at /etc/kubernetes/manifests/kube-apiserver.yaml
   ```
   apiVersion: v1
   kind: Pod
   metadata:
     name: kube-apiserver
     namespace: kube-system
   spec:
     containers:
     - command:
       - kube-apiserver
         <content-hidden>
       image: k8s.gcr.io/kube-apiserver-amd64:v1.11.3
       name: kube-apiserver
       volumeMounts:
       - mountPath: /tmp/users
         name: usr-details
         readOnly: true
     volumes:
     - hostPath:
         path: /tmp/users
         type: DirectoryOrCreate
       name: usr-details
   ```

3. Modify the kube-apiserver startup options to include the basic-auth file
   ```
   apiVersion: v1
   kind: Pod
   metadata:
     creationTimestamp: null
     name: kube-apiserver
     namespace: kube-system
   spec:
     containers:
     - command:
       - kube-apiserver
       - --authorization-mode=Node,RBAC
         <content-hidden>
       - --basic-auth-file=/tmp/users/user-details.csv
   ```

4. Create the necessary roles and role bindings for these users:
   ```
   ---
   kind: Role
   apiVersion: rbac.authorization.k8s.io/v1
   metadata:
     namespace: default
     name: pod-reader
   rules:
   - apiGroups: [""] # "" indicates the core API group
     resources: ["pods"]
     verbs: ["get", "watch", "list"]
    
   ---
   # This role binding allows "jane" to read pods in the "default" namespace.
   kind: RoleBinding
   apiVersion: rbac.authorization.k8s.io/v1
   metadata:
     name: read-pods
     namespace: default
   subjects:
   - kind: User
     name: user1 # Name is case sensitive
     apiGroup: rbac.authorization.k8s.io
   roleRef:
     kind: Role #this must be Role or ClusterRole
     name: pod-reader # this must match the name of the Role or ClusterRole you wish to bind to
     apiGroup: rbac.authorization.k8s.io
   ```

5. Once created, you may authenticate into the kube-api server using the users credentials
   ```
   curl -v -k https://localhost:6443/api/v1/pods -u "user1:password123"
   ```
## TLS Introduction

## TLS Basics

## TLS in Kubernetes
* kube-apiserver
  * Exposes an HTTPS service
  * apiserver.crt - Certificate
  * apiserver.key - Key
* ETCD
  * ectdserver.crt
  * etcdserver.key
* Kubelet
  * Exposes an HTTPS service
  * kubelet.crt - Certificate
  * kubelet.key - Key
* Clients
  * kubectl
  * restapi
  * admin.crt
  * admin.key
  * Scheduler is treated as a client
  * Kube-controller-manager is treated as a client
  * Kube-proxy is treated as a client
  * kube-apiserver when talking to etcd and kubelet is treated as a client.
## TLS in Kubernetes - Certificate Creation
* To generate a certificate for the CA using openssl tool
  * Create private key.
    ```
    openssl genrsa -out ca.key 2048
    ```
  * Send a certificate signing request.
    ```
    # openssl req -new -key <ca_key_name> -subj "/CN=<name_of_object>" -out ca.csr

    openssl req -new -key ca.key -subj "/CN=KUBERNETES-CA" -out ca.csr
    ```
  * Sign certificate (self signed).
    ```
    openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt
    ```
* Generating the client certificates.
  * Create private key.
    ```
    openssl genrsa -out admin.key 2048
    ```
  * Send a certificate signing request.(A group is required for admin users as they need to kube admin access)
    ```
    # openssl req -new -key <admin_key_name> -subj "/CN=<name_of_object>/O=<group_name>" -out admin.csr

    openssl req -new -key admin.key -subj "/CN=KUBE-ADMIN/O=system:masters" -out admin.csr
    ```
  * Sign certificate (self signed).
    ```
    openssl x509 -req -in admin.csr -CA ca.crt -CAkey ca.key -out admin.crt
    ```
  * For all system components. i.e kube-scheduler, kube-controller-manager, and kube-proxy.
    * Create private key.
        ```
        openssl genrsa -out admin.key 2048
        ```
    * Send a certificate signing request.(A group is required for admin users as they need to kube admin access)
      ```
      # openssl req -new -key <component_key_name> -subj "/CN=SYSTEM:<name_of_object>/O=<group_name>" -out <component>.csr

      openssl req -new -key kube-scheduler.key -subj "/CN=SYSTEM:KUBE-SCHEDULER/O=system:masters" -out kube-scheduler.csr
      ```
    * Sign certificate (self signed).
      ```
      openssl x509 -req -in kube-scheduler.csr -CA ca.crt -CAkey ca.key -out kube-scheduler.crt
      ```
    * To connect with a client certificate.
      * Use curl
        ```
        curl https://kube-apiserver:6443/api/vi/pods \
          --key admin.key \
          --cert admin.crt \
          --cacert ca.crt
        ```
      * Use kube-config.yaml
        ```yaml
        apiVersion: v1
        clusters:
        - cluster:
            certificate-authority: ca.crt
            server: https://kube-apiserver:6443
          name: kubernetes
        kind: Config
        user:
        - name: kubernetes-admin
          user:
            client-certificate: admin.crt
            client-key: admin.key
        ```
* Server Certificates
  * ETCD
  * Kube-apiserver
    * Generate an openssl config file.
      ```
      [req]
      req_extension = v3_req
      distinguished_name = req_distinguished_name
      [ v3_req ]
      basicConstraints = CA:FALSE
      keyUsage = nonRepudiation,
      subjectAltName = @alt_names
      [alt_names]
      DNS.1 = kubernetes
      DNS.2 = kuberenetes.default
      DNS.3 = kubernetes.default.svc
      DNS.4 = kubernetes.default.svc.cluster.local
      IP.1 = <external_ip>
      IP.2 = <internal_ip>
      ```
    * Pass that config to the open ssl key generation request.
      ```
      openssl req -new -key apiserver.key -subj "/CN=kube-apiserver" -out apiserver.csr -config openssl.cnf
      ```
    * Sign the certificate.
      ```
      openssl x509 -req -in apiserver.csr -CA ca.crt -CAkey ca.key -out apiserver.crt
      ```
    * Where to specify keys.
      * keys to specify
        ```
        --etcd-cafile
        --etcd-certfile
        --etcd-keyfile
        --kubelet-certificate-authority
        --kubelet-client-certificate
        --kubelet-client-key
        --client-ca-file
        --tls-cert-file
        --tls-private-key
        ```
      * kube-apiserver service.
      * kube-apiserver manifest file.
  * Kubelet Server
    * Certificate per node in the cluster. Named differently.
      * node01.crt,node02.crt, etc.
    * Specify these certificates in the kubelet-config.yaml per node.
      ```yaml
      kind: KubeletConfiguration
      apiVersion: kubelet.config.k8s.io/v1beta1
      authentication:
        x509:
          clientCAFile: "/var/lib/kubernetes/ca.pem"
      authorization:
        mode: Webhook
      clusterDomain: "cluster.local"
      clusterDNS:
        - "<external_ip_noder>"
      podCIDR: "${POD_CIDR}"
      resolvConf: "/run/systemd/resolve/resolve.conf"
      runtimeRequestTimeout: "15m"
      tlsCertFile: "/var/lib/kubelet/node01.crt"
      tlsPrivateKeyFile: "/var/lib/kubelet/node01.key"
      ```
## View Certificate Details
* The hard way
  * Generate and deploy all certificates manually
* Kubeadm
  * Genrates and deploys certficates automatically
  * How to gather certificate information
    * Gather up all certificate paths by querying the kube-apiserver.yaml
      ```
      kubectl describe pod <kube-apiserver> -n kube-system

      OR

      cat /etc/kubernetes/manifests/kube-apiserver.yaml
      ```
    * Look inside the certifcate file
      ```
      openssl x509 -in /etc/kubernetes/pki/apiserver.crt -text -noout
      ```
      * Subject - Common name used
      * X509 Subject Alternative Names - Alternative names used by this services.
      * Validity(Not After) - The expiration of the certificate
      * Issuer - Name of the ca who issued the certificate
## Certificates API
### Process to create a user cert throught the certificates api.
* User creates a key.
* User creates a CSR with that key
* User sends that key to administrator
* Administrator base64 encodes the csr
  ```
  cat jane.csr |base64 |tr -d "\n"
  ```
* Adminstrator creates a certificate signing request object in k8s. i.e jane-csr.yaml
  ```yaml
  apiVersion: certificates.k8s.io/v1beta1
  kind: CertificateSigningRequest
  metadata:
    name: jane
  spec:
    groups:
    - system:authenticated
    usages:
    - digital signature
    - key encipherment
    - server auth
    request:
         LXXXXXXXXXXXXXXXXXXXXXXXXX
         XXXXXXXXXXXXXXXXXXXXXXXXXX
         XXXXXXXXXXXXXXXXXXXXXXXXXX
         XXXXXXXXXXXXXXXXXXXXXXXXXX
         XXXXXXXXXXXXXX
  ```
* Submit the request to k8s.
  ```
  kubectl apply -f jane-csr.yaml
  ```
### Get and Approve CSR requests in K8s.
* Gather up the available csr requests in k8s.
  ```
  kubectl get csr
  ```
* Approve CSR request.
  ```
  kubectl certificate approve jane
  ```
### View Certificate generated after approval.
* Gather the certificate information from the csr.
  ```
  kubectl get csr jane -o yaml |grep certificate
  ```
* Decode the base64 to get the actually certificate.
  ```
  echo "XXXXXXXXXXXXXXXXXXXXX=" |base64 -d
  ```

### Information about the Certificate API
* The kube-controller-manager is tasked with managing the certificate api.
* Controller manager has two options that specify the ca signing cert and key.
  ```
  --cluster-signing-cert-file
  --cluster-signing-key-file
  ```
## KubeConfig
* Kubectl defaultly looks for kubeconfig files in the `$HOME/.kube/config` location. 
* Format:
  * Clusters - Clusters you have access too.
  * Contexts - Define which user account will be used to acces which cluster
  * Users - User accounts that have access to these clusters.
* Kubeconfig Example:
  ```yaml
  apiVersion: v1
  kind: Config

  clusters:
  - name: my-kube-playground
    cluster:
      certificate-authority: ca.crt
      server: https://my-kube-playgound:6443

  contexts:
  - name: my-kube-admin@my-kube-playground
    context:
      cluster: my-kube-playground
      user: my-kube-admin

  users:
  - name: my-kube-admin
    user:
      client-certificate: admin.crt
      client-key: admin.key
  ```
* To set a default context in the kube config add the following entry between the kind and clusters section in the KubeConfig.
  ```yaml
  kind: Config
  current-context: my-kube-admin@my-kube-playground

  clusters:
  ```
* To get the KubeConfig with kubectl.
  ```
  kubectl config view
  ```
* To get informaation about a non default kubeconfig.
  ```
  kubectl config view --kubeconfig=my-custom-config
  ```
* Supplying a namespace in a context in the kubeconfig.
  ```yaml
  contexts:
  - name: my-kube-admin@my-kube-playground
    context:
      cluster: my-kube-playground
      user: my-kube-admin
      namespace: finance
  ```
* To add a certificate not as a file do the following.
  * Encode the certificate.
    ```
    cat ca.crt |base64
    ```
  * Supply the base64 encoded value to the `certificate-authority-data` field in cluster instead of the `certificate-authority` field.
    ```yaml
    clusters:
    - name: my-kube-playground
      cluster:
        certificate-authority-data: L0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX==
        server: https://my-kube-playgound:6443
    ```
## API Groups
* Check api version.
  ```
  curl https://kube-master:6443/version
  ```
### Groups
* /version - Used to check the version of the api on the cluster.
* /api - Core group is where the all core functionality exists
  * namespaces
  * pods
  * rc
  * events
  * endpoints
  * nodes
  * bindings
  * PV
  * PVC
  * configmaps
  * secrets
  * services
* /apis - Named group is where all new funtionality will be stored currently and in the future.
  * /apps
    * /v1
      * /deployments (resources)
        * list (actions/verbs)
        * get
        * create
        * delete
        * update
        * watch
      * /replicasets
      * /statefulsets
  * /extensions
  * /networking.k8s.io
    * /v1
      * /networkpolicies
  * /storage.k8s.io
  * /authentication.k8s.io
  * /certificates.k8s.io
    * /v1
      * /certificatesigningrequests
* /logs
* /healthz - Monitor health of cluster
* /metrics - Monitor health of cluster

### Listing out all groups and resources within those groups.
```
curl http://localhost:6443 -k \
  --key admin.key
  --cert admin.crt
  --cacert ca.crt

curl http://localhost:6443/apis -k --key admin.key --cert admin.crt --cacert ca.crt |grep "name"
```

You can also use kubectl proxy to proxy requests using the certs provided in your kubeconfig like so..
```
kubectl proxy
starting to serve on 127.0.0.1:8001

curl http://localhost:8001 -k
```
## Authorization

## Role Based Access Controls

## Cluster Roles and Bindings

## Service Accounts (Not require for CKA)

## Service Account Updates

## Image Security

## Pre-requisite - Security in Docker

## Security Contexts

## Network Policy

## Kubectx and Kubens - Command line Utilities
