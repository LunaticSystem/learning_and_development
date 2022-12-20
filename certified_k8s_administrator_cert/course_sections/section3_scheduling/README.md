# SECTION 3: SCHEDULING

## Manual Scheduling

Because you are unable to schedule a pod by changing the "nodeName" option in the pod definition you will need to create a binding pod definition and make a curl request to the a curl request to the kubeapi in order to manually schedule the pod on the node you want. See the process below:

1. Create a pod definition and apply it:
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: nginx
    labels:
      name: nginx
  spec:
    containers:
    - name: nginx
      image: nginx
      resources:
        limits:
          memory: "128Mi"
          cpu: "500m"
      ports:
        - containerPort: 12345
  ```
2. Create a pod binding definition yaml file:
   ```yaml
   apiVersion: app/v1
   kind: Binding
   metadata:
     name: nginx
   target:
     apiVersion: app/v1
     kind: Node
     name: <name_of_node>
   ```
3. Send a curl request to the kubeapi binding endpoint with the yaml converted to json as the data.
   ```shell
   curl --header "Content-Type:applicatoin/json" --request POST --data '{"apiVersion": "app/v1", "kind": "Binding", "metadata": {"name":"nginx"},"target": {"apiVersion": "app/v1","kind": "Node","name": "<name_of_node>"}}' http://$SERVER/api/v1/namespaces/default/pods/$PODNAME/binding/
   ```


  curl --header "Content-Type:applicatoin/json" --request POST --data '{"apiVersion": "v1", "kind": "Binding", "metadata": {"name":"nginx"},"target": {"apiVersion": "v1","kind": "Node","name": "node1"}}' http://10.4.250.9/api/v1/namespaces/default/pods/nginx/binding/


  ## Labels & Selectors

  Labels - Are properties that are attached to each object for their class, kind, color, etc.

  Selectors - Used to filter objects based on their labels

  Label example:
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: nginx
    labels:                      <----
      name: nginx                <----
      env: prod                  <----
      featureflag: somethingcool <----
  spec:
    containers:
    - name: nginx
      image: nginx
      resources:
        limits:
          memory: "128Mi"
          cpu: "500m"
      ports:
        - containerPort: 12345
  ```

Selector Exmaple:
```
kubectl get pods --selector env=prod
```

Annotations - Record other details for information purposes.

## Taints & Tolerations

Taints - Set on nodes

Tolerations - Set on pods.

Adding taints to nodes via command line:
```
kubectl taint nodes node-name key=value:taint-effect
```

Gather taint definitions from nodes:
```
k describe node controlplane |grep -i taint
```

Taint Effects:
* NoSchedule - Pods will not be scheduled on the node
* PreferNoSchedule - System will try to avoid placing a pod on a node **not guaranteed**
* NoExecute - New pods will not be scheduled on the node and existing pods on the node will be evicted unless they tolerate the taint.

Adding a toleration to a pod:
```yaml
apiVersion:
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: nginx-container
    image: nginx
  tolerations: 
  - key: "app"
    operator: "Equal"
    value: "blue"
    effect: "NoSchedule"
```

## Node Selectors
Requires that the node is labelled and that the pod definition uses nodeSelector to select by label name and value.

Limitations:
* 
```yaml
apiVersion:
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: nginx-container
    image: nginx
  nodeSelector:
    size: Large
```

## Node Affinity & AntiAffinity
Primary purpose of the Node affinity feature is to ensure pods are scheduled on the appropriate nodes.

```yaml
apiVersion:
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: nginx-container
    image: nginx
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: size
            operator: In/NotIn/Exists
            values:
            - Large
```

## Mutliple Schedulers

* Example Scheduler Yaml Config:<br />
  <br />my-scheduler.yaml
  ```yaml
  apiVersion: kubescheduler.config.k8s.io/v1
  kind: KubeSchedulerConfiguration
  profiles:
   - schedulerName: my-scheduler
  ```

* To set up a new scheduler copy the original kube-scheduler.service file into a new file and make sure the config is pointed at the new config and the exec start binary location is correct.
  ```bash
  ExecStart=/usr/local/bin/kube-scheduler --config=/etc/kubernetes/config/my-scheduler.yaml
  ```

* Deploying scheduler as a pod. I.e for Kubeadm clusters.
  <br /><br /> my-custom-scheduler.yaml
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: my-custom-scheduler
    namespace: kube-system
  spec:
    containers:
    - command:
      - kube-scheduler
      - --address=127.0.0.1
      - --kubeconfig=/etc/kubernetes/scheduler.conf
      - --config=/etc/kubernetes/my-scheduler-config.yaml

      image: k8s.gcr.io/kube-scheduler-amd64:v1.11.3
      name: kube-scheduler
  ```

* When using multiple masters it is advised to also add th leaderElection block to your scheduler config like the below:
  ```yaml
  apiVersion: kubescheduler.config.k8s.io/v1
  kind: KubeSchedulerConfiguration
  profiles:
   - schedulerName: my-scheduler
  leaderElection:
    leaderElect: true
    resourceNamespace: kube-system
    resourceName: lock-object-my-scheduler
  ```

* How to tell a pod to use the new scheduler.
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: nginx
  spec:
    containers:
    - image: nginx
      name: nginx
    schedulerName: my-customer-scheduler
  ```

* Get events to check which scheduler picked up your pod.
  ```bash
  $ kubectl get events -o wide -n kube-system |grep Scheduled
  ```

* View logs of scheduler when running into issues.
  ```
  $ kubectl logs my-custom-scheduler -n kube-system
  ```

* Sheduler Process:
  * Scheduling Queue (Plugin: PrioritySort) - This is where the scheduler checks the priorityClass of the pod to prioritize pods with higher priority classes then others.
    * Priority Class Manifest:
      ```yaml
      apiVersion: scheduling.k8s.io/v1
      kind: PriorityClass
      metadata:
        name: high-priority
      value: 1000000
      globalDefault: false
      description: "This priority calss should be used for XYZ service pods only"
      ```
    * Adding priority class to pod manifests:
      ```yaml
      apiVersion: v1
      kind: Pod
      metadata:
        name: simple-webapp-color
      spec:
        priorityClassName: high-priority   <-------
        containers:
        - name: simple-webapp-color
          image: simple-webapp-color
            resources:
              requests:
                memory: "1Gi"
                cpu: 10
      ```
  * Filtering (Plugins: NodeResourcesFit, NodeName, NodeUnschedulable) - This is where the scheduler filters out the nodes that are unable to run the pod based on resources, taints & tolerations, and Node-Affinity.
  * Scoring (Plugin: NodeResourcesFit, ImageLocality) - This is the phase where it scores different weights
  * Binding (Plugin: DefaultBinder) - The phase where a pod is bound to the node with the highest score.

* Scheduler Profiles:
  * Configuring different scheduler profiles:
    ```yaml
      apiVersion: kubescheduler.config.k8s.io/v1
      kind: KubeSchedulerConfiguration
      profiles:
       - schedulerName: my-scheduler-1
         plugins:
           score:
             disabled:
              - name: TaintToleration
             enabled:
              - name: MyCustomPluginA
              - name: MyCustomPluginB
       - schedulerName: my-scheduler-2
         plugins:
           preScore:
             disable:
              - name: '*'
           score:
             disable:
              - name: '*'
    ```

[Section 4: Logging & Monitoring](https://github.com/LunaticSystem/learning_and_development/tree/main/certified_k8s_administrator_cert/course_sections/section4_logging_monitoring)

  
