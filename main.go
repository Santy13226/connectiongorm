package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"strconv"

	//"net/http/httptest"

	"time"

	//"errors"

	//"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"

	//"io/ioutil"
	"log"
	"net/http"
	"os"

	//"regexp"s

	"strings"

	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/connection"
	"golang.org/x/crypto/bcrypt"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"

	//"google.golang.org/grpc"

	//"github.com/Santy13226/connectiongorm.git/CAPA_LOGICA"
	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/database/repositories"
	"github.com/Santy13226/connectiongorm.git/CAPA_DE_DATOS/models"
	capalogica "github.com/Santy13226/connectiongorm.git/CAPA_LOGICA"
	//"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	IdUsuario int
	Email     string
	Nombres   string
	Cedula    string
}

func errorFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func leerInputProductos(mensaje, tipo string) string {
	fmt.Print(mensaje)

	resp := bufio.NewReader(os.Stdin)
	input, err := resp.ReadString('\n')
	// Eliminar el salto de línea del final
	input = strings.TrimSpace(input)
	switch tipo {
	case "opcion":
		break
	case "nombre":
		isValidNombre := capalogica.ValidarNombreProducto(input)
		if !isValidNombre {
			fmt.Println("El nombre del producto no son válido.")
			return leerInputProductos(mensaje, tipo)
		}
		break
	case "codigo":
		isValidNombre := capalogica.ValidarItemCodigo(input)
		if !isValidNombre {
			fmt.Println("El código del producto no es válido. Debe tener 6 dígitos")
			return leerInputProductos(mensaje, tipo)
		}
		break

	case "descripcion":
		isValidCed := capalogica.ValidarDescripcion(input)
		if !isValidCed {
			fmt.Println("La descripción del producto no es válida.")
			return leerInputProductos(mensaje, tipo)
		}

		break
	case "stock":
		isValidCed := capalogica.ValidarStock(input)
		if !isValidCed {
			fmt.Println("El stock no es válido.")
			return leerInputProductos(mensaje, tipo)
		}
		break
	case "pvp":
		isValidCed := capalogica.ValidarPVP(input)
		if !isValidCed {
			fmt.Println("El precio no es válido.")
			return leerInputProductos(mensaje, tipo)
		}
		break

	}

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(input)
	return input
}

///////////////////TEMAPLATES

type Clientes struct {
	Cedula   string
	Nombre   string
	Apellido string
	Correo   string
}

func Index(rw http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("CAPA_DE_PRESENTACION/templates/index.html")

	//clientesBase := Clientes{nombre, apellido, edad}
	//clientes := Clientes{"David", "Vallejo", 21}

	if err != nil {
		panic((err))
	} else {
		template.Execute(rw, nil) //envio de datos
	}
}

func Login(rw http.ResponseWriter, r *http.Request) {
	checkSession(rw, r)
	template, err := template.ParseFiles("CAPA_DE_PRESENTACION/templates/clientes/login.html")

	//clientesBase := Clientes{nombre, apellido, edad}
	//clientes := Clientes{"David", "Vallejo", 21}

	if err != nil {
		panic((err))
	} else {
		template.Execute(rw, nil) //envio de datos
	}

}

func MostrarCliente(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Ingreso a la funcion mostrar cliente")
	// Obtener el ID del cliente de los parámetros de la URL
	cedula := r.URL.Query().Get("id")

	conn, err := connection.GetElephantSQLConnection()
	errorFatal(err)

	// Crear instancia del repositorio de clientes
	clienteRepo := repositories.NewClienteRepository(conn)

	// Realizar la lógica necesaria para obtener el cliente por su ID
	cliente, err := clienteRepo.ConsultarClientePorCedula(cedula)
	errorFatal(err)
	fmt.Println(cliente)

	if cliente == nil {
		// El cliente no existe, puedes mostrar un mensaje de error o redirigir al index
		http.Redirect(rw, r, "/error", http.StatusFound)
		return
	}

	// Renderizar el template "clientes.html" y enviar los datos del cliente
	template, err := template.ParseFiles("CAPA_DE_PRESENTACION/templates/clientes/cliente.html")
	if err != nil {
		panic(err)
	}

	template.Execute(rw, cliente)
}

