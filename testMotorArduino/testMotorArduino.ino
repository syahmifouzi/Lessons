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
  robot.motor_left(50);
  robot.motor_right(50);
  delay(500);
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
