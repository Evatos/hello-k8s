# Hello Kubernetes

Мой первый проект на Kubernetes - простой Go веб-сервер с двумя репликами.

## Что это

Веб-сервер на Go который:
- Отвечает "Hello from Kubernetes!"
- Показывает имя пода
- Имеет health check endpoint
- Работает в 2 репликах с load balancing

## Технологии

- Go 1.26
- Docker
- Kubernetes (kind)
- kubectl


## 📦 Установка и запуск

### Предварительные требования
# Проверь что установлено:
```bash
go version
docker --version
kubectl version
kind version
```

### Шаг 1: Создать kind кластер (если ещё нет)
```bash
kind create cluster --name learning
```

### Шаг 2: Собрать Docker образ
```bash
docker build -t hello-k8s:v1 .
```

### Шаг 3: Загрузить образ в kind
```bash
kind load docker-image hello-k8s:v1 --name learning
```

### Шаг 4: Задеплоить в Kubernetes
```bash
kubectl apply -f k8s/
```

### Шаг 5: Проверить что работает

# Смотрим поды
```bash
kubectl get pods
```

### Шаг 6: Подключиться к сервису

# Пробрасываем порт
```bash
kubectl port-forward svc/hello-k8s 8080:8080
```

### Шаг 7: Тестировать

Открой новый терминал:
```bash
curl localhost:8080
```

# Тестируем health check
```bash
curl localhost:8080/health
```

## Troubleshooting

### Под в статусе ImagePullBackOff

# Образ не загружен в kind
```bash
kind load docker-image hello-k8s:v1 --name learning
```

### Под в статусе CrashLoopBackOff

# Смотри логи
```bash
kubectl logs <pod-name>
```

# Смотри events
```bash
kubectl describe pod <pod-name>
```

## 🔄 Очистка

# Удалить deployment и service
```bash
kubectl delete -f k8s/
```

# Удалить кластер (если нужно)
```bash
kind delete cluster --name learning
```

# Удалить Docker образ
```bash
docker rmi hello-k8s:v1
```
Дата создания: 27 февраля 2026  
Автор: Evatos