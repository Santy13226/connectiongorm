package main

import (
	"fmt"
	"log"

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
	//registrarCliente(conn)
	consultarClientes(conn)
	//registrarSucursal(conn)
	fmt.Println("Conexión exitosa")
}
 //METODO DE REGISTRO
 /*
func registrarCliente(conn *gorm.DB) {
	cliente := &models.Cliente{
		Cedula:             "0604855866",
		Nombres:            "Bryan Santiago",
		Apellidos:          "Guaylla Ashqui",
		DireccionDomicilio: "España y Febres Cordero",
		NumeroCelular:      "0984242647",
		CorreoElectronico:  "santiguaylla@gmail.com",
		Contrasena:         "123456789",
		Edad:               22,
		Sexo:               "MAS",
	}

	result := conn.Create(cliente)
	if result.Error != nil {
		log.Println(result.Error)
	}
}
*/
/*
func registrarSucursal(conn *gorm.DB) {
	p := models.Sucursal{
		Nombre:         "Farmacia",
		Direccion:      "Farmacia sana sana",
		Telefono:       "032-9420242",
		HoraDeAtencion: "08:00-21:00",
	}
	rs := conn.Create(p)
	if rs.Error != nil {
		log.Println(rs.Error)
	}
}
*/
//METODO DE CONSULTA
func consultarClientes(conn *gorm.DB){
	var clientes[]models.Cliente						//Se guarda el registro de clientes usando el modelo
	if err:= conn.Find(&clientes).Error; err != nil{
		log.Println(err.Error())						//manejo de error
	}

	for _, p := range clientes{ 			//Ciclo para imprimir los registros
		fmt.Println(p)
	}
}