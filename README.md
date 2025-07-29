# GoStorage
Проект с микросервисной архитектурой, включающий три основных сервиса:
•	API-шлюз (API Gateway) – обеспечивает маршрутизацию запросов и единую точку входа для клиентов.
•	Сервис авторизации (Authentication Service) – отвечает за аутентификацию и управление доступом.
•	Сервис хранилища (Storage Service) - сохраняет пользовательские данные и предоставляет их по запросу

Требования:

Docker

Docker Compose

Запуск проекта

1. Клонируйте репозиторий:
```bash
 git clone https://github.com/DanialKassym/GoStorage.git
```
2. Запустите сервисы с помощью Docker Compose:

```bash
 docker-compose up --build
```
