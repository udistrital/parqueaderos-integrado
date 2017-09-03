
# Módulo WIFI ESP8266
--------------------
Este módulo se comunica por protocolo serial (UART http://www.mbedded.ninja/electronics/communication-protocols/uart-protocol) a una taza predeterminada de 115200 baudios (bits por segundo https://www.arduino.cc/en/Serial/Begin).

* Guía rápida de inicio con el módulo WiFi ESP8266, contiene recomendaciones, conexiones y ejemplo de comandos AT se usa con Arduino UNO pero no se recomienda alimentación desde esta, se da alternativamente instrucciones para FTDI 3.3V Board. http://rancidbacon.com/files/kiwicon8/ESP8266_WiFi_Module_Quick_Start_Guide_v_1.0.4.pdf
* Se da una presentación de conexión de pines completa para todo el módulo, se realiza también una forma de actualizar el firmware del dispositivo. http://www.instructables.com/id/The-First-Usage-of-ESP8266-With-Arduino-Uno/
* Se destaca por dar un ejemplo de comandos AT. http://rayshobby.net/first-impression-on-the-esp8266-serial-to-wifi-module/
* Es el datasheet del integrado base, se destaca la muestra de potencia consumida en modo Transmisor, Receptor y en estados de hibernación. http://www.electroschematics.com/wp-content/uploads/2015/02/esp8266-datasheet.pdf
* Muestra de otra forma la documentación habitual, tiene documentación de los comandos AT sin ejemplos. https://nurdspace.nl/ESP8266
* Conexión del módulo for dummies (torpes). http://www.teomaragakis.com/hardware/electronics/how-to-connect-an-esp8266-to-an-arduino-uno/
* Se muestra un mapa conceptual de los comandos AT con ejemplos y una lista extendida de ellos. Referencias de documentación, características técnicas y la forma de actualizar el firmware. https://www.itead.cc/wiki/ESP8266_Serial_WIFI_Module
* Contiene información necesaria para configurar el módulo como un Access Point. http://dominicm.com/esp8266-configure-access-point/
* La wiki definitiva para los comandos AT: https://github.com/espressif/ESP8266_AT/wiki

# Módulo Ultrasonido MB1010 LV-MaxSonar®-EZ1
-------------------------------------------
Este módulo se utiliza para detectar la presencia o ausencia de objetos, este se puede controlar por medio de 3 métodos, el análogo, el PWM y con comunicación serial. El más acosejable es PWM (http://playground.arduino.cc/Main/MaxSonar).

* Documentación oficial de implementación del sensor. http://www.maxbotix.com/articles/085.htm
* Características técnicas del dispositivo sumado a enlaces del datasheet, explicaciones, guía de selección del sensores de la misma familia. https://www.sparkfun.com/products/8502

# Módulo RF 433MHz Transmisor y Receptor
---------------------------------------
* Enlaces de un transmisor RF dirigido a muchos Receptores. https://www.youtube.com/watch?v=h3OevEdl674 https://github.com/Simsso/Arduino-Wireless-Module-Multiple-Receivers
* Ideas de cómo conectar 20 sensores wireless a un solo receptor. Un transmisor solo emite cuando es necesario y puede retransmitir algunas veces con base de tiempo aleatorio para asegurarse de la llegada del dato. Las colisiones son bajas si la concurrencia de eventos no es baja. Se puede también transmitir a diferentes bases de tiempo para reducir el número de colisiones. http://electronics.stackexchange.com/questions/74272/how-to-connect-20-wireless-sensors-in-one-receiver-with-arduino
* Uso de la biblioteca rc-switch que sirve enviar códigos RC, sirve para controlar dispositivos de radio remoto. En este se consigue la información de 3 sensores de temperatura remotos para hacer un promedio de temperatura, se utilizan arduinos nano v3 y un display I2C, se muestra el código. http://www.electronicsmayhem.com/?p=68 https://github.com/sui77/rc-switch
* RF Link Receiver - 4800bps (434MHz), 500ft (152.4m) rango (en las mejores condiciones), 4800bps data rate, 5V supply voltage. https://www.sparkfun.com/products/10532
* RF Link Transmitter - 434MHz, 500ft (152.4m) rango (en las mejores condiciones), 4800bps data rate, 5V supply voltage. https://www.sparkfun.com/products/10534
* Tutorial completo de módulos RF para Arduino, código, conexiones y vídeos. http://www.instructables.com/id/RF-315433-MHz-Transmitter-receiver-Module-and-Ardu/?ALLSTEPS
* Breve documentación de la biblioteca VirtualWire. http://www.pjrc.com/teensy/td_libs_VirtualWire.html
* Se destaca por tener una comparación de dispositivos de Radio Frecuencia "económicos". http://www.raviyp.com/embedded/152-low-cost-zigbee-modules-for-using-in-projects
* 11 protocolos de internet de las cosas (IoT) sobre los que usted necesita saber. https://www.rs-online.com/designspark/eleven-internet-of-things-iot-protocols-you-need-to-know-about
* Hardware abierto para internet de las cosas. http://www.postscapes.com/internet-of-things-hardware/

# NRF24L01 transceptor 2.4-2.5GHz en la banda ISM (Industrial, Científico, Médico)
Este módulo transmite y recibe datos. Tiene un bajo consumo de potencia y el rango de trabajo del tensión es de 1.9 hasta 3.6 voltios. Tiene multiples frecuencas (125 puntos).
* https://hallard.me/nrf24l01-real-life-range-test/
* https://github.com/maniacbug/RF24
* https://harizanov.com/2013/06/nrf24l01-range-testing/
* http://playground.arduino.cc/InterfacingWithHardware/Nrf24L01
* https://arduino-info.wikispaces.com/Nrf24L01-Info
* http://www.nordicsemi.com/eng/Products/2.4GHz-RF/nRF24L01
* https://arduino-info.wikispaces.com/Nrf24L01-2.4GHz-HowTo?responseToken=fcd06dad835d084c7d89d1d4a147b989
* http://howtomechatronics.com/tutorials/arduino/arduino-wireless-communication-nrf24l01-tutorial/
* http://projectsfromtech.blogspot.com.co/2013/05/nrf24l01-arduino-communication-on.html
* http://howtomechatronics.com/tutorials/arduino/arduino-wireless-communication-nrf24l01-tutorial/

