-- CREAR Y USAR BASE DE DATOS
CREATE DATABASE challenge;
USE challenge;

-- TABLA NOMBRES Y APELLIDOS
CREATE TABLE nombres_apellidos_client (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50) DEFAULT NULL,
    first_surname VARCHAR(50) NOT NULL,
    second_surname VARCHAR(50) DEFAULT NULL
);

-- TABLA DOMICILIO
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

-- TABLA CLIENTES
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

/* DESARROLLADO POR LAUTARO DELGADO */