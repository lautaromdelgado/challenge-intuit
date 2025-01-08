# Challenge Intuit Backend

Este proyecto es un backend desarrollado en **Golang** con el framework **Echo** y utiliza **GORM** como ORM para la interacción con la base de datos MySQL. El propósito del sistema es gestionar clientes, incluyendo sus domicilios y nombres completos, mediante un conjunto de endpoints RESTful.

## **Tecnologías utilizadas**

- **Lenguaje:** Go (Golang)
- **Framework:** Echo
- **ORM:** GORM
- **Base de datos:** MySQL
- **Gestión de dependencias:** Go Modules

## **Estructura del proyecto**

```
backend/
├── cmd
│   └── server.go              # Inicialización del servidor Echo
├── config
│   └── .env                   # Archivo de configuración de variables de entorno
├── database
│   ├── script
│   │   └── database.go        # Conexión y configuración de la base de datos
├── internal
│   ├── handlers
│   │   └── client
│   │       └── client_handlers.go  # Controladores para manejar las solicitudes
│   ├── models
│   │   ├── clients
│   │   │   └── clients_models.go   # Modelos de clientes
│   │   ├── domicilio
│   │   │   └── domicilio_models.go # Modelos de domicilio
│   │   └── nombres_apellidos
│   │       └── nombres_apellidos_models.go  # Modelos de nombres y apellidos
│   ├── repositories
│   │   └── client
│   │       └── client_repositories.go  # Repositorios para acceso a la base de datos
│   ├── services
│   │   └── client
│   │       └── client_services.go      # Lógica de negocio y servicios
│   └── utils
│       └── utils.go                # Funciones auxiliares (validaciones)
├── routes
│   └── routes.go                  # Definición de rutas del servidor
└── go.mod                         # Archivo de módulos de Go
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
        "cuit": "20-12345678-9",
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
- **Detalles de implementación:** Se ha implementado un filtro en la base de datos que excluye a los clientes eliminados, es decir, aquellos cuyo campo `eliminado_el` no es nulo.

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
    "fecha_de_nacimiento": "1990-05-15",
    "cuit": "20-12345678-9",
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
- **Descripción:** Elimina un cliente de manera lógica estableciendo la fecha actual en el campo `eliminado_el`. Esto indica que la cuenta ha sido dada de baja pero permanece en el sistema para propósitos de historial.

#### **Respuestas posibles**
- **200 OK:** Cliente actualizado o eliminado correctamente.
- **400 Bad Request:** Error en los datos enviados o el cliente no existe.

## **Funciones auxiliares (utils)**

Las funciones auxiliares en `utils/utils.go` verifican si las estructuras anidadas están vacías antes de proceder a actualizarlas:

- `IsEmptyDomicilio`: Verifica si los campos de la estructura `Domicilio` están vacíos.
- `IsEmptyNombresApellidos`: Verifica si los campos de la estructura `NombresApellidos` están vacíos.

## **Pruebas**
Para probar el funcionamiento del backend, se recomienda utilizar herramientas como **Postman** o **cURL**. Se incluyen tres ejemplos de JSON para pruebas:

### **1. JSON con todos los campos**
```json
{
  "nombres_apellidos": {
    "first_name": "Ana",
    "first_surname": "López"
  },
  "fecha_de_nacimiento": "1985-03-20",
  "cuit": "23-98765432-1",
  "domicilio": {
    "calle": "Av. Libertador",
    "numero": "1010",
    "ciudad": "Autonoma de Buenos Aires",
    "provincia": "Buenos Aires",
    "codigo_postal": "C100",
    "pais": "Argentina"
  },
  "telefono": "1122334455",
  "email": "ana.lopez@example.com"
}
```

### **2. JSON con pocos cambios**
```json
{
  "telefono": "1231231234"
}
```

### **3. JSON con algunos campos actualizados**
```json
{
  "nombres_apellidos": {
    "first_name": "Luis",
    "first_surname": "Martínez"
  },
  "domicilio": {
    "calle": "San Martín",
    "numero": "505"
  }
}
```

## **Licencia**
Este proyecto está desarrollado por www.linkedin.com/in/lautaromdelgado.

