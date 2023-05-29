package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Santy13226/connectiongorm.git/connection"
	"github.com/Santy13226/connectiongorm.git/models"
	"gorm.io/gorm"
)

func errorFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	conn, err := connection.GetConnection("localhost", "postgres", "200018S@nty", "chatbot", "5432")
	_ = conn

	errorFatal(err)
	//CRUD CLIENTES
	//--------------------------------------------------------------------------
	//registrarCliente(conn)
	consultarClientes(conn)
	//consultarClientesPorCedula(conn, "0604855867")
	//updateCliente(conn)
	//deleteCliente(conn, "0604855866")

	//CRUD PRODUCTOS
	//--------------------------------------------------------------------------
	//registrarProducto(conn)
	//consultarProducto(conn)
	//updateProducto(conn)
	//deleteProducto(conn, "Paracetamol")

	//CRUD PEDIDOS
	//--------------------------------------------------------------------------
	//registrarPedido(conn)
	//consultarPedido(conn)
	//updatePedido(conn)
	//deletePedido(conn, 1)

	fmt.Println("Conexión exitosa")
}

// CRUD CLIENTES
//--------------------------------------------------------------------------
// METODO DE REGISTRO CLIENTES
func registrarCliente(conn *gorm.DB) {
	cliente := &models.Cliente{
		Id:                 2,
		Cedula:             "0604934505",
		Nombres:            "Jaime",
		Apellidos:          "Burgos",
		DireccionDomicilio: "9 de Octubre e Italia",
		NumeroCelular:      "098",
		CorreoElectronico:  "sebas.burgos@gmail.com",
		Contrasena:         "123456789",
		Edad:               22,
		Sexo:               "MAS",
	}

	result := conn.Create(cliente)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("REGISTRO DE CLIENTE REALIZADO CORRECTAMENTE...")
	}
}

// METODO DE CONSULTA GENERAL
func consultarClientes(conn *gorm.DB) {
	var clientes []models.Cliente //Se guarda el registro de clientes usando el modelo
	if err := conn.Find(&clientes).Error; err != nil {
		log.Println(err.Error()) //manejo de error
	} else {
		log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	}

	for _, p := range clientes { //Ciclo para imprimir los registros
		fmt.Println(p)
	}
}

// METODO DE CONSULTA POR PARÁMETRO
func consultarClientesPorCedula(conn *gorm.DB, Cedula string) {
	var cliente models.Cliente //Se guarda el registro de clientes usando el modelo
	result := conn.Where("cedula=?", Cedula).Limit(1).Find(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return
	} else {
		log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	}

	if result.RowsAffected == 0 { //Comprobacion si se encuentra el registro=1, si no se encuentra =0
		log.Println("No se encontro el registro", Cedula)
		return
	}
	fmt.Println(cliente)
}

//MÉTODO DE ACTUALIZACIÓN DE CLIENTES
func updateCliente(conn *gorm.DB) {
	var cliente = models.Cliente{Nombres: "Bryan Santiago", Apellidos: "Guaylla Ashqui", Edad: 22}
	result := conn.Where("Cedula=?", "0604855866").Updates(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return
	} else {
		log.Println("REGISTRO ACTUALIZADO CORRECTAMENTE....")
	}
}

//MÉTODO DE ELIMINACIÓN DE CLIENTES REGISTRADOS
func deleteCliente(conn *gorm.DB, Cedula string) {
	result := conn.Where("Cedula=?", Cedula).Delete(&models.Cliente{})
	if result.Error != nil {
		log.Println(result.Error)
		return
	} else {
		log.Println("REGISTRO BORRADO CORRECTAMENTE....")
	}
	if result.RowsAffected == 0 { //Comprobacion si se encuentra el registro=1, si no se encuentra =0
		log.Println("NO SE ELIMINO", Cedula)
		return
	}
}


