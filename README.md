# Infra-Operator

## Project to demo kubernetes operator

- Infra-Operator is kubernetes operator developed using Operator SDK
- It will demonstrate working of reconciler, watches and webhooks.
- It has few custom resources defined and have controller recocilers implemeted for those CRD's 

## Prerequisites
You need to have [Operator SDK](https://sdk.operatorframework.io/docs/installation/install-operator-sdk/) and Go installed to be able to build the Infra-Operator. 

## Deployment
Use [Make file](Makefile) to build and deploy project.

Follwing are the commands to build and deploy Operator 
    
 - To run the operator locally use
   
    make install && make run ENABLE_WEBHOOKS=false WATCH_NAMESPACE=default 

 - To create and push docker image use
    
    make docker-push  IMG=<Your Repository>/<Image Name>:<Tag>


