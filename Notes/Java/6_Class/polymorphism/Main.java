class Animal {
  public void animalSound() {
    System.out.println("The animal makes a sound");
  }
}

class Pig extends Animal {
  public void animalSound() {
    System.out.println("The pig says: wee wee");
  }
}

class Dog extends Animal {
  public void animalSound() {
    System.out.println("The dog says: bow wow");
  }
}
// polymorphism
// means many form. when have many class related to each other by inheritance

// run time poplymorphism
// by method overloading, need inherit
class Main {
  public static void main(String[] args) {
    Animal myAnimal = new Animal();  // Create a Animal object
    Animal myPig = new Pig();  // Create a Pig object
    Animal myDog = new Dog();  // Create a Dog object
    myAnimal.animalSound();
    myPig.animalSound();
    myDog.animalSound();
  }
}


// this is compile time polymorphism
// by method overloading, no need inherit
class MathOps {
  int add(int a, int b) { return a + b; }
  double add(double a, double b) { return a + b; }
}