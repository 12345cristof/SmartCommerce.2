# SmartCommerce

## Autor

Cristofer Jhonatan Tupiza Cadena

## Asignatura

Programación Orientada a Objetos

## Universidad

Universidad Internacional del Ecuador

## Fecha

Agosto 2026

## Descripción

SmartCommerce es una plataforma de comercio electrónico desarrollada utilizando Go (Golang) y el framework Gin. El sistema permite gestionar usuarios, productos y pedidos mediante servicios web REST, aplicando conceptos de Programación Orientada a Objetos, estructuras de datos y serialización de información mediante JSON.

## Objetivo General

Desarrollar una plataforma de comercio electrónico que permita administrar información de usuarios, productos y pedidos mediante la aplicación de Programación Orientada a Objetos, estructuras de datos y servicios web REST.

## Tecnologías Utilizadas

- Go (Golang)
- Gin Framework
- JSON
- GitHub
- Visual Studio Code
- Postman

## Funcionalidades Principales

- Registro de usuarios
- Consulta de usuarios
- Registro de productos
- Consulta de productos
- Registro de pedidos
- Consulta de pedidos
- Servicio de autenticación
- Servicio de reportes

## Servicios Web Implementados

### Usuarios

POST /api/usuarios

GET /api/usuarios

### Productos

POST /api/productos

GET /api/productos

### Pedidos

POST /api/pedidos

GET /api/pedidos

### Autenticación

POST /api/login

### Reportes

GET /api/reportes

## Estructura del Proyecto

```text
SmartCommerce
│
├── cmd
│   └── main.go
│
├── internal
│   ├── handlers
│   ├── models
│   ├── services
│   └── interfaces
│
├── go.mod
└── README.md
```

## Ejecución

Instalar dependencias:

```bash
go mod tidy
```

Ejecutar aplicación:

```bash
go run cmd/main.go
```

Servidor:

```text
http://localhost:8080
```

## Resultados

El proyecto permitió desarrollar una API REST funcional capaz de administrar usuarios, productos y pedidos mediante intercambio de información en formato JSON, integrando los conocimientos adquiridos durante la asignatura de Programación Orientada a Objetos.
