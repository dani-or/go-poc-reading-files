# go-poc-reading-files
Este proyecto lee el archivo `resource/FINACLE_NEQUICARTERA_20200508_VENCIDOS.csv` en el bucket `nequi-s3-select-tmp-2` 

### Diagrama de Clases 
![Diagram de clases ](/go-poc-reading-files/resources/images/clases.png?raw=true "Diagrama de clases")

# Install local
Para instalar localmente el proyecto 
## Pre-requisites
##- Docker

## Construir la imagen 
```bash
docker image build -t go-poc-reading-files .
```
## Correr la imagen
```bash
docker run --network host -d go-poc-reading-files
```
## Probar la funcionalidad
Para probar esto solo es hacer el despliegue y mirar los logs del docker , para ver los logs del docker se ejecuta este comando:
```bash
docker logs #iddeldocker
```

# Install on EKS
Esta guia de instalación contiene:
- Crear artefactos de red
- Crear clúster
- Crear proveedor de identididad
- Crear de política
- Crear service account
- Crear del deployment

 Nota: La ejecución debe ser en orden

## Pre-requisites
##- Kubectl —  https://kubernetes.io/docs/tasks/tools/install-kubectl/
##- AWS CLI -  https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
##- Aws iam authenticator — https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html
##- eksctl — https://github.com/weaveworks/eksctl

## Crear artefactos de red -- AWS CLI
Se crean artefactos de red: vpc, subnets públicas y privadas, grupos de seguridad entre otros
```bash
aws cloudformation deploy --template-file red.yaml --stack-name my-new-stack
```

## Crear clúster
Para la creación del clúster se debe cambiar en el archivo cluster.yaml el id de las vpc y las subnets
```bash
eksctl create cluster -f cluster.yaml
```

## Crear proveedor de identididad
El proveedor de identidad se crea para que los pods puedan asumir un role de IAM y así acceder a los diferentes recursos de AWS, este paso es opcional para mejor documentación [aquí](https://dzone.com/articles/how-to-use-aws-iam-role-on-aws-eks-pods).
```bash
eksctl utils associate-iam-oidc-provider --cluster <CLUSTER_NAME> --approve
```

## Crear el policy
Este política se la vamos asignar al role que vamos a crear en el siguiente paso
```bash
aws iam create-policy --policy-name go-poc-reading-files-role-pilicy --policy-document file://policy.json
```

## Crear service account
Acá creamos el role que le vamos asigna a nuestro pod y le asociamos la política que creamos en el paso anterior

```bash
eksctl create iamserviceaccount \
  --cluster=EKS-Demo-Cluster1 \
  --role-name=go-poc-reading-files-role-1 \
  --namespace=default \
  --name=go-poc-reading-files-service-account-name-1  \
  --attach-policy-arn=arn:aws:iam::851560454673:policy/go-poc-reading-files-role-pilicy \
  --approve
```

## Crear del deployment
El despliegue incluye la arn del repositorio ecr donde tenemos la imagen que queremos desplegar, el puerto que queremos exponer, los labels con los que queremos clasificar nuestra app y el service account que creamos en el paso anterior que es el que nos permite que el o los pods asuman un rol de IAM
```bash
kubectl apply -f deployment.yaml
```

# Pruebas en EKS
Para probar esto solo es hacer el despliegue y mirar los logs del pod , para ver los logs del pod se ejecuta este comando:
```bash
kubectl logs id_del_pod_que_se_desplegó_en_el_deployment
```