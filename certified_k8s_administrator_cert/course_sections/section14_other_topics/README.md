# SECTION 14: Other Topics

## Advanced Kubectl Commands with JSON Path
* Node names:
  ```
  kubectl get nodes -o=jsonpath='{.items[*].metadata.name}'
  ```
* Node architectures:
  ```
  kubectl get nodes -o=jsonpath='{.items[*].status.nodeInfo.architecture}'
  ```
* Nodes cpu capacity:
  ```
  kubectl get nodes -o=jsonpath='{.items[*].status.capacity.cpu}'
  ```
* Multiple json path queries merged:
  ```
  kubectl get nodes -o=jsonpath='{.items[*].metadata.name}{.items[*].status.capacity.cpu}'
  ```
* Formatting output:
  ```
  kubectl get nodes -o=jsonpath='{.items[*].metadata.name}{"\n"}{.items[*].status.capacity.cpu}'
  ```
* Loops - Range:
  ```
  kubectl get nodes -o=jsonpath='{range.items[*].metadata.name}{"\t"}{.status.capacity.cpu}{"\n}{end}'
  ```
* Json path for custom columns:
  ```
  kubectl get nodes -o=custom-columns=NODE:.metadata.name,CPU:.status.capacity.cpu
  ```
* JSON Path for sorting based on name:
  ```
  kubectl get nodes --sort-by=.metadata.name
  ```
* Searching within jsonpath:
  ```
  kubectl config view --kubeconfig=/root/my-kube-config -o=jsonpath="{.contexts[?(@.context.user=='aws-user')].name}" > /opt/outputs/aws-context-name
  ```