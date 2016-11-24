/*  ----------------------------------------------------------------
    http://www.prometec.net/duplex-nrf2401
    Prog_79B_Emisor

    Usando un NRF2401 para comunicar dos Arduinos en modo Duplex
    Programa Emisor:
   --------------------------------------------------------------------
 */

/**
  * Dependencias módulo RF
  */
#include "RF24.h"
#include "nRF24L01.h"
#include <SPI.h>
/*Fin*/


/**
 * Variables Módulo RF
 */
RF24 radio(9, 10);
const uint64_t pipes[2] = {0xF0F0F0F0E1LL,
                           0xF0F0F0F0D2LL}; // LongLong = 64 bits.
/*Fin*/

/**
 * Variables Módulo Ultrasonido
 */
#define trigPin 6
#define echoPin 5
#define ledEstado 8 // Indicador para deteccion del vehículo

int limite = 212; // Distancia max entre el suelo y el vehiculo (cms)
// 212 cms Toyota Hilux
long duracion, distancia;
volatile byte estadoIsla = 0;
/*Fin*/


void setup(void) {
  configSerial();
  configRF();
  configUltrasonido();
}

void configSerial() { Serial.begin(9600); }

void configRF() {
  pinMode(10, OUTPUT);
  radio.begin();

  radio.setRetries(15, 15); // Maximos reintentos
  // radio.setPayloadSize(8);   // Reduce el payload de 32 si tienes problemas

  // Open pipes to other nodes for communication
  radio.openWritingPipe(pipes[0]);
  radio.openReadingPipe(1, pipes[1]);
}

void configUltrasonido() {
  pinMode(trigPin, OUTPUT);
  pinMode(echoPin, INPUT);
  pinMode(ledEstado, OUTPUT);
}

void loop(void) {
  enviarDatoRF();
  medirDistancia();
  delay(1000);
}

void enviarDatoRF() {
  radio.stopListening(); // Paramos la escucha para poder hablar
  char idIsla = 1;
  //char estado = 0;
  char msg[2] = {idIsla, estadoIsla};
  // unsigned long time = millis();
  Serial.print("Enviando  ");
  Serial.print((int)msg[0]);
  Serial.print(",");
  Serial.println((int)msg[1]);
  // bool ok = radio.write(&time, sizeof(unsigned long));
  bool ok = radio.write(msg, 2);

  if (ok) {
    Serial.println("ok...");
  } else {
    Serial.println("failedEstado");
  }

  radio.startListening(); // Volvemos a la escucha
  //
  // unsigned long started_waiting_at = millis();
  // bool timeout = false;
  // while (!radio.available() && !timeout) { // Esperasmos repsuesta hasta
  // 200ms
  //   if (millis() - started_waiting_at > 200) {
  //     timeout = true;
  //   }
  // }
  //
  // if (timeout) {
  //   Serial.println("FailedEstado, response timed out");
  // } else { // Leemos el mensaje recibido
  //   char got_isla[2];
  //   radio.read(&got_isla, 2);
  //
  //   Serial.print("Respuesta = ");
  //   Serial.print((int)got_isla[0]);
  //   Serial.println((int)got_isla[1]);
  // }
}

void medirDistancia() {
  digitalWrite(trigPin,
               LOW); // Nos aseguramos de que el trigger está desactivado
  // delayMicroseconds(2);              // Para asegurarnos de que el trigger
  // esta LOW
  digitalWrite(trigPin, HIGH); // Activamos el pulso de salida
  // delayMicroseconds(10);             // Esperamos 10µs. El pulso sigue active
  // este tiempo
  digitalWrite(trigPin, LOW); // Cortamos el pulso y a esperar el echo
  duracion = pulseIn(echoPin, HIGH);
  distancia = duracion / 2 / 29.1;
  if (distancia < limite) {
    Serial.println(String(distancia) + " cm.");
    digitalWrite(ledEstado, HIGH);
    estadoIsla = 1;
  } else {
    Serial.println(String(distancia) + " cm.");
    digitalWrite(ledEstado, LOW);
    estadoIsla = 0;
  }
  // delay (500) ;                  // Para limitar el número de mediciones
}
