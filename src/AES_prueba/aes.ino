 /*
 * Esta utiliza la biblioteca https://github.com/DavyLandman/AESLib
 */

#include <Arduino.h>
#include <AESLib.h>

const int BlockSize = 16;

void setup()
{
        //  Usage example for AES256:
        Serial.begin(57600);
        uint8_t key[] = {111, 109, 97, 114, 108, 101, 111, 110, 97, 114, 100, 111, 122, 97, 109, 98};
        char data[] = "HelloWorldHolaMu";
        //Serial.println(strlen(data));
        uint8_t iv[] = {97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112};
        
        aes128_cbc_enc(key, iv, data, strlen(data));
        //Serial.println(sizeof(data));
        Serial.print("encrypted:");
        Serial.println(data);
         for (int i=0; i<sizeof(data);i++)
        {
          //int x =int(data[i]);
          uint8_t x = uint8_t(data[i]);
          Serial.print(x);
          Serial.print(" ");
        }
        Serial.println("");  
        
        aes128_cbc_dec(key, iv, data, strlen(data));
        Serial.print("decrypted:");
        Serial.println(data);
        Serial.print("iv:");
        for (int i=0; i<sizeof(iv);i++)
        {
          Serial.print(char(iv[i]));
        }
        Serial.println("");             
}

void loop()
{
  
}

