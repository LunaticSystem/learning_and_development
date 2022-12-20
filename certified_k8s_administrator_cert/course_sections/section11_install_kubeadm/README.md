# SECTION 11: Install K8s with Kubeadm

## Steps to setup cluster with kubeadm
* Provision vm's for master and worker nodes
* Install a container runtime on all nodes
* Install kubeadm tool on all nodes.
* Initialize the master server
* Pod network
* Join the worker nodes to the master node

[Vagrant File to deploy three node cluster to test kubeadm](https://github.com/kodekloudhub/certified-kubernetes-administrator-course.git)

## Deploying a kubernetes cluster using kubeadm

* Create k8s.conf file with the following command:
  ```
  cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
  net.bridge.bridge-nf-call-ip6tables = 1
  net.bridge.bridge-nf-call-iptables = 1
  EOF
  ```
* Run sysctl to reload the configuration with the above properties. 
  ```
  sudo sysctl --system
  ```
* Update your apt repos.
  ```
  sudo apt-get update
  ```
* Install apt-transport-https ca-certificates and curl.
  ```
  sudo apt-get install -y apt-transport-https ca-certificates curl
  ```
* Pull down kubernetes gpg keys locally.
  ```
  sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
  ```
* Echo the below contents into the kubernetes.list file to configure the apt repo.
  ```
  echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
  ```
* Update your apt repos again.
  ```
  sudo apt-get update
  ```
* Install kubelet kubeadm and kubectl
  ```
  sudo apt-get install -y kubelet=1.24.0-00 kubeadm=1.24.0-00 kubectl=1.24.0-00
  ```
* Freeze all packages above at the their current version.
  ```
  sudo apt-mark hold kubelet kubeadm kubectl
  ```


## Installing cluster via kubeadm on controlplane
* kubeadm init --apiserver-advertise-address 10.44.59.9 --pod-netwrok-cidr 10.244.0.0/16 --apiserver-cert-extra-sans controlplane

[Section 12: End to End Testing](https://github.com/LunaticSystem/learning_and_development/tree/main/certified_k8s_administrator_cert/course_sections/section12_e2e_tests)