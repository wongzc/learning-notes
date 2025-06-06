public class Person {
    String name;
    String job;
    int age;

    // Default constructor
    public Person() {
        this("Unknown", 0);  // call param constructor
    }

    // Param constructor
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    public Person(int age, String job) {
        this.job = job;
        this.age = age;
    }
}
