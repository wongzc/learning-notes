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

        // passing parameter to class constructor
        Another2 xxx = new Another2(100142);
        System.err.println(xxx.z);


        //access modifier
        // class: only public or default
        // attribute/ methods/ constructor/ inner class
            // private: only access in same class
            // default: access in different class ok if same package
            // protected: access in different package ok if same it is subclass ( extend)
            // public: other class also can access, anywhere no limit
        // class->package->subclass->any

        //Non-access modifier
        // class: 
            // final: cannot be inherit
            // abstract: must be inherit
        // attribute/ methods/ constructor
            // final: cannot be overridden/ modified
            // static: attribute and methods belong to class, not object
            // abstract: only used in abstract class ( only methods, which do not have body. ( can have concrete method in abstract also))
            // transient: skip attribute and methods when serializing object containing them ( like password, thread, cache etc)
            // synchronized: method accessed by 1 thread at a time
            // volatile: value of attribute not cached thread-locally, always read from main memory

        // Java package
            // can import with package.name.* to import whole package


    }
}

class abc {

    int cde;

    public abc(int efg) { // constructor with parameter
        cde = efg;
    }
}
