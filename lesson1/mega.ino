
// Basic serial communication with ESP8266
// Uses serial monitor for communication with ESP8266
//
//  Pins
//  Arduino pin 2 (RX) to ESP8266 TX
//  Arduino pin 3 to voltage divider then to ESP8266 RX
//  Connect GND from the Arduiono to GND on the ESP8266
//  Pull ESP8266 CH_PD HIGH
//
// When a command is entered in to the serial monitor on the computer
// the Arduino will relay it to the ESP8266
//

// #include <SoftwareSerial.h>
// SoftwareSerial ESPserial(10, 11); // RX | TX
// Need to apply external pull-up resistor to use other pins as serial

String a;
char buf[80];

int readline(int readch, char *buffer, int len)
{
    static int pos = 0;
    int rpos;

    if (readch > 0)
    {
        switch (readch)
        {
        case '\r': // Ignore CR
            break;
        case '\n': // Return on new-line
            rpos = pos;
            pos = 0; // Reset position index ready for next time
            return rpos;
        default:
            if (pos < len - 1)
            {
                buffer[pos++] = readch;
                buffer[pos] = 0;
            }
        }
    }
    return 0;
}

void setup()
{
    Serial.begin(115200); // communication with the host computer
    // while (!Serial)   { ; }

    // Start the software serial for communication with the ESP8266
    // ESPserial.begin(115200);
    // pinMode(10, INPUT_PULLUP);

    Serial1.begin(115200);

    Serial.println("");
    Serial.println("Remember to to set Both NL & CR in the serial monitor.");
    Serial.println("Ready");
    Serial.println("");
}

void loop()
{
    // listen for communication from the ESP8266 and then write it to the serial monitor
    // while( ESPserial.available() )   {
    // a = ESPserial.readString();
    // Serial.println(a);  }

    if (readline(Serial1.read(), buf, 80) > 0)
    {
        Serial.print("You entered: >");
        Serial.print(buf);
        Serial.println("<");
        // Serial1.println("TESTING!!");
        // Serial1 doesnot print to serial monitor
        // only serial can print to serial monitor
    }

    // listen for user input and send it to the ESP8266
    //    if ( Serial.available() )       {  ESPserial.write( Serial.read() );  }
}