//https://www.w3schools.com/java/java_methods.asp

public class Main {
  // create a method
  static void myMethod(String fname, int age) {
    System.out.println(fname + " is "+ age);
  }

  // method with return value
  static int myMethod2(int x) {
    return 5+x;
  }

  // method overloading,  method that work for differtn type
  static int plusMethod(int x, int y) {
    return x+y;
  }
  static double plusMethod(double x, double y){
    return x+y;
  }

  public static void main(String[] args) {
    myMethod("lol", 4); // calling a method
    System.out.println(myMethod2(55));

    
    int myNum1 = plusMethod(8, 5);
    double myNum2 = plusMethod(4.3, 6.26);
    System.out.println("int: " + myNum1);
    System.out.println("double: " + myNum2);
  }
}