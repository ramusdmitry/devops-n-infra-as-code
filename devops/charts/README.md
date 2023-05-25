# Charts


## Запуск с помощью Helm

```sh
helm install chart_name --generate-name
```

## Генерация чарта и вывод
```sh
helm template chart_name --generate-name
```

## Чарт PostgreSQL

Помимо `deployment.yaml` и `service.yaml`, в папке лежит `pvc.yaml`, отвечающий за создание PersistentVolumeClaims
