/*
Esta utiliza las bibliotecas:
http://www.prometec.net/nrf2401/
https://www.arduino.cc/en/Tutorial/WebClientRepeating
https://github.com/DavyLandman/AESLib
 */
/**
 * Bibilioteca arduino
 */
#include <Arduino.h>
/*Fin*/

/**
 * Dependencias módulo de Ethernet
 */
#include <Ethernet.h>
#include <SPI.h>
/*Fin*/

/**
 * Dependencias módulo RF
 */
#include "RF24.h"
#include "nRF24L01.h"
#include <SPI.h>
/*Fin*/

/**
 * Dependencias AES
 */
#include <AESLib.h>
/*Fin*/

/**
 * Variables módulo de Ethernet
 */
// Enter a MAC address for your controller below.
// Newer Ethernet shields have a MAC address printed on a sticker on the shield
byte mac[] = {0xDE, 0xBD, 0xBE, 0xEF, 0xFE, 0xED};
// if you don't want to use DNS (and reduce your sketch size)
// use the numeric IP instead of the name for the server:
// IPAddress server(10, 145, 20, 62); // numeric IP for Google (no DNS)
IPAddress server(172, 16, 0, 1);
// char server[] = "192.168.1.238";    // name address for Google (using DNS)
// IP Servicio Web Destino de Datos

// Set the static IP address to use if the DHCP fails to assign
IPAddress ip(172, 16, 0, 2); // IP Arduino

// Initialize the Ethernet client library
// with the IP address and port of the server
// that you want to connect to (port 80 is default for HTTP):
EthernetClient client;
/*Fin*/

/**
 * Variables Módulo RF
 */
RF24 radio(9, 53);
const uint64_t pipes[2] = {0xF0F0F0F0E1LL, 0xF0F0F0F0D2LL};
/*Fin*/

/**
 * Variables AES
 */
uint8_t key[] = {111, 109, 97, 114, 108, 101, 111, 110, 97, 114, 100, 111, 122, 97, 109, 98};
/*End*/

void setup() {
  configSerial();
  configEthernet();
  // sendDataEthernet();
  configRF();
}

void configSerial() {
  // Open serial communications and wait for port to open:
  Serial.begin(9600);
  while (!Serial) {
    ; // wait for serial port to connect. Needed for Leonardo only
  }
}

void configEthernet() {
  // start the Ethernet connection:
  Serial.println("Try to config Ethernet using DHCP");
  //if (Ethernet.begin(mac) == 0) {//uncomment
    Serial.println("Failed to configure Ethernet using DHCP");
    // no point in carrying on, so do nothing forevermore:
    // try to congifure using IP address instead of DHCP:
    Ethernet.begin(mac, ip);
  //}
  // give the Ethernet shield a second to initialize:
  delay(100);
  Serial.println("Connecting...");
}

void configRF() {
  Serial.println("Config RF 2.4GHz");
  pinMode(53, OUTPUT);
  radio.begin();
  radio.setRetries(15, 15);
  // radio.setPayloadSize(8);
  radio.startListening();
  radio.openWritingPipe(pipes[1]);
  radio.openReadingPipe(1, pipes[0]);
}

// http://www.prometec.net/operaciones-bits/
uint8_t GetFirst4Bits(uint8_t b) { return b >> 4; }
uint8_t GetSecond4Bits(uint8_t b) {
  // b = b << 4;
  byte a = B00001111;
  return b & a;
}

void convert2Hex(char input[], byte tamanoInput, String &output) {
  String hexa[] = {"0", "1", "2", "3", "4", "5", "6", "7",
                   "8", "9", "A", "B", "C", "D", "E", "F"};
  for (uint8_t i = 0; i < tamanoInput; i++) {
    // Serial.println(i);
    /* code */
    uint8_t b = (uint8_t)input[i];
    // Serial.print("Bin: ");
    // Serial.println(b, BIN);
    // Serial.println(GetFirst4Bits(b), BIN);
    // Serial.println(GetSecond4Bits(b), BIN);
    String first = hexa[GetFirst4Bits(b)];
    String second = hexa[GetSecond4Bits(b)];
    // Serial.print(first);
    // Serial.println(second);
    output += first + second;
    // String stringOne = String(input[i], HEX);
    // Serial.println(stringOne);
  }
}

