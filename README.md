# Challenge Intuit Backend

Este proyecto es un backend desarrollado en **Golang** utilizando el framework **Echo** y **GORM** como ORM para la gestión de la base de datos MySQL. El sistema permite la gestión de clientes, sus domicilios y sus nombres completos a través de un conjunto de endpoints RESTful.

## **Tecnologías utilizadas**

- **Lenguaje:** Go (Golang)
- **Framework:** Echo
- **ORM:** GORM
- **Base de datos:** MySQL
- **Gestor de dependencias:** Go Modules

## **Estructura del proyecto**

```plaintext
backend/
├── cmd
│   └── server.go                         # Inicialización del servidor Echo
├── config
│   └── .env                              # Variables de entorno
├── database
│   └── script
│       └── database.go                   # Conexión y configuración de la base de datos
├── internal
│   ├── handlers
│   │   └── client
│   │       └── client_handlers.go        # Controladores de las solicitudes REST
│   ├── models
│   │   ├── clients
│   │   │   └── clients_models.go         # Modelos de clientes
│   │   ├── domicilio
│   │   │   └── domicilio_models.go       # Modelos de domicilio
│   │   └── nombres_apellidos
│   │       └── nombres_apellidos_models.go  # Modelos de nombres y apellidos
│   ├── repositories
│   │   └── client
│   │       └── client_repositories.go    # Repositorios de clientes
│   ├── services
│   │   └── client
│   │       └── client_services.go        # Lógica de negocio y servicios
│   └── utils
│       └── UpdateClient.go               # Funciones auxiliares para actualizaciones
├── routes
│   ├── clients_routes.go                 # Definición de rutas de clientes
│   └── routes.go                         # Configuración de rutas del servidor
├── go.mod                                # Archivo de módulos de Go
├── go.sum                                # Suma de verificación de módulos
└── README.md                             # Documentación del proyecto
```

## **Estructura de la base de datos**

```sql
-- CREAR Y USAR BASE DE DATOS
CREATE DATABASE challenge;
USE challenge;

-- TABLA nombres_apellidos_client
CREATE TABLE nombres_apellidos_client (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50) DEFAULT NULL,
    first_surname VARCHAR(50) NOT NULL,
    second_surname VARCHAR(50) DEFAULT NULL
);

-- TABLA domicilio
CREATE TABLE domicilio (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    calle VARCHAR(55) NOT NULL,
    numero VARCHAR(10) NOT NULL,
    piso VARCHAR(10) DEFAULT NULL,
    departamento VARCHAR(10) DEFAULT NULL,
    ciudad VARCHAR(50) NOT NULL,
    provincia VARCHAR(50) NOT NULL,
    codigo_postal VARCHAR(10) NOT NULL,
    pais VARCHAR(50) NOT NULL
);

-- TABLA clientes
CREATE TABLE clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_nombres_apellidos INT UNIQUE,
    fecha_de_nacimiento DATE NOT NULL,
    cuit BIGINT NOT NULL,
    id_domicilio INT UNIQUE,
    telefono VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    FOREIGN KEY (id_nombres_apellidos) REFERENCES nombres_apellidos_client(id),
    FOREIGN KEY (id_domicilio) REFERENCES domicilio(id)
);

-- AGREGAR COLUMNA DE ALTA Y BAJA
ALTER TABLE clientes ADD COLUMN creado_el TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE clientes ADD COLUMN eliminado_el TIMESTAMP DEFAULT NULL;
```

## **Instalación y ejecución**

1. **Clonar el repositorio:**
   ```bash
   git clone <URL_DEL_REPOSITORIO>
   cd backend
   ```

2. **Configurar las variables de entorno:**
   Crea un archivo `.env` en la carpeta `config` con el siguiente contenido:
   ```env
   DB_USER=<usuario_db>
   DB_PASSWORD=<password_db>
   DB_NAME=<nombre_db>
   DB_HOST=<host_db>
   DB_PORT=<puerto_db>
   ```

3. **Instalar las dependencias:**
   ```bash
   go mod tidy
   ```

4. **Ejecutar el servidor:**
   ```bash
   go run cmd/server.go
   ```

   El servidor estará disponible en `http://localhost:7001`.

## **Endpoints disponibles**

### **Clientes**

