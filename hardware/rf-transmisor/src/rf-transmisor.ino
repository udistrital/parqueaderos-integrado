// Se instaló con:
// platformio lib search VirtualWire
// platformio lib install VirtualWire
#include <VirtualWire.h>
void setup() {
  // Initialize the IO and ISR
  vw_setup(2000); // Bits per sec
}
void loop() {
  // Se debe encender Dispositivo
  send((char *)"param1=value1&param2=value2");
  // Se debe apagar dispositivo (Pin Salida Digital de Control?)
  delay(1000);
}
void send(char *message) {
  vw_send((uint8_t *)message, strlen(message));
  vw_wait_tx(); // Wait until the whole message is gone
}
