package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	//"regexp"
	"strconv"
	"strings"

	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/connection"
	"golang.org/x/crypto/bcrypt"

	//"github.com/Santy13226/connectiongorm.git/CAPA_LOGICA"
	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/database/repositories"
	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/models"
	capalogica "github.com/Santy13226/connectiongorm.git/CAPA_LOGICA"
	//"golang.org/x/crypto/bcrypt"
)

func errorFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func leerInput(mensaje, tipo string) string {
	fmt.Print(mensaje)

	resp := bufio.NewReader(os.Stdin)
	input, err := resp.ReadString('\n')
	// Eliminar el salto de línea del final
	input = strings.TrimSpace(input)
	switch tipo {
	case "opcion":

	case "nombres":
		isValidNombre := capalogica.ValidarNombreApellido(input)
		if !isValidNombre {
			fmt.Println("El o los nombres no son válidos.")
			os.Exit(1)
		}

	case "apellidos":
		isValidNombre := capalogica.ValidarNombreApellido(input)
		if !isValidNombre {
			fmt.Println("El o los apellidos no son válidos.")
			os.Exit(1)
		}

	case "cedula":
		isValidCed := capalogica.ValidarCedula(input)
		if !isValidCed {
			fmt.Println("La cédula no es válido.")
			os.Exit(1)
		}

	case "direccion":
		isValidCed := capalogica.ValidarDireccionDomicilio(input)
		if !isValidCed {
			fmt.Println("La dirección no es válida.")
			os.Exit(1)
		}

	case "ceular":
		isValidCed := capalogica.ValidarNumeroCelular(input)
		if !isValidCed {
			fmt.Println("El celular no es válido.")
			os.Exit(1)
		}

	case "correo":
		isValidCed := capalogica.ValidarCorreoElectronico(input)
		if !isValidCed {
			fmt.Println("El correo no es válido.")
			os.Exit(1)
		}

	case "contraseña":
		isValidCed := capalogica.ValidarContrasena(input)
		if !isValidCed {
			fmt.Println("La contraseña no es válida. La contraseña debe ser mayor a 8 dígitos")
			os.Exit(1)
		}
	case "sexo":
		isValidCed := capalogica.ValidarSexo(input)
		if !isValidCed {
			fmt.Println("El sexo no es válido.")
			os.Exit(1)
		}
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(input)
	return input
}

func leerInputProductos(mensaje, tipo string) string {
	fmt.Print(mensaje)

	resp := bufio.NewReader(os.Stdin)
	input, err := resp.ReadString('\n')
	// Eliminar el salto de línea del final
	input = strings.TrimSpace(input)
	// switch tipo {
	// case "opcion":
	// 	break
	// case "nombre":
	// 	isValidNombre := validarNombreProducto(input)
	// 	if !isValidNombre {
	// 		fmt.Println("El nombre del producto no son válido.")
	// 		os.Exit(1)
	// 	}
	// 	break
	// case "codigo":
	// 	isValidNombre := validarItemCodigo(input)
	// 	if !isValidNombre {
	// 		fmt.Println("El código del producto no es válido.")
	// 		os.Exit(1)
	// 	}
	// 	break

	// case "descripcion":
	// 	isValidCed := validarDescripcion(input)
	// 	if !isValidCed {
	// 		fmt.Println("La descripción del producto no es válida.")
	// 		os.Exit(1)
	// 	}

	// 	break
	// case "stock":
	// 	isValidCed := validarStock(input)
	// 	if !isValidCed {
	// 		fmt.Println("El stock no es válido.")
	// 		os.Exit(1)
	// 	}
	// 	break
	// case "pvp":
	// 	isValidCed := validarPVP(input)
	// 	if !isValidCed {
	// 		fmt.Println("El precio no es válido.")
	// 		os.Exit(1)
	// 	}
	// 	break

	// }

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(input)
	return input
}

func main() {

	server := "localhost"
	database := "chatbot"

	//Obtener la Conexion con POSTGRES
	conn, err := connection.GetConnection(server, "postgres", "200018S@nty", database, "5432")

	// Obtener la Conexión con SQL SERVER
	//conn, err := connection.GetConnectionSQLS(server, database)

	//Obtener la conexion con MYSQL
	//conn, err := connection.GetMySQLConnection("localhost", "root", "", "chatbot", "3306")

	//conn, err := connection.GetMySQLConnection("root", "", database, server, "3306")
	fmt.Println("El servidor esta corriendo en el puerto localhost:3000")

	//------------------------------------------------------------------------------
	if err != nil {
		log.Fatal(err)
	}

	errorFatal(err)

	// Crear instancia del repositorio de clientes
	clienteRepo := repositories.NewClienteRepository(conn)

	// Crear instancia del repositorio de productos
	productoRepo := repositories.NewProductoRepository(conn)

	// Crear instancia del repositorio de productos
	//pedidoRepo := repositories.NewPedidoRepository(conn)

	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Clientes")
		fmt.Println("2. Productos")
		fmt.Println("3. Pedidos")
		fmt.Println("0. Salir")

		opcion := leerInput("Ingrese el número de opción: ", "opcion")

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

				clienteOpcion := leerInput("Ingrese el número de opción: ", "opcion")

				if clienteOpcion == "0" {
					break
				}

				switch clienteOpcion {
				case "1":

					// Insertar cliente
					cliente := &models.Cliente{
						Cedula:             leerInput("Ingrese la cédula: ", "cedula"),
						Nombres:            leerInput("Ingrese la nombres: ", "nombres"),
						Apellidos:          leerInput("Ingrese los apellidos: ", "apellidos"),
						DireccionDomicilio: leerInput("Ingrese la dirección de domicilio: ", "direccion"),
						NumeroCelular:      leerInput("Ingrese el número de celular: ", "celular"),
						CorreoElectronico:  leerInput("Ingrese el correo electrónico: ", "correo"),
						Contrasena:         "",
						Edad:               0,
						Sexo:               leerInput("Ingrese el sexo: ", "sexo"),
					}

					//leer contraseña
					constrasenaStr := leerInput("Ingrese la contraseña: ", "contraseña")
					hashedPassword, err := bcrypt.GenerateFromPassword([]byte(constrasenaStr), bcrypt.DefaultCost)
					if err != nil {
						log.Fatal(err.Error())
					}
					cliente.Contrasena = string(hashedPassword)

					// Leer la edad como entero
					edadStr := leerInput("Ingrese la edad: ", "edad")
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
					cedula := leerInput("Ingrese la cédula del cliente a actualizar: ", "cedula")
					clienteActualizado := &models.Cliente{
						Nombres:   leerInput("Ingrese los nuevos nombres: ", "nombres"),
						Apellidos: leerInput("Ingrese los nuevos apellidos: ", "apellidos"),
					}

					// Leer la nueva edad como entero
					nuevaEdadStr := leerInput("Ingrese la nueva edad: ", "edad")
					nuevaEdad, err := strconv.Atoi(nuevaEdadStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					clienteActualizado.Edad = nuevaEdad

					err = clienteRepo.ActualizarCliente(cedula, clienteActualizado)
					errorFatal(err)

				case "4":
					// Eliminar cliente
					cedula := leerInput("Ingrese la cédula del cliente a eliminar: ", "cedula")
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

				productoOpcion := leerInput("Ingrese el número de opción: ", "opcion")

				if productoOpcion == "0" {
					break
				}

				switch productoOpcion {
				case "1":
					// Insertar producto
					producto := &models.Producto{
						ItemCodigo:  leerInputProductos("Ingrese el código del producto: ", "codigo"),
						Nombre:      leerInputProductos("Ingrese el nombre del producto: ", "nombre"),
						Descripcion: leerInputProductos("Ingrese la descripción del producto: ", "descripcion"),
						Stock:       0,
						PVP:         0,
					}

					// Leer el stock como entero
					stockStr := leerInputProductos("Ingrese el stock: ", "stock")
					stock, err := strconv.Atoi(stockStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					producto.Stock = stock

					// Leer el PVP como número de punto flotante
					pvpStr := leerInputProductos("Ingrese el PVP: ", "pvp")
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
					nombre := leerInput("Ingrese el nombre del producto a actualizar: ", "nombre")
					productoActualizado := &models.Producto{
						Descripcion: leerInput("Ingrese la nueva descripción del producto: ", "descripcion"),
					}

					// Leer el nuevo stock como entero
					nuevoStockStr := leerInput("Ingrese el nuevo stock: ", "stock")
					nuevoStock, err := strconv.Atoi(nuevoStockStr)
					if err != nil {
						log.Fatal(err.Error())
					}
					productoActualizado.Stock = nuevoStock

					// Leer el nuevo PVP como número de punto flotante
					nuevoPVPStr := leerInput("Ingrese el nuevo PVP: ", "pvp")
					nuevoPVP, err := strconv.ParseFloat(nuevoPVPStr, 64)
					if err != nil {
						log.Fatal(err.Error())
					}
					productoActualizado.PVP = nuevoPVP

					err = productoRepo.ActualizarProducto(nombre, productoActualizado)
					errorFatal(err)

				case "4":
					// Eliminar producto
					nombre := leerInput("Ingrese el nombre del producto a eliminar: ", "nombre")
					err = productoRepo.EliminarProducto(nombre)
					errorFatal(err)

				default:
					fmt.Println("Opción inválida")
				}

			}

			/*
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
								IDCliente:  10,
								IDSucursal: 1,
								Fecha:      time.Now(),
								//Producto:   leerInput("Ingrese el nombre del producto: "),
								Cantidad: 0,
								Precio:   0,
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
			*/
		case "0":
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