// CRUD PRODUCTOS
//--------------------------------------------------------------------------
// Metodo de registro de PRODUCTO
func registrarProducto(conn *gorm.DB) {
	p := models.Producto{
		ID:          1,
		ItemCodigo:  "6822",
		Nombre:      "Paracetamol",
		Descripcion: "Medicamento/Analgésico para el dolor",
		Stock:       20,
		PVP:         10,
		IDSucursal:  1,
	}
	rs := conn.Create(&p)
	if rs.Error != nil {
		log.Println(rs.Error)
	} else {
		log.Println("REGISTRO DE PRODUCTO REALIZADO CORRECTAMENTE...")
	}
}

//Método de consulta de productos
func consultarProducto(conn *gorm.DB) {
	var producto []models.Producto //Se guarda el registro de clientes usando el modelo
	if err := conn.Find(&producto).Error; err != nil {
		log.Println(err.Error()) //manejo de error
	} else {
		log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	}

	for _, p := range producto { //Ciclo para imprimir los registros
		fmt.Println(p)
	}
}
//Método de actualización de productos
func updateProducto(conn *gorm.DB){
	var producto = models.Producto{Descripcion: "Medicamento que sirve como analgésico para el dolor"}
	result := conn.Where("Nombre=?", "Paracetamol").Updates(&producto)
	if result.Error != nil {
		log.Println(result.Error)
		return
	} else {
		log.Println("REGISTRO ACTUALIZADO CORRECTAMENTE....")
	}
}

//Método de eliminación de productos
func deleteProducto(conn *gorm.DB, Nombre string){
	result := conn.Where("Nombre=?", Nombre).Delete(&models.Producto{})
	if result.Error != nil {
		log.Println(result.Error)
		return
	} else {
		log.Println("REGISTRO BORRADO CORRECTAMENTE....")
	}
	if result.RowsAffected == 0 { //Comprobacion si se encuentra el registro=1, si no se encuentra =0
		log.Println("NO SE ELIMINO", Nombre)
		return
	}
}


// CRUD PEDIDOS
//--------------------------------------------------------------------------
// Metodo de registro de PRODUCTO
func registrarPedido(conn *gorm.DB) {
	fecha := time.Date(2023, time.May, 29, 0, 0, 0, 0, time.UTC)
	p := models.Pedido{
	ID         :1,
	IDCliente  :2,
	IDSucursal :1,
	Fecha      :fecha,
	Producto   :"Paracetamol",
	Cantidad   :2,
	Precio     :0.60,
	}
	rs := conn.Create(&p)
	if rs.Error != nil {
		log.Println(rs.Error)
	} else {
		log.Println("REGISTRO DE PEDIDO REALIZADO CORRECTAMENTE...")
	}
}

//Método de consulta de productos
func consultarPedido(conn *gorm.DB) {
	var pedido []models.Pedido //Se guarda el registro de clientes usando el modelo
	if err := conn.Find(&pedido).Error; err != nil {
		log.Println(err.Error()) //manejo de error
	} else {
		log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	}

	for _, p := range pedido { //Ciclo para imprimir los registros
		fmt.Println(p)
	}
}
//Método de actualización de productos
func updatePedido(conn *gorm.DB){
	var pedido = models.Pedido{Cantidad: 10}
	result := conn.Where("id_cliente=?", "2").Updates(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return
	} else {
		log.Println("REGISTRO ACTUALIZADO CORRECTAMENTE....")
	}
}

//Método de eliminación de productos
func deletePedido(conn *gorm.DB, ID int){
	result := conn.Where("id_pedido=?", ID).Delete(&models.Pedido{})
	if result.Error != nil {
		log.Println(result.Error)
		return
	} else {
		log.Println("REGISTRO BORRADO CORRECTAMENTE....")
	}
	if result.RowsAffected == 0 { //Comprobacion si se encuentra el registro=1, si no se encuentra =0
		log.Println("NO SE ELIMINO", ID)
		return
	}
}
