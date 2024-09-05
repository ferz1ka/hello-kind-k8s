# Hello kind k8s!!

Note: For WSL2 users, we need to disable cgroupsv1 (and therefore utilize only cgroupsv2) to use kind clusters! Specifically, you need to add the following lines to the `.wslconfig` file (for my user of Ferz1ka, this would be `C:\Users\Ferz1ka\.wslconfig`) and restart WSL2 by executing `wsl --shutdown`: 

```
[wsl2]
kernelCommandLine = cgroup_no_v1=all systemd.unified_cgroup_hierarchy=1
```

Start creating kind cluster with `kind create cluster --config=.k8s/kind.yaml --name my-cluster`;

Make sure the cluster is up and running with `kubectl get nodes`;

[optional] Build the server.go image with `docker build -t server-go src/.`;

### Creating pods

Apply the k8s pod spec with `kubectl apply -f .k8s/pod.yaml` using the server-go image as the container image in `.k8s/pod.yaml`;

Check if pod is running with `kubectl get pods`;

### Creating replicasets

Apply the k8s replicaset spec with `kubectl apply -f .k8s/replicaset.yaml` using the server-go image as the container image in `.k8s/replicaset.yaml`;

Note: If you apply a new image version to a replicaset spec, old pods will NOT be deleted and will continue to run with the old image. You will need to delete the old pods manually to get the new image running.

### Creating deployments

Apply the k8s deployment spec with `kubectl apply -f .k8s/deployment.yaml` using the server-go image as the container image in `.k8s/deployment.yaml`;

Note: If you apply a new image version to a deployment spec (chaning kind from `replicaset` to `deployment`), a new replicaset will be created and old pods will automatically be terminated and recreated with the new image version.

#### NOTE! 

Kubernetes objects are created in the following order: Deployments > Replicasets > Pods

### Creating services

Apply the k8s service spec with `kubectl apply -f .k8s/service.yaml`. To access the service, use `kubectl port-forward service/go-server-service 9000:8000` and access it at `http://localhost:8080`.

Note: Inside the `.k8s/service.yaml` spec file, the port is the service port and targetPort is the container port. So in this example, 9000 is the host port, 8000 is the service port and 80 is the container port. We also have more types of services, such as `LoadBalancer` (applies a external port to the service, used with cloud providers) and `NodePort` (applies same port to all nodes redirecting to the service).

### Creating ConfigMaps

Apply the k8s configmap spec with `kubectl apply -f .k8s/configmap-env.yaml`. This configmap will be used to set the name and version of the Hello World application.

Note: Any changes applied to the configmap spec will NOT recreate deployments!

### Creating Secrets

Apply the k8s secret spec with `kubectl apply -f .k8s/secret.yaml`. This secret will be used to set the username and password of the Hello World application.

Note: Secrets env keys are base64 encoded!