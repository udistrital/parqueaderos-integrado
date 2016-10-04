#include <SoftwareSerial.h>

//https://www.arduino.cc/en/Reference/SoftwareSerialConstructor
SoftwareSerial mySerial(2,3);//rxPin, TxPin

void setup()
{
  mySerial.begin(115200);                 
  Serial.begin(115200);                 
}

void loop()
{
  if (mySerial.available()){
    Serial.write(mySerial.read());
  }
  if (Serial.available()) {
    
    mySerial.write(Serial.read());
    Serial.write("hi");
  }

}
