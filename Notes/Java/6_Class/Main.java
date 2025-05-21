//https://www.w3schools.com/java/java_classes.asp

// top level class: class not declared inside other class
// can only be no modifier or public
public class Main {// create a class name Main

    int x = 5; // attribute
    final int y = 1999; // declare attribute with final to make it constant

    // static method, can access without create object of class
    static void myMethod() { // create a method
        System.out.println("Hiii");
    }

    // public method, any class or package
    public void myMethod2() {
        System.out.println("Hello");
    }

    public static void myMethod3() {
        System.out.println("Hello");
    }

    void myMethod4() {
        System.out.println("Hello");
    }

    int z; //create a class attribute

    public Main() { // constructor name must match class name, and have no return type like void
        z = 3456; // set initial value for class attribute
    }

    public static void main(String[] args) {
        Main myObj = new Main(); // create an object from the Main class
        // format: ClassName objectName = new ClassName();
        System.out.println(myObj.x);

        // update the attribute value
        myObj.x = 1000;
        System.out.println(myObj.x);

        Another another = new Another(); // can run class from another file
        System.out.println(another.lala);

        Main.myMethod(); // static can call like this (if same package)

        myObj.myMethod2(); // call a public method ( can be call in other package)

        Main.myMethod3(); // public static can call like this (any package)

        myObj.myMethod4(); // call a default

        // static can call like this if same class
        myMethod();
        myMethod3();

        System.out.println(myObj.z);

        abc testabc = new abc(100);
        System.out.println(testabc.cde);
    }
}

class abc {

    int cde;

    public abc(int efg) { // constructor with parameter
        cde = efg;
    }
}
