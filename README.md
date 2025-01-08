# Challenge Intuit

Este es un proyecto de ejemplo para la gestión de clientes utilizando el framework Echo en Go.

## Requisitos

- Go 1.16 o superior
- Git

## Instalación

Sigue estos pasos para clonar el repositorio, instalar las dependencias y ejecutar el proyecto.

### Clonar el repositorio

```bash
git clone https://github.com/tu-usuario/challenge-intuit.git
cd challenge-intuit

``` Instalar dependencias
- go mod tidy

``` Ejecutar el servidor
- go run server.go

### Estructura del proyecto
El proyecto está organizado de la siguiente manera:

internal/handlers/client: Controladores para manejar las solicitudes relacionadas con los clientes.
internal/models/clients: Modelos de datos para los clientes.
internal/models/domicilio: Modelos de datos para los domicilios.
internal/models/nombres_apellidos: Modelos de datos para los nombres y apellidos.
internal/repositories/client: Repositorios para interactuar con la base de datos.
internal/services/client: Servicios para la lógica de negocio relacionada con los clientes.
internal/routes: Definición de las rutas de la API.
Rutas de la API
Clientes
GET /clients/:id: Obtener cliente mediante el ID.
GET /clients: Obtener todos los clientes.
POST /clients/create: Crear un cliente.

Ejemplo de JSON para crear un cliente
```{
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
}```