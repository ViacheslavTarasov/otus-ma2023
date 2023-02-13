# Домашние задания по курсу "Microservice Architecture"

## Домашнее задание 1

Исходные файлы в ./hw1. Для проверки можно скачать docker-образ:
```bash
docker pull viacheslavtarasov/otus-ma2023-hw1:latest
docker run -p 8080:8000 viacheslavtarasov/otus-ma2023-hw1:latest
curl localhost:8080/health
```

## Домашнее задание 2

Добавлены файлы конфигурации deployment, service, ingress.

Исходные файлы в ./hw2. Для проверки:
1. Клонируем репозиторий и применяем конфиг
```bash
kubectl apply -f ./hw2/config/k8s/
```
2. Выполняем запросы
```bash
curl http://arch.homework/health
curl http://arch.homework/otusapp/health
```