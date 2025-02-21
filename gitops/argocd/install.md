Install Argo CD using manifests

kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

to access argocd , we need to expose service 
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'

hit the alb dns on browser , username admin , password run below command
kubectl get secrets -n argocd 

kubectl edit secret argocd-initial-admin-secret -n argocd

get password
echo NnI0NGJEelZYUDd5dTdJcQ== | base64 --decode

6r44bDzVXP7yu7Iq

