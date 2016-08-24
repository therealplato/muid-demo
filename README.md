muid-demo
=========

This repo contains a sample project that uses github.com/therealplato/muid

Usage
-----

kubectl create namespace muid-demo
docker build -t muid-demo:0.1.0 .
kubectl --namespace muid-demo apply -f k8s/svc-muid-generator.yml
kubectl --namespace muid-demo apply-f k8s/dep-muid-generator.yml
kubectl --namespace muid-demo get pods |grep muid
