
#Módulo WIFI ESP8266
--------------------
Este módulo se comunica por protocolo serial (UART http://www.mbedded.ninja/electronics/communication-protocols/uart-protocol) a una taza predeterminada de 115200 baudios (bits por segundo https://www.arduino.cc/en/Serial/Begin). 

* Guía rápida de inicio con el módulo WiFi ESP8266, contiene recomendaciones, conexiones y ejemplo de comandos AT se usa con Arduino UNO pero no se recomienda alimentación desde esta, se da alternativamente instrucciones para FTDI 3.3V Board. http://rancidbacon.com/files/kiwicon8/ESP8266_WiFi_Module_Quick_Start_Guide_v_1.0.4.pdf
* Se da una presentación de conexión de pines completa para todo el módulo, se realiza también una forma de actualizar el firmware del dispositivo. http://www.instructables.com/id/The-First-Usage-of-ESP8266-With-Arduino-Uno/
* Se destaca por dar un ejemplo de comandos AT. http://rayshobby.net/first-impression-on-the-esp8266-serial-to-wifi-module/
* Es el datasheet del integrado base, se destaca la muestra de potencia consumida en modo Transmisor, Receptor y en estados de hibernación. http://www.electroschematics.com/wp-content/uploads/2015/02/esp8266-datasheet.pdf
* Muestra de otra forma la documentación habitual, tiene documentación de los comandos AT sin ejemplos. https://nurdspace.nl/ESP8266
* Conexión del módulo for dummies (torpes). http://www.teomaragakis.com/hardware/electronics/how-to-connect-an-esp8266-to-an-arduino-uno/

#Módulo Ultrasonido MB1010 LV-MaxSonar®-EZ1
Este módulo se utiliza para detectar la presencia o ausencia de objetos, este se puede controlar por medio de 3 métodos, el análogo, el PWM y con comunicación serial. El más acosejable es PWM (http://playground.arduino.cc/Main/MaxSonar).

* Documentación oficial de implementación del sensor. http://www.maxbotix.com/articles/085.htm
* Características técnicas del dispositivo sumado a enlaces del datasheet, explicaciones, guía de selección del sensores de la misma familia. https://www.sparkfun.com/products/8502
