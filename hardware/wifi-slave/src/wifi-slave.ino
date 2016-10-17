#include <SoftwareSerial.h>

// static or const is in read-only parts of the memory?
// https://www.arduino.cc/en/Reference/SoftwareSerialConstructor
SoftwareSerial esp8266(2, 3); // rxPin, TxPin of Arduino rx-tx tx-rx

static char response[255] = "\0"; // static vs const?

void setup() {
  Serial.begin(9600);
  configWIFI();
}

void loop() { redirSerialAT(); }

void configWIFI() {
  esp8266.begin(115200);
  esp8266.println(F("AT"));
  showResponse();
  esp8266.println(F("AT+RST")); // for know state
  showResponse();
  delay(5000);
}

void showResponse() {
  bufferSerial(response);
  Serial.println(response);
}

/**
 * Se utiliza para redirigir los comandos AT del serial del arduino
 * al puerto del módulo WIFI ESP8266
 */
void redirSerialAT() {
  if (esp8266.available()) {
    Serial.write(esp8266.read());
  }
  if (Serial.available()) {
    esp8266.write(Serial.read());
  }
}

/**
 * Realiza la lectura del buffer del serial
 * caracteres hasta que la transmisión serial termina
 * Acumulador de caracteres recursivo
 */
void bufferSerial(char *cad) {
  int i = 0;
  char c = '\0';
  strcpy(cad, "");
  while (esp8266.available()) {
    c = esp8266.read();
    cad[i] = c;
    i++;
  }
}
