// Watch video here: https://www.youtube.com/watch?v=U0hJBc1Ppv4

/*
Accelerometer connection pins (I2C) to Arduino are shown below:

Arduino     Accelerometer ADXL345
  A5            SCL
  A4            SDA
  3.3V          CS
  3.3V          VCC
  GND           GND
  
Arduino     Servo No. 1
  5V          5V  (Red Wire)
  GND         GND (Black Wire)
  D5          (last wire for control - might be white color)
  
Arduino     Servo No. 2
  5V          5V  (Red Wire)
  GND         GND (Black Wire)
  D6          (last wire for control - might be white color)
*/

#include <Wire.h>
#include <ADXL345.h>

#include <Servo.h>

Servo servo1;  // create servo object to control a servo
Servo servo2;

ADXL345 adxl; //variable adxl is an instance of the ADXL345 library

int x, y, z;
int rawX, rawY, rawZ;
int mappedRawX, mappedRawY;

void setup() {
  Serial.begin(9600);
  adxl.powerOn();
  servo1.attach(5);
  servo2.attach(6);
}

void loop() {
  adxl.readAccel(&x, &y, &z); //read the accelerometer values and store them in variables  x,y,z
  Serial.print(" x = ");Serial.print(x);
  Serial.print(" y = ");Serial.print(y);
  Serial.print(" z = ");Serial.println(z);
  rawX = x - 7;
  rawY = y - 6;
  rawZ = z + 10;
  
  if (rawX < -255) rawX = -255; else if (rawX > 255) rawX = 255;
  if (rawY < -255) rawY = -255; else if (rawY > 255) rawY = 255;
   
  mappedRawX = map(rawX, -255, 255, 0, 180);
  mappedRawY = map(rawY, -255, 255, 0, 180);
  
  servo1.write(mappedRawX);
  delay(15);
  servo2.write(180 - mappedRawY);
  delay(15);
  
  //Serial.print(" mappedRawX = "); Serial.print(mappedRawX); // raw data with offset
  //Serial.print(" mappedRawY = "); Serial.println(mappedRawY); // raw data with offset
}
