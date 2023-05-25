# K8s-Manifests

# Создание секрета

```sh
kubectl create secret docker-registry NAME --docker-server=registry.gitlab.com/hse_students/hse_ramus/core --docker-username=USERNAME --docker-password=PASSWORD -n -hse-app
```

Для запуска использовать (создаётся в namespace *hse-app*)
```sh
kubectl apply -f microservice_name-deployment.yaml -n hse-app
kubectl apply -f microservice_name-service.yaml -n hse-app
```

## Запуск POD для PostgreSQL

Нужно смонтировать раздел для Postgres:

```sh
kubectl apply -f postgres-pvc.yaml -n hse-app
```
А затем запустить deployment и service:

```sh
kubectl apply -f postgres-deployment.yaml -n hse-app
kubectl apply -f postgres-service.yaml -n hse-app
```

