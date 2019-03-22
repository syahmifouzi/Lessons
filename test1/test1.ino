/*
  KhalifNovation
  School LineFollower

    Created : 12/02/2019 21:46:35
    Author :     Syed Uthman
  Hardware Target : KhalifTech Sch_LF (Arduino Nano Compatible Board)

  This project is dedicated to all primary and high school teachers around the world,
  especially from Malaysia. I hope this project will help student will get more understanding 
  about robotic.

  In this module we focus on making library to build coinfident in student's understanding in
  robotic field. By using human readable method, we hope student will be able understand. For example,
  usually to move forward we need to set some of digital pin need to set in a certain way. It is
  a little bit complicated for some biginners. So we intruduce robot.forward() as an alternatif.

  We hope you enjoy using this platform to help in teaching primary and high school students.

  Best Regards
  Syed Uthman
  3/3/2019

*/

#include "Motor.h"
#include "Sensor.h"
#include "Hardware.h"


Motor robot;
Sensor input;

// constant for serial 
#define debugMode false
#define serialMode true
#define baudrate 9600

void setup()
{
    init_Serial1();      //initialize serial
    init_Hardware();    //initialize robot hardware
    Serial.flush();
}

void loop()
{
  String LeftMotor = "";
  String RightMotor = "";
  char command = 'n';
  bool done = false;

  while (!done) {
    while(Serial.available()) {
      char character = Serial.read();
      if (command == 'l' && character != 'r') {
        LeftMotor.concat(character);
      } else if (command == 'r' && character != 'q') {
        RightMotor.concat(character);
      } else {
        command = character;
      }
  
      if (command == 'q') {
        done = true;
      }
    }
  }

   robot.motor_left(LeftMotor.toInt());
   robot.motor_right(RightMotor.toInt());

  // let the robot move for 1 sec before we analyse
    delay(500);

  robot.motor_left(0);
  robot.motor_right(0);

//  delay(500);

   int ls0 = input.LS_RAW(0);
   int ls1 = input.LS_RAW(1);
   int ls2 = input.LS_RAW(2);
   int ls3 = input.LS_RAW(3);
   int ls4 = input.LS_RAW(4);
   int ls5 = input.LS_RAW(5);
   int ls6 = input.LS_RAW(6);
  
    char buffer1[64];         //the ASCII of the integer will be stored in this char array
    memset(buffer1, 0, sizeof(buffer1));
    itoa(ls0, buffer1, 10); //(integer, yourBuffer, base)
    Serial.write('a');
    Serial.write(buffer1);
    memset(buffer1, 0, sizeof(buffer1));
    itoa(ls1, buffer1, 10);
    Serial.write('b');
    Serial.write(buffer1);
    memset(buffer1, 0, sizeof(buffer1));
    itoa(ls2, buffer1, 10);
    Serial.write('c');
    Serial.write(buffer1);
    memset(buffer1, 0, sizeof(buffer1));
    itoa(ls3, buffer1, 10);
    Serial.write('d');
    Serial.write(buffer1);
    memset(buffer1, 0, sizeof(buffer1));
    itoa(ls4, buffer1, 10);
    Serial.write('e');
    Serial.write(buffer1);
    memset(buffer1, 0, sizeof(buffer1));
    itoa(ls5, buffer1, 10);
    Serial.write('f');
    Serial.write(buffer1);
    memset(buffer1, 0, sizeof(buffer1));
    itoa(ls6, buffer1, 10);
    Serial.write('g');
    Serial.write(buffer1);
    Serial.write('q');

 
//  delay(1000);
//  robot.forward(0);
    

  //robot.forward(255);           //Move robot forward

  
//  Serial.print("BTN : ");         
//  Serial.println(input.isBTN_press());  //print condition of button 1 = press, 0 = release

  // Serial.print(input.LS_RAW(3));
}

// initialize Serial
void init_Serial1(){
    
    // check if user serial or debug mode
    if(debugMode || serialMode){

        // Additional serial setup can be add here
        Serial.begin(baudrate);

    }

}

// initialize Robot Hardware
void init_Hardware(){

  robot.begin();    //initialize output (Motor)
  input.begin();    //initialize input (Button,Jumper,IR Sensor)
  init_BUZZ();    //initialize buzzer   

}
