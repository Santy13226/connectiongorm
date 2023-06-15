-- Tabla Cliente
CREATE TABLE cliente(
    id_cliente SERIAL PRIMARY KEY,
    cedula CHAR(15),
    nombres VARCHAR(100),
    apellidos VARCHAR(100),
    direccion_domicilio VARCHAR(250),
    numero_celular CHAR(15),
    correo_electronico VARCHAR(150),
    contrasena CHAR(30),
    edad INT,
    sexo CHAR(3)
);

-- Tabla Sucursal
CREATE TABLE sucursal(
    id_sucursal SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    direccion VARCHAR(100),
    telefono CHAR(15),
    hora_de_atencion VARCHAR(100)
);

-- Tabla Conversacion
CREATE TABLE conversacion(
    id_conversacion SERIAL PRIMARY KEY,
    id_cliente INT,
    fecha DATE,
    mensajes VARCHAR(150),
    FOREIGN KEY (id_cliente) REFERENCES cliente(id_cliente)
);

-- Tabla Pedido
CREATE TABLE pedido(
    id_pedido SERIAL PRIMARY KEY,
    id_cliente INT,
    id_sucursal INT,
    fecha DATE,
    producto VARCHAR(100),
    cantidad INT,
    precio FLOAT,
    FOREIGN KEY (id_cliente) REFERENCES cliente(id_cliente),
    FOREIGN KEY (id_sucursal) REFERENCES sucursal(id_sucursal)
);

-- Tabla Producto
CREATE TABLE producto(
    id_producto SERIAL PRIMARY KEY,
    item_codigo CHAR(50),
    nombre VARCHAR(100),
    descripcion VARCHAR(100),
    stock INT,
    p_v_p FLOAT,
    id_sucursal INT,
    FOREIGN KEY (id_sucursal) REFERENCES sucursal(id_sucursal)
);
