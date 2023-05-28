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
	//consultarClientes(conn)
	//consultarClientesPorCedula(conn, "0604855867")
	//registrarSucursal(conn)
	//updateCliente(conn)
	deleteCliente(conn, "0604855866")
	fmt.Println("Conexión exitosa")
}
 //METODO DE REGISTRO
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
	}else{
		log.Println("REGISTRO DE CLIENTE REALIZADO CORRECTAMENTE...")
	}
}


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
	}else{
		log.Println("REGISTRO DE SUCURSAL REALIZADO CORRECTAMENTE...")
	}
}

//METODO DE CONSULTA GENERAL
func consultarClientes(conn *gorm.DB){
	var clientes[]models.Cliente						//Se guarda el registro de clientes usando el modelo
	if err:= conn.Find(&clientes).Error; err != nil{
		log.Println(err.Error())						//manejo de error
	}else{
		log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	}

	for _, p := range clientes{ 			//Ciclo para imprimir los registros
		fmt.Println(p)
	}
}

//METODO DE CONSULTA POR PARÁMETRO
func consultarClientesPorCedula(conn *gorm.DB, Cedula string){
	var cliente models.Cliente						//Se guarda el registro de clientes usando el modelo
	result := conn.Where("cedula=?", Cedula).Limit(1).Find(&cliente)
	if result.Error != nil{
		log.Println(result.Error)
		return
	}else{
		log.Println("CONSULTA REALIZADA CORRECTAMENTE...")
	}

	if result.RowsAffected == 0{                    //Comprobacion si se encuentra el registro=1, si no se encuentra =0
		log.Println("No se encontro el registro", Cedula)
		return
	}
	fmt.Println(cliente)
}

func updateCliente(conn *gorm.DB){
	var cliente=models.Cliente{Nombres: "Bryan Santiago", Apellidos: "Guaylla Ashqui", Edad: 22}
	result := conn.Where("Cedula=?", "0604855866" ).Updates(&cliente)
	if result.Error != nil{
		log.Println(result.Error)
		return
	}else{
		log.Println("REGISTRO ACTUALIZADO CORRECTAMENTE....")
	}
}

func deleteCliente(conn *gorm.DB, Cedula string){
	result := conn.Where("Cedula=?", Cedula).Delete(&models.Cliente{})
	if result.Error != nil{
		log.Println(result.Error)
		return
	}else{
		log.Println("REGISTRO BORRADO CORRECTAMENTE....")
	}
	if result.RowsAffected == 0{                    //Comprobacion si se encuentra el registro=1, si no se encuentra =0
		log.Println("NO SE ELIMINO", Cedula)
		return
	}
}