package capalogica

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ////////////////////////////////////////////////VALIDACIONES///////////////////////////////////////////
func ValidarNombreApellido(nombre string) bool {
	match, _ := regexp.MatchString("^[A-Za-z\\s]+$", nombre)

	return match
}

func ValidarDireccionDomicilio(direccion string) bool {
	// Puedes implementar tus propias validaciones para la dirección de domicilio
	// según los criterios de Ecuador
	return len(direccion) > 0
}

func ValidarNumeroCelular(numero string) bool {
	// Verificar que el número de celular tenga 10 dígitos y comience con "09"
	match, _ := regexp.MatchString("^09\\d{8}$", numero)

	return match
}

func ValidarCorreoElectronico(correo string) bool {

	// Puedes implementar tus propias validaciones para el correo electrónico
	// según los criterios de Ecuador
	// Aquí se utiliza una validación básica de formato de correo electrónico
	match, _ := regexp.MatchString("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$", correo)

	return match
}

func ValidarContrasena(contrasena string) bool {

	return len(contrasena) >= 8
}

func ValidarEdad(fechaNac time.Time) bool {
	// Obtener la fecha actual
	fechaActual := time.Now()

	// Calcular la edad en años
	edad := fechaActual.Year() - fechaNac.Year()

	// Verificar si aún no se ha celebrado el cumpleaños este año
	if fechaActual.Month() < fechaNac.Month() || (fechaActual.Month() == fechaNac.Month() && fechaActual.Day() < fechaNac.Day()) {
		edad--
	}

	// Verificar que la edad sea mayor o igual a 18
	return edad >= 18
}
func ValidarCedula(cedula string) bool {
	// Eliminar espacios en blanco y guiones
	cedula = strings.ReplaceAll(cedula, " ", "")
	cedula = strings.ReplaceAll(cedula, "-", "")

	// Verificar longitud de la cédula
	if len(cedula) != 10 {
		return false
	}

	// Verificar que todos los caracteres sean dígitos
	for _, char := range cedula {
		if char < '0' || char > '9' {
			return false
		}
	}

	// Obtener provincia de la cédula (primeros dos dígitos)
	provincia, err := strconv.Atoi(cedula[:2])
	if err != nil {
		return false
	}

	// Verificar provincia válida (de 1 a 24)
	if provincia < 1 || provincia > 24 {
		return false
	}

	// Verificar el tercer dígito (debe ser 6, 9 o un número entre 0 y 5)
	tipo := cedula[2]
	if tipo != '6' && tipo != '9' && (tipo < '0' || tipo > '5') {
		return false
	}

	// Verificar dígitos de validación
	suma := 0
	coeficientes := []int{2, 1, 2, 1, 2, 1, 2, 1, 2}
	for i := 0; i < 9; i++ {
		digito, err := strconv.Atoi(string(cedula[i]))
		if err != nil {
			return false
		}
		producto := digito * coeficientes[i]
		if producto >= 10 {
			producto -= 9
		}
		suma += producto
	}

	// Verificar el último dígito de validación
	ultimoDigito, err := strconv.Atoi(string(cedula[9]))
	if err != nil {
		return false
	}
	digitoVerificador := (suma + ultimoDigito) % 10
	if digitoVerificador != 0 {
		return false
	}

	return true
}

func ValidarSexo(sexo string) bool {

	// Verificar que el sexo sea uno de los valores permitidos (por ejemplo, "M" o "F")
	sexo = strings.ToUpper(sexo)
	if sexo != "M" && sexo != "F" {
		return false
	}
	// Verificar que no haya caracteres especiales en el sexo
	match, _ := regexp.MatchString("^[A-Za-z]+$", sexo)
	return match
}

// Validar el código del producto
func ValidarItemCodigo(codigo string) bool {

	// Verificar que el código tenga exactamente 6 dígitos
	match, _ := regexp.MatchString("^\\d{6}$", codigo)
	return match
}

// Validar el nombre del producto
func ValidarNombreProducto(nombre string) bool {

	// Verificar que el nombre no esté vacío y no contenga caracteres especiales
	match, _ := regexp.MatchString("^[A-Za-z\\s]+$", nombre)
	return match
}

// Validar la descripción del producto
func ValidarDescripcion(descripcion string) bool {

	// Verificar que la descripción no esté vacía y no contenga caracteres especiales
	match, _ := regexp.MatchString("^[a-zA-Z0-9 ]+$", descripcion)
	return match
}

// Validar el stock del producto
func ValidarStock(stockStr string) bool {

	stock, err := strconv.Atoi(stockStr) // Convertir el valor string a int
	if err != nil {
		return false // Si la conversión falla, retorna falso
	}

	// Verificar que el stock sea mayor o igual a cero
	return stock >= 0
}

// Validar el PVP (Precio de Venta al Público) del producto
func ValidarPVP(pvpStr string) bool {

	pvp, err := strconv.ParseFloat(pvpStr, 64) // Convertir el valor string a float64
	if err != nil {
		return false // Si la conversión falla, retorna falso
	}

	// Verificar que el PVP sea mayor a cero
	return pvp > 0
}
