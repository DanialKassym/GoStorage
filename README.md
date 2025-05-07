# GoStorage
**This project is currently in progress.**  
A simple, S3-inspired object storage system written in Go. This system mimics the basic functionalities of AWS S3, offering features such as object storage, retrieval, and deletion.

## Requirements
- Docker and Docker Compose
- Go (Golang) for running local in environment 
- .env file for running in local environment configuration

## Setup for Local Development
- Add .env file with:
- DB_URL: postgres://YOURPOSTGRESUSER:changemepass@localhost:5432/YOURDBNAME?sslmode=disable
