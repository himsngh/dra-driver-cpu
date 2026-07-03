# Sysfs overlay example

This example supplies a synthetic four-CPU topology. The ConfigMap owns the overlay data, while the values file uses the chart's generic extra arguments, volume mounts, and volumes.

One environment where this is useful is a Kind cluster running in Docker Desktop for macOS, where the container-visible sysfs may not expose complete NUMA topology. The CPU and NUMA portions of this example work on Apple Silicon; PCIe-root discovery is currently limited to AMD64.

Create the overlay ConfigMap and deploy the chart:

```bash
kubectl apply -n kube-system -f hack/examples/sysfs-overlay/configmap.yaml

helm upgrade --install dra-driver-cpu deployment/helm/dra-driver-cpu \
  --namespace kube-system \
  --set fullnameOverride=dracpu \
  -f hack/examples/sysfs-overlay/values.yaml
```

The example exposes two NUMA-grouped devices with two CPUs each. It also defines one PCIe root local to each NUMA node: `pci0000:00` for CPUs 0-1 and`pci0001:80` for CPUs 2-3.

When changing the ConfigMap after deployment, restart the DaemonSet to reload the overlay:

```bash
kubectl rollout restart daemonset/dracpu -n kube-system
```
