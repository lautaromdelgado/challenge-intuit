# Challenge Intuit

Este es un proyecto de ejemplo para la gestión de clientes utilizando el framework **Echo** en Go.

## Requisitos

Antes de comenzar, asegúrate de tener instalados los siguientes requisitos:

- **Go** 1.16 o superior.
- **Git**.

## Instalación

Sigue los pasos a continuación para clonar el repositorio, instalar las dependencias y ejecutar el proyecto.

### Clonar el repositorio

Abre una terminal y ejecuta los siguientes comandos:

```bash
# Clonar el repositorio
git clone https://github.com/tu-usuario/challenge-intuit.git

# Acceder al directorio del proyecto
cd challenge-intuit
```

### Instalar dependencias

Ejecuta el siguiente comando para instalar las dependencias necesarias:

```bash
go mod tidy
```

### Ejecutar el servidor

Finalmente, ejecuta el siguiente comando para iniciar el servidor:

```bash
go run server.go
```

## Estructura del proyecto

El proyecto está organizado de la siguiente manera:

```
internal/
├── handlers/
│   └── client/              # Controladores para manejar las solicitudes relacionadas con los clientes
├── models/
│   ├── clients/             # Modelos de datos para los clientes
│   ├── domicilio/           # Modelos de datos para los domicilios
│   └── nombres_apellidos/   # Modelos de datos para los nombres y apellidos
├── repositories/
│   └── client/              # Repositorios para interactuar con la base de datos
├── services/
│   └── client/              # Servicios para la lógica de negocio relacionada con los clientes
└── routes/                  # Definición de las rutas de la API
```

## Rutas de la API

### Clientes

- **GET** `/clients/:id`: Obtener un cliente mediante el ID.
- **GET** `/clients`: Obtener todos los clientes.
- **POST** `/clients/create`: Crear un nuevo cliente.

### Ejemplo de JSON para crear un cliente

A continuación se muestra un ejemplo de cómo debe ser el JSON enviado para crear un cliente:

```json
{
  "nombres_apellidos": {
    "first_name": "Victoria",
    "first_surname": "Boll"
  },
  "fecha_de_nacimiento": "1990-01-01",
  "cuit": "25441770607",
  "domicilio": {
    "calle": "Gordibuena",
    "numero": "4",
    "ciudad": "Resistencia",
    "provincia": "Chaco",
    "codigo_postal": "S355",
    "pais": "Argentina"
  },
  "telefono": "348012347",
  "email": "VickiBoll@gmail.com"
}
```