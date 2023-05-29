# DevOps and Infrastructure as Code for Classic Three-tier Architecture 

Данный репозиторий представляет реализацию курсового проекта веб-приложения с использованием микросервисной архитектуры на языке Go, развёрнутого в кластере Kubernetes и имеющего скрипты для CI/CD в GitLab. 

<img alt="image" src="https://github.com/ramusdmitry/devops-n-infra-as-code/assets/65753926/9b2512ac-affd-4ddf-bf28-0bb9aa9dc3da">


## Архитектура

Веб-приложение представлено в классической трёхуровневой архитектуре "Клиент - Сервер - Данные". Архитектура проекта и ПО описана через контекстную и контейнерную диаграмму по нотации C4.

#### Level 1: Контекстная C4 диаграмма
![image](https://github.com/ramusdmitry/devops-n-infra-as-code/assets/65753926/92722e57-9ba5-446e-b4d1-8c3cd8cd2bd0)

#### Level 2: Контейнерная C4 диаграмма
![image](https://github.com/ramusdmitry/devops-n-infra-as-code/assets/65753926/d617c9b4-ee51-4f40-b9b1-35780f31a03d)

## Структура репозитория
Проект разбит на группу ``core`` и ``devops``. Группа ``core`` включает в себя исходный код следующих микросервисов:

- ``auth:`` сервис регистрации и авторизации
- ``chat:`` real-time чат
- ``posts:`` сервис публикации постов
- ``comms:`` сервис публикации комментариев к постам, интегрирован с posts
- ``group:`` cервис управления ролями пользователей
- ``frontend:`` сервис фронтенда

В группе ``devops`` собраны инструменты для развёртывания приложения в **Kubernetes**: манифесты для k8s-кластера, шаблоны helm-charts для их автоматизации и диаграммы по нотации **C4** с общим представлением архитектуры веб-приложения. 

## Реализация

**Backend**

Общение между микросервисами, кроме chat-service, происходит с помощью **REST API**. Сервис real-time чата использует **websockets**. Для реализации сервера в ``gо`` проектах был применён фреймворк [gin](https://github.com/gin-gonic/gin), который до ``40х`` быстрее стандартного пакета ``http``. В качестве базы данных была выбрана СУБД PostgreSQL. 

**Frontend**

Использован [Bootstrap](https://getbootstrap.com/) для адаптивной разметки и минималистичного интерфейса, ``js``-пакет [axios](https://axios-http.com/docs/intro) для асинхронных запросов к микросервисам и пакет [react-router-dom](https://reactrouter.com/en/main/start/tutorial) для роутинга между страницами. 

<img width="557" alt="image" src="https://github.com/ramusdmitry/devops-n-infra-as-code/assets/65753926/dcd57102-7b4c-43c7-a21f-2f0d8a440536">

* [Docker](https://www.docker.com/) Инструмент для развёртывания и управления контейнерами
* [Prometheus](https://prometheus.io/) Система мониторинга и сбора метрик
* [Grafana](https://grafana.com/) Операционный дашборд
* [Kubernetes](https://kubernetes.io/) Production-Grade Container Orchestration

## Список используемого и требуемого ПО

- docker-desktop >= 4.13.1
- kubernetes >= 1.25.2
- go >= 1.19
- reactjs >= 18.0.0
- postgresql >= 1.15