type Conversacion struct {
	Pregunta  string
	Respuesta string
}

var sessionID string

var conversaciones []Conversacion

func EnviarPregunta(rw http.ResponseWriter, r *http.Request) {

	projectID := "newagent-lpsi"
	//sessionID := generateSessionID()
	//text := "Hola,tengo dolor de garganta"
	languageCode := "es"
	credentialPath := "newagent-lpsi-00e02bf2e2ec.json"
	creds, err := loadCredentials(credentialPath)
	if err != nil {
		log.Fatalf("Error al cargar las credenciales: %v", err)
	}
	if err != nil {
		log.Fatalf("Error al detectar la intención: %v", err)
	}

	// Obtén la pregunta del usuario
	pregunta := r.FormValue("pregunta")

	fulfillmentText, err := DetectIntentText(projectID, sessionID, pregunta, languageCode, creds)

	// Obtén la respuesta del chatbot
	respuesta := fulfillmentText // Código para obtener la respuesta del chatbot

	// Almacena la conversación en la variable global
	conversaciones = append(conversaciones, Conversacion{
		Pregunta:  pregunta,
		Respuesta: respuesta,
	})

	fmt.Println(conversaciones)
	// Renderiza solo la lista de conversaciones en formato JSON
	conversacionesJSON, err := json.Marshal(conversaciones)
	if err != nil {
		panic(err)
	}

	// Envía la respuesta como JSON
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(conversacionesJSON)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// DESDE AQUI EMPIEZA EL TOKEN DE INICIO DE SESIÓN
type VerifyTokenRequest struct {
	ClientID string `json:"clientID"`
	Token    string `json:"token"`
}

type UserInfo struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	ProfilePicURL string `json:"profilePicURL"`
}

// Agrega una nueva estructura para el token
type TokenInfo struct {
	Name          string `json:"name"`
	ProfilePicURL string `json:"profilePicURL"`
}

