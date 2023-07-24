const { Vonage } = require('@vonage/server-sdk')


function generarPassword(longitud) {
  const caracteres = '!@#$%&*0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
  let password = '';
  for (let i = 0; i < longitud; i++) {
    const caracterAleatorio = caracteres.charAt(Math.floor(Math.random() * caracteres.length));
    password += caracterAleatorio;
  }
  return password;
}
const password = generarPassword(10);
console.log(password);

const vonage = new Vonage({
   apiKey: "16423620",
   apiSecret: "2JcBnpHs9lDPYPzn"
 })

 const from = "Vonage APIs"
 const to = "593984242647"
 const text = "Tu nueva contraseña en PharmacyBot es " + password;

 async function sendSMS() {
          await vonage.sms.send({to, from, text})
         .then(resp => { console.log('Message sent successfully'); console.log(resp); })
         .catch(err => { console.log('There was an error sending the messages.'); console.error(err); });
 }

 sendSMS();