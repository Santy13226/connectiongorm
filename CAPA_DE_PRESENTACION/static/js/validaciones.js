console.log("validaciones2");
// Espera a que el documento esté listo

function validarCamposNumericos(value,tipo) {
  console.log(value)
  let regex = /^[0-9]+$/; // Expresión regular para solo permitir números
  let isValid = regex.test(value);

  if (!isValid) {
    if(tipo==="cedula")
    mensaje("La cédula solo debe tener dígitos numéricos válidos");
    if(tipo==="numero")
    mensaje("El número de celular solo debe tener dígitos numéricos");
  }else{
    if(document.querySelector(".contenedorInputCedula").querySelector(".errorMensajeDiv") && tipo==="cedula"){
      document.querySelector(".contenedorInputCedula").querySelector(".errorMensajeDiv").remove();
      return
    }
    if(document.querySelector(".contenedorInputCelular").querySelector(".errorMensajeDiv") && tipo==="numero"){
      document.querySelector(".contenedorInputCelular").querySelector(".errorMensajeDiv").remove();
      return
    }
  }
}

function validarCamposString(value, tipo) {
  console.log(value);
  let regex = /^[A-Za-z\s]+$/;
  let isValid = regex.test(value);

  if (!isValid) {
    if (tipo === "nombre")
      mensaje("El nombre debe contener solamente letras");
      if (tipo === "apellido")
      mensaje("El apellido debe contener solamente letras");
      
  } else {
  
    if(document.querySelector(".contenedorInputNombre").querySelector(".errorMensajeDiv") && tipo==="nombre"){
      document.querySelector(".contenedorInputNombre").querySelector(".errorMensajeDiv").remove();
      return
    }
    if(document.querySelector(".contenedorInputApellido").querySelector(".errorMensajeDiv") && tipo==="apellido"){
      document.querySelector(".contenedorInputApellido").querySelector(".errorMensajeDiv").remove();
      return
    }
  }
}



$(document).ready(function() {
  // Escucha el evento de envío del formulario
  $('#formulario').submit(function(event) {
    // Evita que el formulario se envíe normalmente
    event.preventDefault();

    // Obtén los datos del formulario
    var formData = $(this).serialize();

    // Realiza la solicitud AJAX
    $.ajax({
      type: 'POST',
      url: '/enviarDatosCliente',
      data: formData,
      success: function(response) {
        var regex = /\(error\)(.*?)\(error\)/g;
        var match = regex.exec(response);

        console.log(match);
        if (match && match.length > 1) {
          var errorMessage = match[1];

          if (errorMessage.includes("cédula")) {

            mensaje(errorMessage);
            document.getElementById("cedula").className = "errorValidacion";
            document.getElementById("cedula").focus();

          } else if (errorMessage.includes("celular")) {

            mensaje(errorMessage);
            document.getElementById("cedula").classList.remove = "errorValidacion";
            document.getElementById("cedula").className = "exitoValidacion";

            document.getElementById("celular").className = "errorValidacion";
            document.getElementById("celular").focus();
            
          }else if (errorMessage.includes("correo")) {

            mensaje(errorMessage);
            document.getElementById("cedula").classList.remove = "errorValidacion";
            document.getElementById("cedula").className = "exitoValidacion";
            
            document.getElementById("celular").classList.remove = "errorValidacion";
            document.getElementById("celular").className = "exitoValidacion";

            document.getElementById("email").className = "errorValidacion";
            document.getElementById("email").focus();
            
          } 

        } else {
          console.log("No se encontró un mensaje de error.");
          window.location.href = "http://localhost:3000/clientes/login";
        }

        if (response.error) {
          console.log(response.message); // Muestra el mensaje de error en la consola
          // Aquí puedes mostrar el mensaje de error en tu página o realizar otras acciones
        } else {
          // El registro fue exitoso, realiza las acciones necesarias
        }
      },
      error: function(xhr, status, error) {
        // Maneja errores de la solicitud AJAX
        console.error(error);
      }
    });
  });
});

const mensaje = (error) => {
  console.log("mensaje")
  let divMensajeError = document.createElement("div");
  let mensajeError = document.createElement("p");
  mensajeError.innerHTML = `${error}`;

  divMensajeError.style.zIndex = 100;
  divMensajeError.style.display = "flex";
  divMensajeError.style.justifyContent = "center";
  divMensajeError.style.alignItems = "center";
  divMensajeError.style.textAlign = "center";
  divMensajeError.style.fontWeight = "bold";
  divMensajeError.style.color = "white";
  divMensajeError.style.padding = "10px";
  divMensajeError.style.backgroundColor = "rgba(194, 61, 28, 0.781)";
  divMensajeError.style.marginTop = "2px";
  divMensajeError.style.marginBottom = "5px";
  divMensajeError.style.borderRadius = "10px";
  divMensajeError.appendChild(mensajeError);

  let contenedorInputCedula = document.querySelector(".contenedorInputCedula");
  let contenedorInputCelular = document.querySelector(".contenedorInputCelular");
  let contenedorInputEmail= document.querySelector(".contenedorInputEmail");
  let contenedorInputNombre= document.querySelector(".contenedorInputNombre");
  let contenedorInputApellido= document.querySelector(".contenedorInputApellido");
  divMensajeError.className="errorMensajeDiv"  


  if (error.includes("cédula")) { 
    if(!contenedorInputCedula.querySelector(".errorMensajeDiv")) 
    contenedorInputCedula.appendChild(divMensajeError.cloneNode(true));
  } 


  if (error.includes("celular")) {
    if(!contenedorInputCelular.querySelector(".errorMensajeDiv"))
    contenedorInputCelular.appendChild(divMensajeError.cloneNode(true));
  } 
  
  if (error.includes("email")) {
    if(!contenedorInputEmail.querySelector(".errorMensajeDiv"))
    contenedorInputEmail.appendChild(divMensajeError.cloneNode(true));
  } 

  if(error.includes("nombre")) {
    if(!contenedorInputNombre.querySelector(".errorMensajeDiv"))
    contenedorInputNombre.appendChild(divMensajeError.cloneNode(true));
  } 

  if(error.includes("apellido")) {
    if(!contenedorInputApellido.querySelector(".errorMensajeDiv"))
    contenedorInputApellido.appendChild(divMensajeError.cloneNode(true));
  }
}