func handleVerifyToken(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud
	var req VerifyTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar el token y obtener la información del usuario
	userInfo, err := verifyToken(req.ClientID, req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Crear objeto TokenInfo con la información obtenida
	tokenInfo := &TokenInfo{
		Name:          userInfo.Name,
		ProfilePicURL: userInfo.ProfilePicURL,
	}

	// Devolver la información del token como respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenInfo)

	// Crear objeto UserInfo con la información obtenida
	userInfoResponse := struct {
		Name          string `json:"name"`
		ProfilePicURL string `json:"profilePicURL"`
	}{
		Name:          userInfo.Name,
		ProfilePicURL: userInfo.ProfilePicURL,
	}

	// Enviar los datos al archivo JavaScript
	responseBytes, err := json.Marshal(userInfoResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Imprimir los datos en la consola del servidor
	fmt.Println(userInfo)

	// Escribir los datos en la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}

func verifyToken(clientID, token string) (*UserInfo, error) {
	ctx := context.Background()

	// Verificar el token de ID
	payload, err := idtoken.Validate(ctx, token, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	// Obtener información del usuario del payload
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	profilePicURL, _ := payload.Claims["picture"].(string)

	// Crear objeto UserInfo con la información obtenida
	userInfo := &UserInfo{
		Email:         email,
		Name:          name,
		ProfilePicURL: profilePicURL,
	}

	return userInfo, nil
}

// HASTA AQUI ES LA IMPLEMENTACION DEL TOKEN DE GOOGLE
// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func registrarConversacionHandler(rw http.ResponseWriter, r *http.Request) error {
	conn, err := connection.GetLocalConnection()
	errorFatal(err)

	conversacionRepo := repositories.NewConversacionRepository(conn)
	cookie, err := r.Cookie("session")
	if err != nil {
		// Manejar el error si no se puede recuperar la cookie
	}

	cookieValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		// Manejar el error si no se puede decodificar la cookie
	}

	var userData UserData
	fmt.Println(userData.IdUsuario)

	err = json.Unmarshal(cookieValue, &userData)
	if err != nil {
		// Manejar el error si no se puede deserializar la cookie
	}

	// Crear una slice para almacenar todos los mensajes de las conversaciones
	var mensajes []string

	// Concatenar todos los mensajes en una sola cadena
	for _, conversacion := range conversaciones {
		mensajes = append(mensajes, conversacion.Pregunta, conversacion.Respuesta)
	}

	// Crear una instancia del modelo Conversacion con los datos necesarios
	conversacionDB := &models.Conversacion{
		IDCliente: userData.IdUsuario, // Llena el ID del cliente utilizando la función obtenerIDCliente()
		Fecha:     time.Now(),
		Mensajes:  mensajes,
	}

	// Registrar la x|conversación en la base de datos utilizando el repositorio ConversacionRepository
	idConversacion, err := conversacionRepo.RegistrarConversacion(conversacionDB)
	if err != nil {
		return err
	}

	// Construir la URL de redirección con el ID de la conversación
	redireccionURL := fmt.Sprintf("/chatbot/pharmacybot/conversacion/?id=%d", idConversacion)

	// Redirigir al usuario a la nueva ruta
	http.Redirect(rw, r, redireccionURL, http.StatusSeeOther)

	return nil
}

func actualizarConversacionHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el ID de conversación de la URL
	idParametro := r.URL.Query().Get("id")
	if idParametro == "" {
		http.Error(rw, `{"error": "ID de conversación no proporcionado"}`, http.StatusBadRequest)
		return
	}

	idConversacion, err := strconv.Atoi(idParametro)
	if err != nil {
		http.Error(rw, `{"error": "Error al convertir el ID a un entero"}`, http.StatusBadRequest)
		return
	}

	conn, err := connection.GetLocalConnection()
	if err != nil {
		http.Error(rw, `{"error": "Error al obtener la conexión"}`, http.StatusInternalServerError)
		return
	}

	conversacionRepo := repositories.NewConversacionRepository(conn)
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(rw, `{"error": "Error al obtener la cookie"}`, http.StatusInternalServerError)
		return
	}

	cookieValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		http.Error(rw, `{"error": "Error al decodificar la cookie"}`, http.StatusInternalServerError)
		return
	}

	var userData UserData
	err = json.Unmarshal(cookieValue, &userData)
	if err != nil {
		http.Error(rw, `{"error": "Error al deserializar la cookie"}`, http.StatusInternalServerError)
		return
	}

	// Obtén la conversación actual por su ID
	conversacionActual, err := conversacionRepo.ConsultarConversacionPorID(idConversacion)
	if err != nil {
		http.Error(rw, `{"error": "Error al obtener la conversación actual"}`, http.StatusInternalServerError)
		return
	}

	if conversacionActual == nil {
		http.Error(rw, `{"error": "No se encontró la conversación con el ID proporcionado"}`, http.StatusNotFound)
		return
	}

	// Agrega los nuevos mensajes al slice de mensajes existente
	var mensajes []string
	mensajes = append(mensajes, conversacionActual.Mensajes...)
	for _, conversacion := range conversaciones {
		mensajes = append(mensajes, conversacion.Pregunta, conversacion.Respuesta)
	}

	// Crea una nueva instancia del modelo Conversacion con los datos actualizados
	conversacionDB := &models.Conversacion{
		IDCliente: userData.IdUsuario,
		Fecha:     time.Now(),
		Mensajes:  mensajes,
	}

	// Actualiza la conversación en la base de datos utilizando el repositorio ConversacionRepository
	err = conversacionRepo.ActualizarConversacion(idConversacion, conversacionDB)
	if err != nil {
		http.Error(rw, `{"error": "Error al actualizar la conversación"}`, http.StatusInternalServerError)
		return
	}

	// Si la actualización es exitosa, devuelves un mensaje de éxito en formato JSON
	successMessage := `{"message": "Conversación actualizada correctamente"}`
	rw.Write([]byte(successMessage))
}

func eliminarconversacionHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el ID de conversación de la URL
	idParametro := r.URL.Query().Get("id")
	if idParametro == "" {
		http.Error(rw, `{"error": "ID de conversación no proporcionado"}`, http.StatusBadRequest)
		return
	}

	idConversacion, err := strconv.Atoi(idParametro)
	if err != nil {
		http.Error(rw, `{"error": "Error al convertir el ID a un entero"}`, http.StatusBadRequest)
		return
	}

	conn, err := connection.GetLocalConnection()
	if err != nil {
		http.Error(rw, `{"error": "Error al obtener la conexión"}`, http.StatusInternalServerError)
		return
	}

	conversacionRepo := repositories.NewConversacionRepository(conn)
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(rw, `{"error": "Error al obtener la cookie"}`, http.StatusInternalServerError)
		return
	}

	cookieValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		http.Error(rw, `{"error": "Error al decodificar la cookie"}`, http.StatusInternalServerError)
		return
	}

	var userData UserData
	err = json.Unmarshal(cookieValue, &userData)
	if err != nil {
		http.Error(rw, `{"error": "Error al deserializar la cookie"}`, http.StatusInternalServerError)
		return
	}

	// Obtén la conversación actual por su ID
	err = conversacionRepo.EliminarConversacion(idConversacion)
	if err != nil {
		http.Error(rw, `{"error": "Error al eliminar la conversación"}`, http.StatusInternalServerError)
		return
	}

	// Si la actualización es exitosa, devuelves un mensaje de éxito en formato JSON
	successMessage := `{"message": "Conversación eliminada correctamente"}`
	rw.Write([]byte(successMessage))
}

type Mensaje struct {
	ID              int
	MensajesChatbot string
}

// Define el tipo ConversacionesChatbot con las etiquetas gorm adecuadas
type MensajesText []string
type ConversacionesChatbot struct {
	ID_CONVERSACION int
	ID_USER         int
	Mensajes        MensajesText
}

func MostrarChatBot(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		// Manejar el error si no se puede recuperar la cookie
	}

	cookieValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		// Manejar el error si no se puede decodificar la cookie
	}

	conn, err := connection.GetLocalConnection()
	errorFatal(err)

	// Crear instancia del repositorio de conversaciones
	conversacionRepo := repositories.NewConversacionRepository(conn)

	// Realizar la lógica necesaria para obtener las conversaciones
	mensajes, err := conversacionRepo.ConsultarConversaciones()
	errorFatal(err)
	fmt.Println(mensajes)

	if mensajes == nil {
		// Si no hay conversaciones, puedes mostrar un mensaje de error o redirigir al index
		http.Redirect(rw, r, "/error", http.StatusFound)
		return
	}

	// Crea un slice de ConversacionesChatbot y almacena las conversaciones en él
	var conversacionesChatBot []ConversacionesChatbot
	for _, conversacion := range mensajes {
		var mensajesSeparados []string
		for _, mensaje := range conversacion.Mensajes {
			mensajesSeparados = append(mensajesSeparados, mensaje)
		}
		conversacionChatbot := ConversacionesChatbot{
			ID_CONVERSACION: conversacion.IDConversacion,
			ID_USER:         conversacion.IDCliente,
			Mensajes:        mensajesSeparados,
		}
		conversacionesChatBot = append(conversacionesChatBot, conversacionChatbot)
	}

	var userData UserData
	fmt.Println(conversacionesChatBot)

	err = json.Unmarshal(cookieValue, &userData)
	if err != nil {
		// Manejar el error si no se puede deserializar la cookie
	}

	template, err := template.ParseFiles("CAPA_DE_PRESENTACION/templates/chatbot/pharmacybot.html")
	if err != nil {
		panic(err)
	} else {
		conversaciones = make([]Conversacion, 0)
		contexto := struct {
			UserData              UserData
			Conversaciones        []Conversacion
			ConversacionesChatbot []ConversacionesChatbot
		}{
			UserData:              userData,
			Conversaciones:        conversaciones,
			ConversacionesChatbot: conversacionesChatBot,
		}

		template.Execute(rw, contexto) // Envío de datos al template
	}
}