#### **Buscar clientes por nombre**
- **URL:** `/clients/search`
- **Método:** `GET`
- **Descripción:** Busca clientes cuyo nombre o apellido coincidan parcialmente con el parámetro de búsqueda proporcionado.
- **Query Param:**
  - `name`: Texto a buscar en los nombres o apellidos.
- **Ejemplo:**
  ```
  GET /clients/search?name=Juan
  ```
- **Respuesta:**
  ```json
  {
    "status": "success",
    "data": [
      {
        "id": 1,
        "nombres_apellidos": {
          "first_name": "Juan",
          "second_name": "Carlos",
          "first_surname": "Pérez",
          "second_surname": "García"
        },
        "fecha_de_nacimiento": "1990-05-15",
        "cuit": "20123456789",
        "domicilio": {
          "calle": "Figueroa Alcorta",
          "numero": "4880",
          "ciudad": "Autonoma de Buenos Aires",
          "provincia": "Buenos Aires",
          "codigo_postal": "S100",
          "pais": "Argentina"
        },
        "telefono": "1234567890",
        "email": "juan.perez@example.com",
        "creado_el": "2025-01-01T12:00:00Z"
      }
    ]
  }
  ```

#### **Obtener un cliente por ID**
- **URL:** `/clients/:id`
- **Método:** `GET`
- **Descripción:** Obtiene un cliente específico mediante su ID.

#### **Obtener todos los clientes**
- **URL:** `/clients`
- **Método:** `GET`
- **Descripción:** Devuelve una lista de todos los clientes registrados que no han sido eliminados.

#### **Obtener clientes eliminados**
- **URL:** `/clients/deleted`
- **Método:** `GET`
- **Descripción:** Devuelve una lista de todos los clientes que han sido eliminados lógicamente.
- **Detalles de implementación:** Se ha implementado un filtro en la base de datos que incluye solo a los clientes eliminados, es decir, aquellos cuyo campo `eliminado_el` no es nulo.

- **Respuesta:**
  ```json
  {
    "status": "success",
    "data": [
      {
        "id": 2,
        "nombres_apellidos": {
          "first_name": "María",
          "first_surname": "González"
        },
        "fecha_de_nacimiento": "1988-11-20",
        "cuit": "27987654321",
        "domicilio": {
          "calle": "Av. Belgrano",
          "numero": "123",
          "ciudad": "Rosario",
          "provincia": "Santa Fe",
          "codigo_postal": "S200",
          "pais": "Argentina"
        },
        "telefono": "9876543210",
        "email": "maria.gonzalez@example.com",
        "eliminado_el": "2025-01-05T15:30:00Z"
      }
    ],
    "totalClients": {
      "totaldeleted": 1
    }
  }

#### **Crear un cliente**
- **URL:** `/clients/create`
- **Método:** `POST`
- **Descripción:** Crea un nuevo cliente.
- **Body:**
  ```json
  {
    "nombres_apellidos": {
      "first_name": "Juan",
      "second_name": "",
      "first_surname": "Pérez",
      "second_surname": ""
    },
    "fecha_de_nacimiento": "2002-03-25T15:04:05Z",
    "cuit": "20123456789",
    "domicilio": {
      "calle": "Figueroa Alcorta",
      "numero": "4880",
      "ciudad": "Autonoma de Buenos Aires",
      "provincia": "Buenos Aires",
      "codigo_postal": "S100",
      "pais": "Argentina"
    },
    "telefono": "1234567890",
    "email": "juan.perez@example.com"
  }
  ```

#### **Actualizar un cliente**
- **URL:** `/clients/update/:id`
- **Método:** `PUT`
- **Descripción:** Actualiza un cliente específico mediante su ID.
- **Body:** (Ejemplo con algunos campos actualizados)
  ```json
  {
    "nombres_apellidos": {
      "first_name": "Carlos",
      "first_surname": "Gómez"
    },
    "telefono": "9876543210"
  }
  ```

#### **Eliminar un cliente**
- **URL:** `/clients/delete/:id`
- **Método:** `PUT`
- **Descripción:** Elimina un cliente de manera lógica estableciendo la fecha actual en el campo `eliminado_el`.

## **Licencia**
Este proyecto está desarrollado por [Lautaro M. Delgado](https://www.linkedin.com/in/lautaromdelgado/).

