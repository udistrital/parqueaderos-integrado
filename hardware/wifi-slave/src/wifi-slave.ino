#include <SoftwareSerial.h>

// static or const is in read-only parts of the memory?
// https://www.arduino.cc/en/Reference/SoftwareSerialConstructor
SoftwareSerial esp8266(2, 3); // rxPin, TxPin of Arduino rx-tx tx-rx

static char response[255] = "\0"; // static vs const?

void setup() {
  configSerial();
  configWIFI();
}

void configSerial() { Serial.begin(115200); }

void configWIFI() {
  esp8266.begin(115200);
  esp8266.println(F("AT")); // Print both? NL & CR
  saveAndShowResponse();
  esp8266.println(F("AT+CWMODE=1"));
  saveAndShowResponse();
  esp8266.println(F("AT+RST")); // for know state
  saveAndShowResponse();
  delay(5000);
  enviarDato();
}

void enviarDato() {
  esp8266.println(F("AT+CWLAP"));
  saveAndShowResponse();
  esp8266.println(F("AT+CWJAP=\"Familia\",\"useche;)\""));
  saveAndShowResponse();
  esp8266.println(F("AT+CIFSR"));
  saveAndShowResponse();
  esp8266.println(F("AT+CIPMUX=0"));
  saveAndShowResponse();
  esp8266.println(F("AT+CIPSTART=\"TCP\",\"192.168.0.11\",8080"));
  saveAndShowResponse();
  String msj =
      "GET / HTTP/1.1\r\nHost: 192.168.0.11\r\nConnection: close\r\n\r\n";
  Serial.println(msj.length());
  esp8266.println(F("AT+CIPSEND=57"));
  saveAndShowResponse();
  esp8266.print(msj);
  saveAndShowResponse();
  Serial.println("Terminé?");
}

void loop() {
  enviarDato2();
  redirSerialAT();
}

void enviarDato2() {
  esp8266.println(F("AT+CIPSTART=\"TCP\",\"192.168.0.11\",8080"));
  saveAndShowResponse();
  String msj =
      "GET / HTTP/1.1\r\nHost: 192.168.0.11\r\nConnection: close\r\n\r\n";
  Serial.println(msj.length());
  esp8266.println(F("AT+CIPSEND=57"));
  saveAndShowResponse();
  esp8266.print(msj);
  saveAndShowResponse();
}

/**
 * Solo se muestra la respuesta,
 * la taza de muestreo debería ser la misma para que no ocurran problemas
 */
void showResponse() {
  if (esp8266.available()) {
    Serial.write(esp8266.read());
  }
}

/**
 * Se guarda la respuesta en Response y se imprime en serial
 */
void saveAndShowResponse() {
  saveResponse();
  Serial.println(response);
}

/**
 * Se guarda la respuesta en Response
 */
void saveResponse() { bufferSerial(response); }

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