func MostrarChatBotConversacion(rw http.ResponseWriter, r *http.Request) {
	idParametro := r.URL.Query().Get("id")
	if idParametro == "" {
		// Manejar el caso cuando no se proporciona el ID de la conversación
		http.Redirect(rw, r, "/error", http.StatusFound)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		// Manejar el error si no se puede recuperar la cookie
	}

	cookieValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		// Manejar el error si no se puede decodificar la cookie
	}

	conn, err := connection.GetLocalConnection()
	errorFatal(err)

	// Crear instancia del repositorio de conversaciones
	conversacionRepo := repositories.NewConversacionRepository(conn)

	idConversacion, err := strconv.Atoi(idParametro)
	if err != nil {
		// Manejar el error si no se puede convertir el ID a un entero
		http.Redirect(rw, r, "/error", http.StatusFound)
		return
	}

	// Realizar la lógica necesaria para obtener las conversaciones
	mensajes, err := conversacionRepo.ConsultarConversaciones()
	errorFatal(err)
	fmt.Println(mensajes)

	if mensajes == nil {
		// Si no hay conversaciones, puedes mostrar un mensaje de error o redirigir al index
		http.Redirect(rw, r, "/error", http.StatusFound)
		return
	}

	var conversacionesChatBotTodos []ConversacionesChatbot
	for _, conversacion := range mensajes {
		var mensajesSeparados []string
		for _, mensaje := range conversacion.Mensajes {
			mensajesSeparados = append(mensajesSeparados, mensaje)
		}
		conversacionChatbot := ConversacionesChatbot{
			ID_CONVERSACION: conversacion.IDConversacion,
			ID_USER:         conversacion.IDCliente,
			Mensajes:        mensajesSeparados,
		}
		conversacionesChatBotTodos = append(conversacionesChatBotTodos, conversacionChatbot)
	}

	// Realizar la lógica necesaria para obtener la conversación
	conversacion, err := conversacionRepo.ConsultarConversacionPorID(idConversacion)
	errorFatal(err)
	fmt.Println(conversacion)

	if conversacion == nil {
		// Si no se encontró la conversación, puedes mostrar un mensaje de error o redirigir al index
		http.Redirect(rw, r, "/error", http.StatusFound)
		return
	}

	// Convertir los mensajes en un slice de strings
	var mensajesSeparados []string
	for _, mensaje := range conversacion.Mensajes {
		mensajesSeparados = append(mensajesSeparados, mensaje)
	}

	// Crea un slice de ConversacionesChatbot y almacena la conversación en él
	conversacionChatbot := ConversacionesChatbot{
		ID_CONVERSACION: conversacion.IDConversacion,
		ID_USER:         conversacion.IDCliente,
		Mensajes:        mensajesSeparados,
	}
	conversacionesChatBot := []ConversacionesChatbot{conversacionChatbot}

	var userData UserData
	fmt.Println(conversacionesChatBot)

	err = json.Unmarshal(cookieValue, &userData)
	if err != nil {
		// Manejar el error si no se puede deserializar la cookie
	}

	template, err := template.ParseFiles("CAPA_DE_PRESENTACION/templates/chatbot/pharmacybot/conversacion.html")
	if err != nil {
		panic(err)
	} else {
		conversaciones = make([]Conversacion, 0)
		contexto := struct {
			UserData                   UserData
			Conversaciones             []Conversacion
			ConversacionesChatbot      []ConversacionesChatbot
			ConversacionesChatbotTodos []ConversacionesChatbot // Cambio de nombre aquí
			IDConversacion             int
		}{
			UserData:                   userData,
			Conversaciones:             conversaciones, // Puedes dejar esto como nil si no se necesita en el template
			ConversacionesChatbot:      conversacionesChatBot,
			ConversacionesChatbotTodos: conversacionesChatBotTodos, // Cambio de nombre aquí
			IDConversacion:             idConversacion,
		}

		template.Execute(rw, contexto) // Envío de datos al template
	}
}

func Registro() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		checkSession(rw, r)
		fmt.Println("Ingreso al template registro")
		templateFile := "CAPA_DE_PRESENTACION/templates/clientes/registro.html"

		errorType := r.URL.Query().Get("error")

		data := struct {
			ShowError    bool
			ErrorMessage string
		}{
			ShowError:    errorType != "",
			ErrorMessage: getErrorMessage(errorType),
		}

		fmt.Println(data)
		template, err := template.ParseFiles(templateFile)
		if err != nil {
			panic(err)
		}
		err = template.Execute(rw, data) // Send data to the template
		if err != nil {
			panic(err)
		}
	}
}

func formularioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ingreso al formulario registro")

	localConn, err := connection.GetLocalConnection()
	errorFatal(err)
	elephantSQLConn, err := connection.GetElephantSQLConnection()
	errorFatal(err)

	clienteRepoLocal := repositories.NewClienteRepository(localConn)
	clienteRepoElephantSQL := repositories.NewClienteRepository(elephantSQLConn)

	if r.Method == "POST" {
		// Procesar los datos del formulario enviado

		cedula := r.FormValue("cedula")
		nombres := r.FormValue("nombres")
		apellidos := r.FormValue("apellidos")
		direccion := r.FormValue("direccion")
		celular := r.FormValue("celular")
		email := r.FormValue("email")
		password := r.FormValue("password")
		fechaNac := r.FormValue("fechaNac")
		sexo := r.FormValue("sexo")

		// Parsear la fecha de nacimiento
		fechaNacParsed, err := time.Parse("2006-01-02", fechaNac)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(cedula)
		fmt.Println(nombres)

		if err != nil {
			log.Fatal(err.Error())
		}

		if capalogica.ValidarCedula(cedula) &&
			capalogica.ValidarNombreApellido(nombres) &&
			capalogica.ValidarNombreApellido(apellidos) &&
			capalogica.ValidarDireccionDomicilio(direccion) &&
			capalogica.ValidarNumeroCelular(celular) &&
			capalogica.ValidarCorreoElectronico(email) &&
			capalogica.ValidarContrasena(password) &&
			capalogica.ValidarEdad(fechaNacParsed) &&
			capalogica.ValidarSexo(sexo) {

			fmt.Println("Ingreso a la condicion de validacion")

			// Verificar si el cliente ya existe en la primera base de datos
			existeClienteLocal, err := clienteRepoLocal.ClienteExistente(cedula, email, celular)
			if err != nil {
				log.Fatal(err.Error())
			}

			// Verificar si el cliente ya existe en la segunda base de datos
			existeClienteElephantSQL, err := clienteRepoElephantSQL.ClienteExistente(cedula, email, celular)
			if err != nil {
				log.Fatal(err.Error())
			}

			var redirectURL string

			if existeClienteLocal == "cedula" || existeClienteElephantSQL == "cedula" {
				redirectURL = "/clientes/registro?error=cedula"
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
				return
			} else if existeClienteLocal == "celular" || existeClienteElephantSQL == "celular" {
				redirectURL = "/clientes/registro?error=celular"
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
				return
			} else if existeClienteLocal == "email" || existeClienteElephantSQL == "email" {
				redirectURL = "/clientes/registro?error=email"
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
				return
			}

			// Insertar cliente en la primera base de datos
			clienteLocal := &models.Cliente{
				Cedula:             cedula,
				Nombres:            nombres,
				Apellidos:          apellidos,
				DireccionDomicilio: direccion,
				NumeroCelular:      celular,
				CorreoElectronico:  email,
				Contrasena:         "",
				FechaNacimiento:    fechaNacParsed,
				Sexo:               sexo,
			}

			// Insertar cliente en la segunda base de datos
			clienteElephantSQL := &models.Cliente{
				Cedula:             cedula,
				Nombres:            nombres,
				Apellidos:          apellidos,
				DireccionDomicilio: direccion,
				NumeroCelular:      celular,
				CorreoElectronico:  email,
				Contrasena:         "",
				FechaNacimiento:    fechaNacParsed,
				Sexo:               sexo,
			}

			// Leer contraseña
			constrasenaStr := password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(constrasenaStr), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal(err.Error())
			}
			clienteLocal.Contrasena = string(hashedPassword)
			clienteElephantSQL.Contrasena = string(hashedPassword)

			// Insertar cliente en la primera base de datos
			err = clienteRepoLocal.RegistrarCliente(clienteLocal)
			errorFatal(err)

			// Insertar cliente en la segunda base de datos
			err = clienteRepoElephantSQL.RegistrarCliente(clienteElephantSQL)
			errorFatal(err)

			fmt.Println("El usuario se creó correctamente")

			// Redireccionar o mostrar una respuesta al usuario
			http.Redirect(w, r, "/clientes/login", http.StatusSeeOther)

		} else {
			fmt.Println("No se pudo crear la cuenta debido a que no todos los datos ingresados fueron correctos")
			// Redireccionar al usuario a la página de registro con el error en la URL

			http.Redirect(w, r, "/clientes/registro?error=cedula", http.StatusSeeOther)
		}
	} else {
		if r.URL.Path == "/clientes/registro" && r.URL.Query().Get("error") == "true" {
			Registro()
		} else {
			Registro()
		}
	}
}

