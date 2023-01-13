# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_resource', 'helm_resource', 'helm_repo')

compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

#################### ArgoCD ##################

helm_repo('argo', 'https://argoproj.github.io/argo-helm')
helm_resource(
  'argocd',
  'argo/argo-cd',
  namespace='argocd',
  flags=[
    '--create-namespace',
    '--set',
    'configs.secret.argocdServerAdminPassword=$2a$10$txDzpiHZBbiL9c5vYTGD2uutyZkjIFsxHO0/cMf45CbVS.OlHIv8C',
  ],
  deps=['collector/developer.values.yaml'],
  resource_deps=['argo'],
  labels='argocd',
)

local_resource(
  'argocd-port-forwarding',
  '',
  serve_cmd='kubectl port-forward service/argocd-server -n argocd 8080:443',
  resource_deps=['argocd'],
  labels='argocd',
)
