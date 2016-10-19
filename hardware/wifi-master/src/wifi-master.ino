#include <SoftwareSerial.h>

// based on: http://dominicm.com/esp8266-configure-access-point/
// https://www.arduino.cc/en/Reference/SoftwareSerialConstructor
SoftwareSerial esp8266(2, 3); // rxPin, TxPin of Arduino rx-tx tx-rx

static char response[255] = "\0"; // static vs const?

void setup() {
  configSerial();
  configWIFI();
  hacerRedWifi();
}

void configSerial() { Serial.begin(9600); }

void configWIFI() {
  esp8266.begin(115200);
  esp8266.println(F("AT")); // Print both? NL & CR
  showResponse();
  delay(5000);
  esp8266.println(F("AT+RST")); // for know state
  showResponse();
  esp8266.println(F("AT+CIOBAUD=115200"));
  showResponse();
  // esp8266.println(F("AT+CWMODE?"));
  // showResponse();
  esp8266.println(F("AT+CWMODE=3")); //	1= Sta, 2= AP, 3=both, Sta is the
                                     // default mode of router, AP is a normal
                                     // mode for devices
  showResponse();

  Serial.println("Configuración Wifi Terminada");
}

void hacerRedWifi() {
  esp8266.println(
      F("AT+CWSAP=\"ESP8266\",\"123\",11,0")); // Sin seguridad, no funcionaron
  // ssid: string, ESP8266 softAP’ SSID
  // pwd: string, MAX: 64 bytes
  // chl: channel id
  // ecn:
  // 0 OPEN
  // 2 WPA_PSK
  // 3 WPA2_PSK
  // 4 WPA_WPA2_PSK
  showResponse();
  // AT+CWSAP?
  // AT+CIPAP?
  esp8266.println(F("AT+CIPAP=\"192.168.4.1\""));
  showResponse();

  esp8266.println(F("AT+CIPAPMAC=\"1a:fe:34:a4:35:42\""));
  showResponse();

  esp8266.println(F("AT+CWDHCP=0,1"));
  // mode
  // 0 : set ESP8266 softAP
  // 1 : set ESP8266 station
  // 2 : set both softAP and station
  // en
  // 0 : Disable DHCP
  // 1 : Enable DHCP
  showResponse();
}

void loop() { redirSerialAT(); }

/**
 * Solo se muestra la respuesta,
 * la taza de muestreo debería ser la misma para que no ocurran problemas
 */
void showResponse() {
  if (esp8266.available()) {
    Serial.write(esp8266.read());
  }
  Serial.println("");
  delay(500);
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
