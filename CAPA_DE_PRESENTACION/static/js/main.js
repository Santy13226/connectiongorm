const btnUser = document.querySelector('.button-user');
const menuUser = document.querySelector('#inactive');
const btnNewChat = document.querySelector('.button-newchat');
const conversationContainer = document.querySelector('.container-conversation');
const btnSettings = document.querySelector('.ajustes');
const containerOverlay = document.getElementById('inactivo');
let isOverlayVisible = false;
const btnCerrar = document.querySelector('.cerrar');

btnUser.addEventListener('click', toggleMenuUser);
btnNewChat.addEventListener('click', createNewChat);

function toggleMenuUser(){
    menuUser.style.display = (menuUser.style.display === 'none') ? 'block' : 'none';
}
function createNewChat(){
    conversationList.push({
        name: 'Chat1'
    });
}

btnSettings.addEventListener('click', toggleContainerOverlay);

function toggleContainerOverlay() {
  isOverlayVisible = !isOverlayVisible;
  containerOverlay.style.display = isOverlayVisible ? 'block' : 'none';
}

btnCerrar.addEventListener('click', toggleCerrar);

function toggleCerrar(){
    isOverlayVisible = !isOverlayVisible;
    containerOverlay.style.display = isOverlayVisible ? 'block' : 'none';
}


// Codido para crear el Objeto de las Conversaciones y mostrarlas
const conversationList = [];
conversationList.push({
    name: 'Chat1'
});
conversationList.push({
    name: 'Medicamentos disponibles'
});
conversationList.push({
    name: 'Medicamentos disponibles'
});

function renderConversation(arr){
    for(chat of arr){
        const conversation = document.createElement('button');
        conversation.classList.add('conversation');
        const span = document.createElement('span');
        span.classList.add('icon-message');
        const text = document.createTextNode(chat.name);
        conversation.appendChild(span);
        conversation.appendChild(text);
        conversationContainer.appendChild(conversation);     
    }
}

renderConversation(conversationList);

// Llamada a la API para obtener las conversaciones históricas
// function getChatbotHistory() {
//     var projectId = '';
//     var sessionId = '1'; // Puede ser un ID único para identificar cada sesión
  
//     var url = `https://dialogflow.googleapis.com/v3beta1/projects/newagent-lpsi/locations/global/agents/-/sessions/123456789/entityTypes:batchUpdate`;
  
//     // Realiza la llamada a la API usando fetch o tu librería HTTP preferida
//     fetch(url, {
//       method: 'GET',
//       headers: {
//         'Authorization': 'Bearer YOUR_AUTH_TOKEN'
//       }
//     })
//     .then(response => response.json())
//     .then(data => {
//       // Procesa los datos de las conversaciones históricas y muéstralos en tu página web
//       console.log(data);
//     })
//     .catch(error => {
//       console.error('Error al obtener las conversaciones históricas:', error);
//     });
//   }





// document.addEventListener("DOMContentLoaded", function() {
//     var chatDisplay = document.getElementById("chat-display");
//     var userInput = document.getElementById("user-input");
//     var sendButton = document.getElementById("send-button");

//     sendButton.addEventListener("click", function() {
//         var userMessage = userInput.value;
//         displayUserMessage(userMessage);
//         sendUserMessage(userMessage);
//         userInput.value = "";
//     });

//     function displayUserMessage(message) {
//         var userMessageElement = document.createElement("div");
//         userMessageElement.className = "user-message";
//         userMessageElement.innerHTML = "<strong>Tú:</strong> " + message;
//         chatDisplay.appendChild(userMessageElement);
//     }

//     function displayAssistantMessage(message) {
//         var assistantMessageElement = document.createElement("div");
//         assistantMessageElement.className = "assistant-message";
//         assistantMessageElement.innerHTML = "<strong>Asistente:</strong> " + message;
//         chatDisplay.appendChild(assistantMessageElement);
//     }

//     function sendUserMessage(message) {
//         // Envía el mensaje del usuario a Dialogflow o cualquier otro servicio de backend
//         // y obtén la respuesta del asistente.
//         // Aquí puedes agregar la lógica de comunicación con Dialogflow mediante API.
//         // Por simplicidad, en este ejemplo simplemente mostraremos una respuesta predefinida.
//         var assistantResponse = "¡Hola! Soy el asistente. ¿En qué puedo ayudarte?";
//         displayAssistantMessage(assistantResponse);
//     }
    
//     // Agregar el iframe al DOM
//     var iframe = document.createElement("iframe");
//     iframe.setAttribute("allow", "microphone;");
//     iframe.setAttribute("width", "350");
//     iframe.setAttribute("height", "430");
//     iframe.setAttribute("src", "https://console.dialogflow.com/api-client/demo/embedded/8dc5a833-3bf9-4cfe-b2de-607a8147e9f1");
//     chatDisplay.appendChild(iframe);
// });