void sendDataEthernet(int valor, int valor2) {
  // close any connection before send a new request.
  // This will free the socket on the WiFi shield
  client.stop();
  if (client.connect(server, 9000)) {
    Serial.println("Connected");
    // // Make a HTTP request:
    // char data[30] = {0};
    // sprintf(data, "id=%i&st=%i", valor, valor2);
    // // sprintf(data, "id=%i&st=%i&t=%i", valor, valor2, millis());
    // Serial.print("data: ");
    // Serial.println(data);

    uint8_t iv[] = {97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112};

    // char data[] = "id=" + (char)valor + "&st=" + (char)valor + "&t=";
    //char data[] = "id=2&st=1&t=0000";
    char data[17] = {0};
    sprintf(data, "id=%i&st=%i&t=0000", valor, valor2);
    Serial.print("data: ");
    Serial.println(data);
    aes128_cbc_enc(key, iv, data, strlen(data));
    Serial.print("encrypted: ");
    Serial.println(data);
    // aes256_dec_single(key, data);
    // Serial.print("decrypted: ");
    // Serial.println(data);
    // encoding
    String dataHEX = "";
    byte tamano = sizeof(data) - 1;
    Serial.print("tamano: ");
    Serial.println(tamano);
    convert2Hex(data, tamano, dataHEX);
    Serial.print("hex: ");
    Serial.println(dataHEX);

    // client.println("GET /dev?data=hola HTTP/1.1");
    // Serial.println("GET /dev?data=hola HTTP/1.1");
    // client.println("GET /dev?data=" + String(encoded) + " HTTP/1.1");
    // Serial.println("GET /dev?data=" + String(encoded) + " HTTP/1.1");
    client.println("GET /dev?data=" + dataHEX + " HTTP/1.1");
    Serial.println("GET /dev?data=" + dataHEX + " HTTP/1.1");
    // client.println("Host: 10.145.20.62");
    // client.println("Host: 10.145.20.71");
    client.println("Host: 172.16.0.1:9000");
    client.println("Connection: close");
    client.println();
  } else {
    // kf you didn't get a connection to the server:
    Serial.println("Connection failed");
  }
  // interactEthernet();
  // readEthernet();
}

void readEthernet() {
  if (client.available()) {
    char c = client.read();
    Serial.print(c);
  }
}

void loop() {
  // interactEthernet();
  interactRF();
  // delay(1000); // Quitar, solo sirve para pruebas
  // sendDataEthernet(10, 1);
}

void interactEthernet() {
  // if there are incoming bytes available
  // from the server, read them and print them:
  if (client.available()) {
    char c = client.read();
    Serial.print(c);
  }

  // if the server's disconnected, stop the client:
  if (!client.connected()) {
    Serial.println();
    Serial.println("Disconnecting...");
    client.stop();
    // sendDataEthernet();
    // do nothing forevermore:
    // while(true);
  }
}

void interactRF() {
  // Importante!!! Borrar true
  if (radio.available()) { // Si hay datos disponibles
  //if (true || radio.available()) { // Solo para pruebas
    char got_isla[2];
    bool done = false;
    while (!done) {
      // Espera aqui hasta recibir algo

      done = radio.read(&got_isla, 2);
      Serial.print("Dato Recibido = ");
      Serial.print((int)got_isla[0]);
      Serial.print(",");
      Serial.println((int)got_isla[1]);
      delay(20); // Para dar tiempo al emisor
    }

    // radio.stopListening(); // Dejamos d escuchar para poder hablar
    //
    // radio.write(&got_isla, 2);
    // Serial.println("Enviando Respuesta");
    // radio.startListening(); // Volvemos a la escucha para recibir mas
    // paquetes
    sendDataEthernet(got_isla[0], got_isla[1]); // Se envía get
  }
}
