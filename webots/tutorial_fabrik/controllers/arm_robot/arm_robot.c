/*
 * File:          arm_robot.c
 * Date:
 * Description:
 * Author:
 * Modifications:
 */

/*
 * You may need to add include files like <webots/distance_sensor.h> or
 * <webots/motor.h>, etc.
 */
#include <webots/robot.h>
#include <stdio.h>
#include <stdlib.h>
#include <webots/motor.h>
#include <math.h>
#include <webots/supervisor.h>

/*
 * You may want to add macros here.
 */
#define TIME_STEP 64

void find_point_coord_at_yz(
  double* ref_point,
  double arm_length,
  double theta,
  double offset_theta,
  double* target_point
  ) {
  theta += offset_theta;
  target_point[0] = ref_point[0];
  target_point[1] = ref_point[1] + (arm_length * sin(theta));
  target_point[2] = ref_point[2] + (arm_length * cos(theta));
}

void find_point_coord_at_xz(
  double* ref_point,
  double arm_length,
  double theta,
  double offset_theta,
  double* target_point
  ) {
  theta += offset_theta;
  target_point[0] = ref_point[0];
  target_point[1] = ref_point[1] + (arm_length * sin(theta));
  target_point[2] = ref_point[2] + (arm_length * cos(theta));
}

void IK_3D_formula() {
  double A[3] = {0,0,0};
  double B[3] = {0,0,0};
  double arm_length_or_link = 0;

  double unit_vector_AB_magnitude = sqrt(pow(A[0]-B[0],2) + pow(A[1]-B[1],2) + pow(A[2]-B[2],2));    
  double unit_vector_AB_i = (A[0]-B[0]) / unit_vector_AB_magnitude;
  unit_vector_AB_i *= -1; // invert the direction
  double unit_vector_AB_j = (A[1]-B[1]) / unit_vector_AB_magnitude;
  unit_vector_AB_j *= -1; // invert the direction
  double unit_vector_AB_k = (A[2]-B[2]) / unit_vector_AB_magnitude;
  unit_vector_AB_k *= -1; // invert the direction
  
  double A_prime_i = B[0] - (arm_length_or_link * unit_vector_AB_i);
  double A_prime_j = B[1] - (arm_length_or_link * unit_vector_AB_j);
  double A_prime_k = B[2] - (arm_length_or_link * unit_vector_AB_k);
}

void IK_at_xyz(
  double* A,
  double* B,
  double arm_length,
  double* unit_vector_AB,
  double* A_prime
  ) {

  double unit_vector_AB_magnitude = sqrt(pow(A[0]-B[0],2) + pow(A[1]-B[1],2) + pow(A[2]-B[2],2));
  unit_vector_AB[0] = (A[0]-B[0]) / unit_vector_AB_magnitude;
  unit_vector_AB[0] *= -1; // invert the direction
  unit_vector_AB[1] = (A[1]-B[1]) / unit_vector_AB_magnitude;
  unit_vector_AB[1] *= -1; // invert the direction
  unit_vector_AB[2] = (A[2]-B[2]) / unit_vector_AB_magnitude;
  unit_vector_AB[2] *= -1; // invert the direction
  
  A_prime[0] = B[0] - (arm_length * unit_vector_AB[0]);
  A_prime[1] = B[1] - (arm_length * unit_vector_AB[1]);
  A_prime[2] = B[2] - (arm_length * unit_vector_AB[2]);
}

void IK_at_yz(
  double* A,
  double* B,
  double arm_length,
  double* unit_vector_AB,
  double* A_prime
  ) {

  double unit_vector_AB_magnitude = sqrt(pow(A[0]-B[0],2) + pow(A[1]-B[1],2) + pow(A[2]-B[2],2));
  unit_vector_AB[1] = (A[1]-B[1]) / unit_vector_AB_magnitude;
  unit_vector_AB[1] *= -1; // invert the direction
  unit_vector_AB[2] = (A[2]-B[2]) / unit_vector_AB_magnitude;
  unit_vector_AB[2] *= -1; // invert the direction
  
  A_prime[0] = A[0];
  A_prime[1] = B[1] - (arm_length * unit_vector_AB[1]);
  A_prime[2] = B[2] - (arm_length * unit_vector_AB[2]);
}

