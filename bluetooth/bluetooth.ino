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
}

void loop()
{

//  if(Serial.available() == 2){
//        char command = Serial.read();
//        int value = Serial.read();
//        if(command == 'm') { 
//            robot.forward(value);
//        }
//        if(command == 'l') {
//            // contro lights
//        }
//        Serial.write(command);
//        Serial.write(value);
//    }
//  delay(1000);
//  robot.forward(0);
    

  //robot.forward(255);           //Move robot forward

  runCalibration();
  
//  Serial.print("BTN : ");         
//  Serial.println(input.isBTN_press());  //print condition of button 1 = press, 0 = release
//  
//  Serial.print("JMP : ");
//  Serial.println(input.isJMP_connected()); //print condition of jumper 1 = close, 0 = open

  Serial.println("Threshold 0");
  Serial.print(input.getIntData(THRESSHOLD, 0));  
  Serial.print("\t"); 
  Serial.print(input.getIntData(HIGH_VAL, 0));
  Serial.print("\t");
  Serial.println(input.getIntData(LOW_VAL, 0));

  Serial.println("Threshold 1");
  Serial.print(input.getIntData(THRESSHOLD, 1));  
  Serial.print("\t"); 
  Serial.print(input.getIntData(HIGH_VAL, 1));
  Serial.print("\t");
  Serial.println(input.getIntData(LOW_VAL, 1));

  Serial.println("Threshold 2");
  Serial.print(input.getIntData(THRESSHOLD, 2));  
  Serial.print("\t"); 
  Serial.print(input.getIntData(HIGH_VAL, 2));
  Serial.print("\t");
  Serial.println(input.getIntData(LOW_VAL, 2));

  Serial.println("Threshold 3");
  Serial.print(input.getIntData(THRESSHOLD, 3));  
  Serial.print("\t"); 
  Serial.print(input.getIntData(HIGH_VAL, 3));
  Serial.print("\t");
  Serial.println(input.getIntData(LOW_VAL, 3));

  Serial.println("Threshold 4");
  Serial.print(input.getIntData(THRESSHOLD, 4));  
  Serial.print("\t"); 
  Serial.print(input.getIntData(HIGH_VAL, 4));
  Serial.print("\t");
  Serial.println(input.getIntData(LOW_VAL, 4));

  Serial.println("Threshold 5");
  Serial.print(input.getIntData(THRESSHOLD, 5));  
  Serial.print("\t"); 
  Serial.print(input.getIntData(HIGH_VAL, 5));
  Serial.print("\t");
  Serial.println(input.getIntData(LOW_VAL, 5));

  Serial.println("Threshold 6");
  Serial.print(input.getIntData(THRESSHOLD, 6));  
  Serial.print("\t"); 
  Serial.print(input.getIntData(HIGH_VAL, 6));
  Serial.print("\t");
  Serial.println(input.getIntData(LOW_VAL, 6));
  
  
  Serial.println("Sensor");
  Serial.print(input.LS_RAW(0));      
  Serial.print("\t");
  Serial.print(input.LS_RAW(1));
  Serial.print("\t");
  Serial.print(input.LS_RAW(2));
  Serial.print("\t");
  Serial.print(input.LS_RAW(3));
  Serial.print("\t");
  Serial.print(input.LS_RAW(4));
  Serial.print("\t");
  Serial.print(input.LS_RAW(5));
  Serial.print("\t");
  Serial.println(input.LS_RAW(6));    //print IR analog sensor value (IR0 - IR6)
  
//  BUZZ_ON();                //Turn on buzzer
//  delay(1000);
//  BUZZ_OFF();               //Turn off buzzer
  delay(1000);

}

void runCalibration() {
  while (!input.isBTN_press());
  input.calibration(60);
  BUZZ_ON();                //Turn on buzzer
  delay(1000);
  BUZZ_OFF(); 
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
