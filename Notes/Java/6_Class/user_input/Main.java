import java.util.Scanner;  
// Scanner clss: to get user input

class Main {
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);  // Create a Scanner object
    System.out.println("Enter username");

    String userName = myObj.nextLine();  
    System.out.println("Username is: " + userName);  

    // nextBoolean()	Reads a boolean 
    // nextByte()	Reads a byte 
    // nextDouble()	Reads a double 
    // nextFloat()	Reads a float 
    // nextInt()	Reads a int 
    // nextLine()	Reads a String 
    // nextLong()	Reads a long 
    // nextShort()	Reads a short 
  }
}