func getErrorMessage(errorType string) string {
	switch errorType {
	case "cedula":
		return "(error)La cédula no es válida o ya existe(error)"
	case "email":
		return "(error)El correo electrónico no es válido o ya existe(error)"
	case "celular":
		return "(error)El número de celular no es válido o ya existe(error)"
	default:
		return ""
	}
}

func validarCedulaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cedula := r.FormValue("cedula")

		// Realizar la validación de la cédula
		esValida := capalogica.ValidarCedula(cedula)

		// Crear una respuesta JSON con el resultado de la validación
		respuesta := struct {
			Error string `json:"error"`
		}{
			Error: "",
		}

		if !esValida {
			respuesta.Error = "La cédula no es válida"
		}

		// Convertir la respuesta a JSON
		respuestaJSON, err := json.Marshal(respuesta)
		if err != nil {
			// Manejar el error si no se pudo convertir a JSON
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Establecer el encabezado de respuesta para indicar que es JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribir la respuesta JSON
		_, err = w.Write(respuestaJSON)
		if err != nil {
			// Manejar el error si no se pudo escribir la respuesta
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Manejar el método de solicitud incorrecto si no es POST
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// //////////////////////////
// Genera un ID de sesión aleatorio
func generateSessionID() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Manejo de error
	}
	sessionID := hex.EncodeToString(randomBytes)
	return sessionID
}

func loadCredentials(credentialPath string) ([]byte, error) {
	creds, err := os.ReadFile(credentialPath)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo JSON de las credenciales: %v", err)
	}
	return creds, nil
}

func DetectIntentText(projectID, sessionID, text, languageCode string, creds []byte) (string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx, option.WithCredentialsJSON(creds))
	if err != nil {
		return "", fmt.Errorf("Error al crear el cliente de sesión: %v", err)
	}
	defer sessionClient.Close()

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", fmt.Errorf("Error al detectar la intención: %v", err)
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}

func ChangePassword(rw http.ResponseWriter, r *http.Request) {
	pass := r.FormValue("change")
	conn, err := connection.GetElephantSQLConnection()
	errorFatal(err)
	clienteRepo := repositories.NewClienteRepository(conn)
	if clienteRepo.ChangePass(pass) {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
	} else {
		fmt.Println("Error")
	}
}

func ResetearPassword(rw http.ResponseWriter, r *http.Request) {
	email := r.FormValue("reset")
	conn, err := connection.GetElephantSQLConnection()
	errorFatal(err)
	clienteRepo := repositories.NewClienteRepository(conn)
	fmt.Println(clienteRepo)
	if compareWithDatabaseEmail(email) {
		// Los datos coinciden, el usuario existe en la base de datos
		usuario := getUserDataFromDatabase(email)
		correo := usuario.Email
		telefono := clienteRepo.ObtenerTelefono(correo)
		fmt.Println(telefono)
	} else {
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
	}
}

