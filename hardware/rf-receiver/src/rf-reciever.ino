// Arduino Version 1.6.9
// platformio lib install VirtualWire
#include <VirtualWire.h>
byte message[VW_MAX_MESSAGE_LEN]; // a buffer to store the incoming messages
byte messageLength = VW_MAX_MESSAGE_LEN; // the size of the message
String readString = String(30);

void setup() {
  Serial.begin(9600);
  Serial.println("El dispositivo esta listo");
  // Initialize the IO and ISR
  vw_setup(2000); // Bits per sec
  vw_rx_start();  // Start the receiver
}

void loop() {
  if (vw_get_message(message, &messageLength)) { // Non-blocking
    Serial.print("Received: ");
    for (int i = 0; i < messageLength; i++) {
      Serial.write(message[i]);
      readString += (char)message[i];
    }
    Serial.println();
    // Serial.println(readString);

    // Inicio
    String param1 = "param1";
    readValueHTTP(param1);
    Serial.println(param1);
    // Fin

    // Inicio
    String param2 = "param2";
    readValueHTTP(param2);
    Serial.println(param2);
    // Fin
  }
  readString = "";
}

// RFC 7231 - Hypertext Transfer Protocol (HTTP/1.1)
void readValueHTTP(String &keyParam) {
  String parametro = keyParam + "=";
  int indice = readString.indexOf(parametro);
  if (indice > -1) {
    keyParam = readString.substring(indice + parametro.length());
    int indice2 = keyParam.indexOf("&");
    if (indice2 > -1) {
      indice = indice2;
    }
    keyParam = keyParam.substring(0, indice);
  }
}