void IK_at_xz(
  double* A,
  double* B,
  double arm_length,
  double* unit_vector_AB,
  double* A_prime
  ) {

  double unit_vector_AB_magnitude = sqrt(pow(A[0]-B[0],2) + pow(A[1]-B[1],2) + pow(A[2]-B[2],2));
  unit_vector_AB[0] = (A[0]-B[0]) / unit_vector_AB_magnitude;
  unit_vector_AB[0] *= -1; // invert the direction
  unit_vector_AB[2] = (A[2]-B[2]) / unit_vector_AB_magnitude;
  unit_vector_AB[2] *= -1; // invert the direction
  
  A_prime[0] = B[0] - (arm_length * unit_vector_AB[0]);
  A_prime[1] = A[1];
  A_prime[2] = B[2] - (arm_length * unit_vector_AB[2]);
}

void FK_at_xyz(
  double* A,
  double arm_length,
  double* unit_vector_AB,
  double* B_primeprime
  ) {
  
  B_primeprime[0] = A[0] + (arm_length * unit_vector_AB[0]);
  B_primeprime[1] = A[1] + (arm_length * unit_vector_AB[1]);
  B_primeprime[2] = A[2] + (arm_length * unit_vector_AB[2]);
}

void FK_at_yz(
  double* A,
  double arm_length,
  double* unit_vector_AB,
  double* B_primeprime
  ) {
  
  B_primeprime[0] = A[0];
  B_primeprime[1] = A[1] + (arm_length * unit_vector_AB[1]);
  B_primeprime[2] = A[2] + (arm_length * unit_vector_AB[2]);
}

void FK_at_xz(
  double* A,
  double arm_length,
  double* unit_vector_AB,
  double* B_primeprime
  ) {
  
  B_primeprime[0] = A[0] + (arm_length * unit_vector_AB[0]);
  B_primeprime[1] = A[1];
  B_primeprime[2] = A[2] + (arm_length * unit_vector_AB[2]);
}


/*
 * This is the main program.
 * The arguments of the main function can be specified by the
 * "controllerArgs" field of the Robot node
 */
