function enviarPregunta() {
  var preguntaInput = document.getElementById("preguntaInput");
  var pregunta = preguntaInput.value.trim();

  if (pregunta === "") {
    return;
  }

  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/pregunta", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
      var conversaciones = JSON.parse(xhr.responseText);
      actualizarConversaciones(conversaciones);

      // Desplazar hacia abajo
      var conversacionesList = document.getElementById("conversacionesList");
      conversacionesList.scrollTop = conversacionesList.scrollHeight;
    }
  };

  var data = "pregunta=" + encodeURIComponent(pregunta);
  xhr.send(data);

  preguntaInput.value = "";
}

function agregarConversacion(pregunta, respuesta) {

  var conversacionesList = document.getElementById("conversacionesList");

  var li = document.createElement("li");
  var contenedorPregunta = document.createElement("div");
  var preguntaLabel = document.createElement("label");
  var contenedorRespuesta = document.createElement("div");
  var respuestaLabel = document.createElement("label");

  li.className = "conversacion";
  contenedorPregunta.className = "contenedorPregunta";
  preguntaLabel.className = "pregunta burbuja";
  contenedorRespuesta.className = "contenedorRespuesta";
  respuestaLabel.className = "respuesta burbuja";

  preguntaLabel.innerHTML = pregunta;
  respuestaLabel.innerHTML = respuesta;

  contenedorPregunta.appendChild(preguntaLabel);
  contenedorRespuesta.appendChild(respuestaLabel);
  li.appendChild(contenedorPregunta);
  li.appendChild(contenedorRespuesta);
  conversacionesList.appendChild(li);
}

$(document).ready(function() {
  // Manejar el evento click del botón
  $('#btnNewChat').click(function() {
      // Realizar la llamada AJAX a la función registrarConversacionHandler
      $.ajax({
          type: 'POST', // Método HTTP POST
          url: '/registrarConversacion', // Reemplaza con la URL correcta hacia tu función registrarConversacionHandler
          success: function(data) {
              // Manejar la respuesta exitosa de la llamada AJAX, si es necesario
              console.log('Conversación registrada exitosamente');
          },
          error: function(xhr, status, error) {
              // Manejar errores de la llamada AJAX, si es necesario
              console.log('Error al registrar la conversación:', error);
          }
      });
  });
});


document.getElementById("enviarBtn").addEventListener("click", enviarPregunta);

