package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Santy13226/connectiongorm.git/connection"
	"github.com/Santy13226/connectiongorm.git/database/repositories"
	"github.com/Santy13226/connectiongorm.git/models"
)

func errorFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func leerInput(mensaje string) string {
	fmt.Print(mensaje)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}

	// Eliminar el salto de línea del final
	input = strings.TrimSpace(input)

	return input
}

func main() {
	conn, err := connection.GetConnection("localhost", "postgres", "200018S@nty", "chatbot", "5432")
	errorFatal(err)

	// Crear instancia del repositorio de clientes
	clienteRepo := repositories.NewClienteRepository(conn)

	// Crear instancia del repositorio de productos
	productoRepo := repositories.NewProductoRepository(conn)

	// Crear instancia del repositorio de productos
	pedidoRepo := repositories.NewPedidoRepository(conn)

	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Clientes")
		fmt.Println("2. Productos")
		fmt.Println("3. Pedidos")
		fmt.Println("0. Salir")

		opcion := leerInput("Ingrese el número de opción: ")

		switch opcion {
		case "1":
			// Opciones para clientes
			for {
				fmt.Println("Seleccione una opción:")
				fmt.Println("1. Insertar cliente")
				fmt.Println("2. Consultar clientes")
				fmt.Println("3. Actualizar cliente")
				fmt.Println("4. Eliminar cliente")
				fmt.Println("0. Volver")

				clienteOpcion := leerInput("Ingrese el número de opción: ")

				if clienteOpcion == "0" {
					break
				}

				switch clienteOpcion {
				case "1":
					// Insertar cliente
					cliente := &models.Cliente{
						Cedula:             leerInput("Ingrese la cédula: "),
						Nombres:            leerInput("Ingrese los nombres: "),
						Apellidos:          leerInput("Ingrese los apellidos: "),
						DireccionDomicilio: leerInput("Ingrese la dirección de domicilio: "),
						NumeroCelular:      leerInput("Ingrese el número de celular: "),
						CorreoElectronico:  leerInput("Ingrese el correo electrónico: "),
						Contrasena:         leerInput("Ingrese la contraseña: "),
						Edad:               0,
						Sexo:               leerInput("Ingrese el sexo: "),
					}

					// Leer la edad como entero
					edadStr := leerInput("Ingrese la edad: ")
					edad, err := strconv.Atoi(edadStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					cliente.Edad = edad

					err = clienteRepo.RegistrarCliente(cliente)
					errorFatal(err)

				case "2":
					// Consultar clientes
					clientes, err := clienteRepo.ConsultarClientes()
					errorFatal(err)
					for _, c := range clientes {
						fmt.Println(c)
					}

				case "3":
					// Actualizar cliente
					cedula := leerInput("Ingrese la cédula del cliente a actualizar: ")
					clienteActualizado := &models.Cliente{
						Nombres:   leerInput("Ingrese los nuevos nombres: "),
						Apellidos: leerInput("Ingrese los nuevos apellidos: "),
					}

					// Leer la nueva edad como entero
					nuevaEdadStr := leerInput("Ingrese la nueva edad: ")
					nuevaEdad, err := strconv.Atoi(nuevaEdadStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					clienteActualizado.Edad = nuevaEdad

					err = clienteRepo.ActualizarCliente(cedula, clienteActualizado)
					errorFatal(err)

				case "4":
					// Eliminar cliente
					cedula := leerInput("Ingrese la cédula del cliente a eliminar: ")
					err = clienteRepo.EliminarCliente(cedula)
					errorFatal(err)

				default:
					fmt.Println("Opción inválida")
				}

			}

		case "2":
			// Opciones para productos
			for {
				fmt.Println("Seleccione una opción:")
				fmt.Println("1. Insertar producto")
				fmt.Println("2. Consultar productos")
				fmt.Println("3. Actualizar producto")
				fmt.Println("4. Eliminar producto")
				fmt.Println("0. Volver")

				productoOpcion := leerInput("Ingrese el número de opción: ")

				if productoOpcion == "0" {
					break
				}

				switch productoOpcion {
				case "1":
					// Insertar producto
					producto := &models.Producto{
						ItemCodigo:  leerInput("Ingrese el código del producto: "),
						Nombre:      leerInput("Ingrese el nombre del producto: "),
						Descripcion: leerInput("Ingrese la descripción del producto: "),
						Stock:       0,
						PVP:         0,
					}

					// Leer el stock como entero
					stockStr := leerInput("Ingrese el stock: ")
					stock, err := strconv.Atoi(stockStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					producto.Stock = stock

					// Leer el PVP como número de punto flotante
					pvpStr := leerInput("Ingrese el PVP: ")
					pvp, err := strconv.ParseFloat(pvpStr, 64)
					if err != nil {
						log.Fatal(err.Error())
					}
					producto.PVP = pvp

					err = productoRepo.RegistrarProducto(producto)
					errorFatal(err)

				case "2":
					// Consultar productos
					productos, err := productoRepo.ConsultarProductos()
					errorFatal(err)
					for _, p := range productos {
						fmt.Println(p)
					}

				case "3":
					// Actualizar producto
					nombre := leerInput("Ingrese el nombre del producto a actualizar: ")
					productoActualizado := &models.Producto{
						Descripcion: leerInput("Ingrese la nueva descripción del producto: "),
					}

					// Leer el nuevo stock como entero
					nuevoStockStr := leerInput("Ingrese el nuevo stock: ")
					nuevoStock, err := strconv.Atoi(nuevoStockStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					productoActualizado.Stock = nuevoStock

					// Leer el nuevo PVP como número de punto flotante
					nuevoPVPStr := leerInput("Ingrese el nuevo PVP: ")
					nuevoPVP, err := strconv.ParseFloat(nuevoPVPStr, 64)
					if err != nil {
						log.Fatal(err.Error())
					}
					productoActualizado.PVP = nuevoPVP

					err = productoRepo.ActualizarProducto(nombre, productoActualizado)
					errorFatal(err)

				case "4":
					// Eliminar producto
					nombre := leerInput("Ingrese el nombre del producto a eliminar: ")
					err = productoRepo.EliminarProducto(nombre)
					errorFatal(err)

				default:
					fmt.Println("Opción inválida")
				}

			}
		case "3":
			// Opciones para pedidos
			for {
				fmt.Println("Seleccione una opción:")
				fmt.Println("1. Insertar pedido")
				fmt.Println("2. Consultar pedidos")
				fmt.Println("3. Actualizar pedido")
				fmt.Println("4. Eliminar pedido")
				fmt.Println("0. Volver")
		
				pedidoOpcion := leerInput("Ingrese el número de opción: ")
		
				if pedidoOpcion == "0" {
					break
				}
		
				switch pedidoOpcion {
				case "1":
					// Insertar pedido
					pedido := &models.Pedido{
						IDCliente:  2,
						IDSucursal: 1,
						Fecha:      time.Now(),
						Producto:   leerInput("Ingrese el nombre del producto: "),
						Cantidad:   0,
						Precio:     0,
					}
		
					// Leer la cantidad como entero
					cantidadStr := leerInput("Ingrese la cantidad: ")
					cantidad, err := strconv.Atoi(cantidadStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					pedido.Cantidad = cantidad
		
					// Leer el precio como número de punto flotante
					precioStr := leerInput("Ingrese el precio: ")
					precio, err := strconv.ParseFloat(precioStr, 64)
					if err != nil {
						log.Fatal(err.Error())
					}
					pedido.Precio = precio
		
					err = pedidoRepo.RegistrarPedido(pedido)
					errorFatal(err)
		
				case "2":
					// Consultar pedidos
					pedidos, err := pedidoRepo.ConsultarPedido()
					if err != nil {
						log.Fatal(err.Error())
					}
					for _, p := range pedidos {
						fmt.Println(p)
					}
				
				case "3":
					// Actualizar pedido
					IDStr := leerInput("Ingrese el ID del pedido a actualizar: ")
					ID, err := strconv.Atoi(IDStr)
					if err != nil {
						log.Fatal(err.Error())
					}

					cantidadStr := leerInput("Ingrese la nueva cantidad del pedido: ")
					cantidad, err := strconv.Atoi(cantidadStr)
					if err != nil {
						log.Fatal(err.Error())
					}

					pedidoActualizado := &models.Pedido{
						Cantidad: cantidad,
					}

					err = pedidoRepo.ActualizarPedido(ID, pedidoActualizado)
					errorFatal(err)
				
		
				case "4":
					// Eliminar pedido
					IDStr := leerInput("Ingrese el ID del pedido a eliminar: ")
					ID, err := strconv.Atoi(IDStr)
					if err != nil {
						log.Fatal(err.Error())
					}
		
					err = pedidoRepo.EliminarPedido(ID)
					errorFatal(err)
		
				default:
					fmt.Println("Opción inválida")
				}
		
			}
		
		case "0":
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