func compareWithDatabaseEmail(email string) bool {
	conn, err := connection.GetElephantSQLConnection()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	// Crear instancia del repositorio de clientes
	clienteRepo := repositories.NewClienteRepository(conn)

	// Verificar si el cliente existe y las credenciales son válidas
	return clienteRepo.ConsultarClienteEmail(email)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Resto del código del manejador de inicio de sesión
	email := r.FormValue("email")
	password := r.FormValue("password")

	if capalogica.ValidarCorreoElectronico(email) &&
		capalogica.ValidarContrasena(password) {
		if compareWithDatabase(email, password) {
			// Los datos coinciden, el usuario existe en la base de datos
			usuario := getUserDataFromDatabase(email)
			fmt.Println(usuario)

			userDataBytes, err := json.Marshal(usuario)
			if err != nil {
				// Manejar el error si falla la codificación JSON
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Codificar los datos en base64
			userDataEncoded := base64.StdEncoding.EncodeToString(userDataBytes)

			cookie := http.Cookie{
				Name:    "session",
				Value:   userDataEncoded,                // Almacena los datos codificados en la cookie
				Expires: time.Now().Add(24 * time.Hour), // Establece la expiración de la cookie
				Path:    "/",
			}

			http.SetCookie(w, &cookie) // Establece la cookie en la respuesta

			// Construir la URL de redirección
			redireccionURL := "/chatbot/pharmacybot/conversacion/?id=59"

			// Redirigir al usuario a la nueva ruta
			http.Redirect(w, r, redireccionURL, http.StatusSeeOther)

			return
		} else {
			// Los datos no coinciden, el usuario no existe o las credenciales son incorrectas
			defer connection.CloseElephantSQLConnection() // Cierra la conexión con la base de datos

			http.Redirect(w, r, "/clientes/login", http.StatusFound)
		}
	} else {
		// Los datos no coinciden, el usuario no existe o las credenciales son incorrectas
		defer connection.CloseElephantSQLConnection() // Cierra la conexión con la base de datos

		http.Redirect(w, r, "/clientes/login", http.StatusFound)
	}
}

func checkSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	fmt.Println("cookie")
	if err == nil && cookie != nil {
		http.Redirect(w, r, "/chatbot/pharmacybot/conversacion/?id=59", http.StatusFound)
	}
	fmt.Println(cookie)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Establece el tiempo de vida de la cookie a un valor negativo para eliminarla
	}

	http.SetCookie(w, cookie)

	template, _ := template.ParseFiles("CAPA_DE_PRESENTACION/templates/clientes/login.html")

	template.Execute(w, nil)
}

func getUserDataFromDatabase(email string) UserData {
	conn, err := connection.GetElephantSQLConnection()
	if err != nil {
		log.Fatal(err.Error())
		return UserData{} // Devuelve una estructura vacía en caso de error
	}

	// Crear instancia del repositorio de clientes
	clienteRepo := repositories.NewClienteRepository(conn)

	// Obtener los datos del cliente
	cliente := clienteRepo.ConsultarCliente(email)
	fmt.Println(cliente)
	return UserData{
		IdUsuario: cliente.IdCliente,
		Email:     cliente.CorreoElectronico,
		Nombres:   cliente.Nombres,
		Cedula:    cliente.Cedula,
	}
}

func compareWithDatabase(email, password string) bool {
	conn, err := connection.GetElephantSQLConnection()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	// Crear instancia del repositorio de clientes
	clienteRepo := repositories.NewClienteRepository(conn)

	// Verificar si el cliente existe y las credenciales son válidas
	return clienteRepo.ConsultarClienteLogin(email, password)
}

func main() {

	http.Handle("/CAPA_DE_PRESENTACION/", http.StripPrefix("/CAPA_DE_PRESENTACION/", http.FileServer(http.Dir("CAPA_DE_PRESENTACION"))))
	http.HandleFunc("/", Index)
	http.HandleFunc("/validar-cedula", validarCedulaHandler)
	http.HandleFunc("/pregunta", EnviarPregunta)

	http.HandleFunc("/enviarDatosCliente", formularioHandler)
	http.HandleFunc("/actualizarConversacion/", actualizarConversacionHandler)
	http.HandleFunc("/eliminarConversacion/", eliminarconversacionHandler)
	sessionID = generateSessionID()
	http.HandleFunc("/registrarConversacion", func(rw http.ResponseWriter, r *http.Request) {
		err := registrarConversacionHandler(rw, r)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// Responde con éxito
		rw.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/loginCliente", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.HandleFunc("/clientes/registro", Registro())
	http.HandleFunc("/clientes/login", Login)

	http.HandleFunc("/clientes/cliente", MostrarCliente)

	http.HandleFunc("/chatbot/pharmacybot/conversacion/", MostrarChatBotConversacion)
	http.HandleFunc("/chatbot/pharmacybot", MostrarChatBot)

	http.HandleFunc("/resetPassword", ResetearPassword)
	http.HandleFunc("/changePassword", ChangePassword)

	http.ListenAndServe("localhost:3000", nil)

	fmt.Println("El servidor esta corriendo en el puerto localhost:3000")

}
