# Select cluster and authenticate

gcloud container --project PROJECT_NAME clusters get-credentials CLUSTER_NAME --zone=ZONE_NAME

gcloud config set account ACCOUNT_NAME

gcloud auth application-default login

# Load Certificates

kubectl create secret generic tls-cert-&#60;VERSION&#62; --from-file=./cert --namespace=ingress-&#60;ENV_NAME&#62;

# Create Site Configurations

sed -i '' "s/MUI_ENV_NAME/mui-&#60;ENV_NAME&#62;/g" upstreams.conf

sed -i '' "s/MSP_ENV_NAME/msp-&#60;ENV_NAME&#62;/g" upstreams.conf

sed -i '' "s/DOMAIN_NAME/mcom-&#60;DOMAIN_NAME&#62;/g" mcom.conf

sed -i '' "s/DOMAIN_NAME/tcom-&#60;DOMAIN_NAME&#62;/g" tcom.conf

kubectl create configmap sites-config-&#60;VERSION&#62; --from-file=./nginx --namespace=ingress-&#60;ENV_NAME&#62;

    # Note: -i '' for Mac mini and -i option for Linux.

# Deployment

kubectl create -f deployment.yaml

