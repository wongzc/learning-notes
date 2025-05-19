class DogNoConstructor {
    String name;
    int age;

    void printInfo() {
        System.out.println("DogNoConstructor -> name: " + name + ", age: " + age);
    }
}

class DogWithConstructor {
    String name;
    int age;

    // Custom constructor
    public DogWithConstructor(String name, int age) {
        this.name = name;
        this.age = age;
    }

    void printInfo() {
        System.out.println("DogWithConstructor -> name: " + name + ", age: " + age);
    }
}

public class TestDogs {
    public static void main(String[] args) {
        // Using the default constructor (fields will be null/0 unless set manually)
        DogNoConstructor dog1 = new DogNoConstructor();
        dog1.name = "Lucky";
        dog1.age = 3;
        dog1.printInfo();

        // Using the constructor to set values
        DogWithConstructor dog2 = new DogWithConstructor("Buddy", 5);
        dog2.printInfo();
    }
}
