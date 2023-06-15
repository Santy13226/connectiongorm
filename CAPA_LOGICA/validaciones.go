package capalogica

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////VALIDACIONES///////////////////////////////////////////
func ValidarNombreApellido(nombre string) bool {
	fmt.Println("Ingreso a validar nombre")
	match, _ := regexp.MatchString("^[A-Za-z\\s]+$", nombre)
	return match
}

func ValidarDireccionDomicilio(direccion string) bool {
	fmt.Println("Validando dirección de domicilio:", direccion)
	// Puedes implementar tus propias validaciones para la dirección de domicilio
	// según los criterios de Ecuador
	return len(direccion) > 0
}

func ValidarNumeroCelular(numero string) bool {
	fmt.Println("Validando número de celular:", numero)
	// Verificar que el número de celular tenga 10 dígitos y comience con "09"
	match, _ := regexp.MatchString("^09\\d{8}$", numero)
	return match
}

func ValidarCorreoElectronico(correo string) bool {
	fmt.Println("Validando correo electrónico:", correo)
	// Puedes implementar tus propias validaciones para el correo electrónico
	// según los criterios de Ecuador
	// Aquí se utiliza una validación básica de formato de correo electrónico
	match, _ := regexp.MatchString("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$", correo)
	return match
}

func ValidarContrasena(contrasena string) bool {
	fmt.Println("Validando contraseña:", contrasena)
	return len(contrasena) >= 8
}

func ValidarEdad(edadStr string) bool {
	fmt.Println("Validando edad:", edadStr)
	edad, err := strconv.Atoi(edadStr)
	if err != nil {
		return false
	}
	// Verificar que la edad esté en un rango válido (por ejemplo, entre 18 y 100)
	if edad < 18 || edad > 100 {
		return false
	}
	// Verificar que no haya caracteres especiales en la edad
	match, _ := regexp.MatchString("^[0-9]+$", edadStr)
	return match
}

func ValidarCedula(cedula string) bool {
	// Eliminar espacios en blanco y guiones
	cedula = strings.ReplaceAll(cedula, " ", "")
	cedula = strings.ReplaceAll(cedula, "-", "")

	// Verificar longitud de la cédula
	if len(cedula) != 10 {
		return false
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

	// Verificar dígitos de validación
	suma := 0
	for i := 0; i < 9; i++ {
		digito, err := strconv.Atoi(string(cedula[i]))
		if err != nil {
			return false
		}

		if i%2 == 0 {
			digito *= 2
			if digito > 9 {
				digito -= 9
			}
		}

		suma += digito
	}

	verificador, err := strconv.Atoi(string(cedula[9]))
	if err != nil {
		return false
	}

	suma += verificador

	// Verificar que la suma sea múltiplo de 10
	return suma%10 == 0
}

func ValidarSexo(sexo string) bool {
	fmt.Println("Validando sexo:", sexo)
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
	fmt.Println("Validando código de producto:", codigo)
	// Verificar que el código tenga exactamente 6 dígitos
	match, _ := regexp.MatchString("^\\d{6}$", codigo)
	return match
}

// Validar el nombre del producto
func ValidarNombreProducto(nombre string) bool {
	fmt.Println("Validando nombre de producto:", nombre)
	// Verificar que el nombre no esté vacío y no contenga caracteres especiales
	match, _ := regexp.MatchString("^[a-zA-Z0-9 ]+$", nombre)
	return match
}

// Validar la descripción del producto
func ValidarDescripcion(descripcion string) bool {
	fmt.Println("Validando descripción de producto:", descripcion)
	// Verificar que la descripción no esté vacía y no contenga caracteres especiales
	match, _ := regexp.MatchString("^[a-zA-Z0-9 ]+$", descripcion)
	return match
}

// Validar el stock del producto
func ValidarStock(stockStr string) bool {
	fmt.Println("Validando stock de producto:", stockStr)
	stock, err := strconv.Atoi(stockStr) // Convertir el valor string a int
	if err != nil {
		return false // Si la conversión falla, retorna falso
	}

	// Verificar que el stock sea mayor o igual a cero
	return stock >= 0
}

// Validar el PVP (Precio de Venta al Público) del producto
func ValidarPVP(pvpStr string) bool {
	fmt.Println("Validando PVP de producto:", pvpStr)
	pvp, err := strconv.ParseFloat(pvpStr, 64) // Convertir el valor string a float64
	if err != nil {
		return false // Si la conversión falla, retorna falso
	}

	// Verificar que el PVP sea mayor a cero
	return pvp > 0
}
