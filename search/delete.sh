kubectl delete deploy nginx-dev --namespace=ingress-dev
kubectl delete deploy nginx-qa --namespace=ingress-qa
kubectl delete deploy nginx-reldev --namespace=ingress-reldev


kubectl delete configmap sites-config-dev --namespace=ingress-dev
kubectl delete configmap sites-config-qa --namespace=ingress-qa
kubectl delete configmap sites-config-reldev --namespace=ingress-reldev

kubectl delete secret tls-cert-dev --namespace=ingress-dev
kubectl delete secret tls-cert-qa --namespace=ingress-qa
kubectl delete secret tls-cert-reldev --namespace=ingress-reldev

kubectl delete namespace ingress-dev
kubectl delete namespace ingress-qa
kubectl delete namespace ingress-reldev