int main(int argc, char **argv) {
  /* necessary to initialize webots stuff */
  wb_robot_init();
  printf("Hello World!\n");
  
  const double offset_arm_1 = 0;
  // const double offset_arm_1 = 0.22348;
  double base_offset_x = 100;
  double arm1_offset_y = -75;
  double base_height = 50;
  double p0[3] = {base_offset_x, 0.0, base_height};
  double d1 = 155;
  double p1[3] = {base_offset_x, arm1_offset_y, base_height + d1};
  double d2 = 180;
  double p2[3] = {base_offset_x, arm1_offset_y, base_height + d1 + d2};
  
  double virtual_d1 = sqrt(pow(p1[1] - p0[1], 2) + pow(p1[2] - p0[2], 2));
  double angle_offset_p1 = atan2(p1[1] - p0[1], p1[2] - p0[2]);
  
  WbNodeRef target_box_node = wb_supervisor_node_get_from_def("target_box");
  if (target_box_node == NULL) {
    fprintf(stderr, "No DEF target_box node found in the current world file\n");
    exit(1);
  }
  WbFieldRef tb_trans_field = wb_supervisor_node_get_field(target_box_node, "translation");

  /*
   * You should declare here WbDeviceTag variables for storing
   * robot devices like this:
   *  WbDeviceTag my_sensor = wb_robot_get_device("my_sensor");
   *  WbDeviceTag my_actuator = wb_robot_get_device("my_actuator");
   */
   WbDeviceTag arm_1 = wb_robot_get_device("arm_rm_1");
   WbDeviceTag arm_2 = wb_robot_get_device("arm_rm_2");
   
   // double newpos = M_PI; // rotate PI rad == 180 degree
   // newpos = M_PI / 2;

  /* main loop
   * Perform simulation steps of TIME_STEP milliseconds
   * and leave the loop when the simulation is over
   */
  while (wb_robot_step(TIME_STEP) != -1) {
    
    /*
     * Read the sensors :
     * Enter here functions to read sensor data, like:
     *  double val = wb_distance_sensor_get_value(my_sensor);
     */
     
    const double *target_in_m = wb_supervisor_field_get_sf_vec3f(tb_trans_field);
    // printf("target_box is at position: %g %g %g\n", values[0], values[1], values[2]);

    const double target[3] = {target_in_m[0]*1000,target_in_m[1]*1000,target_in_m[2]*1000};
    // printf("target_box is at position: %g %g %g\n", target[0], target[1], target[2]);
    const double trans_y = target[1]; // convert to mm
    
    /* INVERSE KINEMATICS START*/
    
    double p1_prime[3] = {0,0,0};
    double unit_vector_1T[3] = {0,0,0};
    IK_at_xyz(p1,target,d2,unit_vector_1T,p1_prime);

    // double unit_vector_1T_magnitude = sqrt(pow(p1[0]-target[0],2) + pow(p1[1]-target[1],2) + pow(p1[2]-target[2],2));    
    // double unit_vector_1T_i = (p1[0]-target[0]) / unit_vector_1T_magnitude;
    // unit_vector_1T_i *= -1; // invert the direction
    // double unit_vector_1T_j = (p1[1]-target[1]) / unit_vector_1T_magnitude;
    // unit_vector_1T_j *= -1; // invert the direction
    // double unit_vector_1T_k = (p1[2]-target[2]) / unit_vector_1T_magnitude;
    // unit_vector_1T_k *= -1; // invert the direction
    
    // double p1_prime_i = target[0] - (d2*unit_vector_1T_i);
    // double p1_prime_j = target[1] - (d2*unit_vector_1T_j);
    // double p1_prime_k = target[2] - (d2*unit_vector_1T_k);
    
    
    double p0_prime[3] = {0,0,0};
    double unit_vector_01prime[3] = {0,0,0};
    IK_at_xyz(p0,p1_prime,virtual_d1,unit_vector_01prime,p0_prime);
    
    
    // double p0_prime[3] = {0,0,0};
    // double unit_vector_01prime[3] = {0,0,0};
    // IK_at_yz(p0,target,virtual_d1,unit_vector_01prime,p0_prime);
    
    
    // double unit_vector_01prime_magnitude = sqrt(pow(p0[0]-p1_prime_i,2) + pow(p0[1]-p1_prime_j,2) + pow(p0[2]-p1_prime_k,2));
    // double unit_vector_01prime_i = (p0[0]-p1_prime_i) / unit_vector_01prime_magnitude;
    // unit_vector_01prime_i *= -1; // invert the direction
    // double unit_vector_01prime_j = (p0[1]-p1_prime_j) / unit_vector_01prime_magnitude;
    // unit_vector_01prime_j *= -1; // invert the direction
    // double unit_vector_01prime_k = (p0[2]-p1_prime_k) / unit_vector_01prime_magnitude;
    // unit_vector_01prime_k *= -1; // invert the direction
    
    
    // printf("unit_vector_1T_i is: %g    ", unit_vector_1T_i);
    // printf("unit_vector_1T_j is: %g    ", unit_vector_1T_j);
    // printf("unit_vector_1T_k is: %g\n", unit_vector_1T_k);
    
    // printf("p1_prime_i is: %g    ", p1_prime_i);
    // printf("p1_prime_j is: %g    ", p1_prime_j);
    // printf("p1_prime_k is: %g\n", p1_prime_k);
    
    
    /* INVERSE KINEMATICS END*/
    
    /* FORWARD KINEMATICS START*/
    
    double p1_primeprime[3] = {0,0,0};
    FK_at_xyz(p0,virtual_d1,unit_vector_01prime,p1_primeprime);
    
    // double p1_primeprime_i = p0[0] + (virtual_d1 * unit_vector_01prime_i);
    // double p1_primeprime_j = p0[1] + (virtual_d1 * unit_vector_01prime_j);
    // double p1_primeprime_k = p0[2] + (virtual_d1 * unit_vector_01prime_k);
    
    double p2_primeprime[3] = {0,0,0};
    FK_at_xyz(p1_primeprime,d2,unit_vector_1T,p2_primeprime);
    
    // double p2_primeprime_i = p1_primeprime_i + (d2 * unit_vector_1T_i);
    // double p2_primeprime_j = p1_primeprime_j + (d2 * unit_vector_1T_j);
    // double p2_primeprime_k = p1_primeprime_k + (d2 * unit_vector_1T_k);
    
    /* FORWARD KINEMATICS END*/
    
    /* ADD OFFSET START*/
    
    // p1_primeprime_i += p1[0];
    // p1_primeprime_j += p1[1];
    // p1_primeprime_k += p1[2];
    
    // p2_primeprime_i += p2[0];
    // p2_primeprime_j += p2[1];
    // p2_primeprime_k += p2[2];
    
    /* ADD OFFSET END*/
    
    /* FIND ROTATION ANGLE START*/
    
    double theta_0_rad = atan2(p1_primeprime[1] - p0[1], p1_primeprime[2] - p0[2]);
    theta_0_rad *= -1;
    double theta_1_rad = atan2(p2_primeprime[0] - p1_primeprime[0], p2_primeprime[2] - p1_primeprime[2]);

    // printf("theta_0_rad is: %g\n", theta_0_rad);
    // printf("angle_offset_p1 is: %g\n", angle_offset_p1);
    
    /* FIND ROTATION ANGLE END*/

    /* Process sensor data here */
    // double result = atan2(y, x); // arctan of y / x
    
    // const double result = atan2(trans_y, 350) * -1; // arctan of Opst / Adj
    
    
    /*
    * Enter here functions to send actuator commands, like:
    * wb_motor_set_position(my_actuator, 10.0);
    */
     
    wb_motor_set_position(arm_1, theta_0_rad + angle_offset_p1);
    // wb_motor_set_position(arm_1, -angle_offset_p1 + theta_0_rad);
    wb_motor_set_position(arm_2, theta_1_rad);
     
    /* UPDATE POINT START*/
    
    
    find_point_coord_at_yz(p0, virtual_d1, theta_0_rad, angle_offset_p1, p1);
    // printf("p0 is at position: %g %g %g\n", p0[0], p0[1], p0[2]);
    // printf("new_point is at position: %g %g %g\n", new_point[0], new_point[1], new_point[2]);
    
    find_point_coord_at_xz(p1, d2, theta_1_rad, 0, p2);
    
    // double new_point[3] = {0,0,0};
    // double new_angle = 0;
    // find_point_coord_at_yz(p1, d2, new_angle, 0, new_point);
    // printf("p1 is at position: %g %g %g\n", p1[0], p1[1], p1[2]);
    // printf("new_point is at position: %g %g %g\n", new_point[0], new_point[1], new_point[2]);
    
    
    // p1[0] = p1_primeprime_i;
    // p1[1] = p1_primeprime_j;
    // p1[2] = p1_primeprime_k;
    
    // p2[0] = p2_primeprime_i;
    // p2[1] = p2_primeprime_j;
    // p2[2] = p2_primeprime_k;
    
    /* UPDATE POINT END*/
  };

  /* Enter your cleanup code here */

  /* This is necessary to cleanup webots resources */
  wb_robot_cleanup();

  return 0;
}
