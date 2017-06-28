gcloud config set compute/zone us-central-1a
gcloud container clusters get-credentials mean-cluster
kubectl create namespace ingress-dev
sed -i 's/ENV_NAME/dev/g' ./kubernetes/deployment.yaml
sed -i 's/ENV_NAME/dev/g' ./kubernetes/secret.json
kubectl create configmap sites-config-dev --from-file=./kubernetes/configmap
kubectl create -f ./kubernetes/secret.json
kubectl create -f ./kubernetes/deployment.yaml