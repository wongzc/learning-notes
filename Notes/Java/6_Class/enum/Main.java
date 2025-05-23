
enum Level {
    LOW,
    MEDIUM,
    HIGH
} // to represent group of constant that cannot be change

public class Main {

    public static void main(String[] args) {
        Level myVar = Level.MEDIUM;
        // myVar will only be value that exist in Level, so it is type safe

        switch (myVar) {
            case LOW:
                System.out.println("Low level");
                break;
            case MEDIUM:
                System.out.println("Medium level");
                break;
            case HIGH:
                System.out.println("High level");
                break;
        }

        // loop through enum
        for (Level xx : Level.values()) {
            System.out.println(xx);
        }

        // loop thtough enum Planet
        for (Planet p : Planet.values()) {
        System.out.println(p.info());
        }
    }
}

//enum with constructor and methods
enum Planet {
  MERCURY(3.3e23, 2.4e6),
  EARTH(5.97e24, 6.4e6),
  MARS(6.4e23, 3.4e6);

  private final double mass;   // in kilograms
  private final double radius; // in meters

  // Constructor (must be private or package-private)
  Planet(double mass, double radius) {
    this.mass = mass;
    this.radius = radius;
  }

  public double surfaceGravity() {
    final double G = 6.67430e-11;
    return G * mass / (radius * radius);
  }

  public String info() {
    return this.name() + ": Gravity = " + surfaceGravity();
  }
}
