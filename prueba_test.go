package main

import (
	"net/http"
	"net/http/httptest"

	//"net/url"
	"strings"
	"testing"
)

/*func TestLoginHandler_ValidCredentials(t *testing.T) {
	t.Run("ValidCredentials", func(t *testing.T) {
		// Crea una solicitud HTTP POST con las credenciales válidas
		form := url.Values{}
		form.Add("email", "santiguaylla@gmail.com")
		form.Add("password", "12345678")

		// Crea una solicitud HTTP con el formulario
		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/login", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}

		// Establece las cabeceras necesarias para la solicitud POST
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Crea un ResponseRecorder para grabar la respuesta
		rr := httptest.NewRecorder()

		// Llamar a la función loginHandler con la solicitud y el ResponseRecorder
		loginHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/chatbot/pharmacybot/conversacion/?id=59"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}
	})

	t.Run("InvalidCredentials", func(t *testing.T) {
		form := url.Values{}
		form.Add("email", "usuario@correo.com")
		form.Add("password", "contraseñaincorrecta")
		req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %v", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Crear un ResponseRecorder para grabar la respuesta
		rr := httptest.NewRecorder()

		// Llamar a la función loginHandler con la solicitud falsa y el ResponseRecorder
		loginHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 302 Found para redireccionar)
		if rr.Code != http.StatusFound {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusFound, rr.Code)
		}

		// Verificar la URL de redirección
		expectedURL := "/clientes/login"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "El usuario no se encuentra en los registros."
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})
	t.Run("InvalidEmailFormat", func(t *testing.T) {
		// Crea una solicitud HTTP POST con un correo electrónico inválido (sin @)
		form := url.Values{}
		form.Add("email", "usuariocorreo.com")
		form.Add("password", "contraseña123")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/login", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		loginHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 400 Bad Request para indicar un error en el formato)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusBadRequest, rr.Code)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "El formato del correo electrónico no es válido."
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})

	t.Run("InvalidPasswordFormat", func(t *testing.T) {
		// Crea una solicitud HTTP POST con una contraseña inválida (menos de 8 caracteres)
		form := url.Values{}
		form.Add("email", "usuario@example.com")
		form.Add("password", "abc123")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/login", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		loginHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 400 Bad Request para indicar un error en el formato)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusBadRequest, rr.Code)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "La contraseña debe tener al menos 8 caracteres."
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})
}

func TestFormularioHandler(t *testing.T) {
	t.Run("ValidFormData", func(t *testing.T) {
		// Crear una solicitud HTTP POST con datos válidos del formulario
		// Puedes modificar los valores aquí según tus necesidades de prueba
		form := url.Values{}
		form.Add("cedula", "1724008063")
		form.Add("nombres", "Andrea")
		form.Add("apellidos", "Melena")
		form.Add("direccion", "Riobamba")
		form.Add("celular", "0987654999")
		form.Add("email", "andreao@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "F")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Crear un ResponseRecorder para grabar la respuesta
		rr := httptest.NewRecorder()

		// Llamar a la función formularioHandler con la solicitud y el ResponseRecorder
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/login"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}
	})

	t.Run("InvalidEmailFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con un correo electrónico inválido (sin @)
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "Dirección")
		form.Add("celular", "0987654321")
		form.Add("email", "usuariocorreo.com") // Correo electrónico inválido
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Crear un ResponseRecorder para grabar la respuesta
		rr := httptest.NewRecorder()

		// Llamar a la función formularioHandler con la solicitud y el ResponseRecorder
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "El formato del correo electrónico no es válido."
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})

	t.Run("InvalidCedulaFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con una cédula inválida (menos de 10 dígitos)
		form := url.Values{}
		form.Add("cedula", "123456789") // Cédula inválida (menos de 10 dígitos)
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "Dirección")
		form.Add("celular", "0987654321")
		form.Add("email", "usuario@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "La cédula no es válida"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})

	t.Run("InvalidNombreFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con un nombre inválido (sin caracteres alfabéticos)
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "123456") // Nombre inválido (sin caracteres alfabéticos)
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "Dirección")
		form.Add("celular", "0987654321")
		form.Add("email", "usuario@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "El nombre no es válido"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})
	t.Run("InvalidApellidoFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con un apellido inválido (sin caracteres alfabéticos)
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "123456") // Apellido inválido (sin caracteres alfabéticos)
		form.Add("direccion", "Dirección")
		form.Add("celular", "0987654321")
		form.Add("email", "usuario@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "El apellido no es válido"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})

	t.Run("InvalidDireccionFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con una dirección inválida (sin caracteres alfabéticos)
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "") // Dirección inválida (vacio)
		form.Add("celular", "0987654321")
		form.Add("email", "usuario@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "La dirección no es válida porque está vacia"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})
	t.Run("InvalidCelularFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con un número de celular inválido (menos de 10 dígitos)
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "Dirección")
		form.Add("celular", "1234567") // Número de celular inválido (menos de 10 dígitos)
		form.Add("email", "usuario@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "El número de celular no es válido"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})

	t.Run("InvalidFechaNacFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con una fecha de nacimiento inválida (formato incorrecto)
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "Dirección")
		form.Add("celular", "0987654321")
		form.Add("email", "usuario@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2022-01-01") // Fecha de nacimiento inválida (menor de 18 años)
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "La fecha de nacimiento no es válida, el cliente debe ser mayor de edad"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})
	t.Run("InvalidPasswordFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con una contraseña inválida (menos de 8 caracteres)
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "Dirección")
		form.Add("celular", "0987654321")
		form.Add("email", "usuario@example.com")
		form.Add("password", "abc123") // Contraseña inválida (menos de 8 caracteres)
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "M")

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "La contraseña debe tener al menos 8 caracteres"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})
	t.Run("InvalidSexFormat", func(t *testing.T) {
		// Crear una solicitud HTTP POST con un valor inválido para el campo "sexo"
		form := url.Values{}
		form.Add("cedula", "1234567890")
		form.Add("nombres", "Nombre")
		form.Add("apellidos", "Apellido")
		form.Add("direccion", "Dirección")
		form.Add("celular", "0987654321")
		form.Add("email", "usuario@example.com")
		form.Add("password", "contraseña123")
		form.Add("fechaNac", "2000-01-01")
		form.Add("sexo", "Otro") // Valor inválido para el campo "sexo"

		req, err := http.NewRequest("POST", "http://localhost:3000/clientes/registro", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatalf("Error al crear la solicitud HTTP: %s", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		formularioHandler(rr, req)

		// Verificar el código de estado de la respuesta (debe ser 303 See Other para redireccionar)
		if rr.Code != http.StatusSeeOther {
			t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusSeeOther, rr.Code)
		}

		// Verificar que la redirección sea a la URL esperada
		expectedURL := "/clientes/registro?error=cedula"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("URL de redirección incorrecta. Esperada: %s, Obtener: %s", expectedURL, location)
		}

		// Verificar el mensaje de error en el cuerpo de la respuesta
		expectedErrorMessage := "El valor del campo sexo es inválido"
		if !strings.Contains(rr.Body.String(), expectedErrorMessage) {
			t.Errorf("Mensaje de error esperado no encontrado en la respuesta. Esperado: %s", expectedErrorMessage)
		}
	})
}
*/

func TestEnviarPregunta(t *testing.T) {
	// Crea una solicitud HTTP POST con la pregunta del usuario
	form := strings.NewReader("pregunta=")
	req, err := http.NewRequest("POST", "http://localhost:3000/enviar-pregunta", form)
	if err != nil {
		t.Fatalf("Error al crear la solicitud HTTP: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Crea un ResponseRecorder para grabar la respuesta
	rr := httptest.NewRecorder()

	// Llamar a la función EnviarPregunta con la solicitud y el ResponseRecorder
	EnviarPregunta(rr, req)

	// Verificar el código de estado de la respuesta (debe ser 200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("El código de estado no es el esperado. Esperado: %d, Obtenido: %d", http.StatusOK, rr.Code)
	}

	// Verificar que se haya recibido alguna respuesta
	if rr.Body.Len() == 0 {
		t.Errorf("No se recibió ninguna respuesta del chatbot")
	}
}
