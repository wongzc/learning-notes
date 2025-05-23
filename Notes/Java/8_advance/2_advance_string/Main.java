// to sort list of objects, need specify rule
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;


class Car {
  public String brand;
  public String model;
  public int year;
  
  public Car(String b, String m, int y) {
    brand = b;
    model = m;
    year = y;
  }
}

// Create a comparator
class SortByYear implements Comparator {
  public int compare(Object obj1, Object obj2) {
    // Make sure that the objects are Car objects
    Car a = (Car) obj1;
    Car b = (Car) obj2;
    
    // Compare the year of both objects
    if (a.year < b.year) return -1; // The first car has a smaller year
    if (a.year > b.year) return 1;  // The first car has a larger year
    return 0; // Both cars have the same year
  }
}

// use below code to direct compare car type
// class SortByYear implements Comparator<Car> {
//   public int compare(Car a, Car b) {
//     return Integer.compare(a.year, b.year);
//   }
// }

public class Main { 
  public static void main(String[] args) { 

    ArrayList<Car> myCars = new ArrayList<>();    
    myCars.add(new Car("BMW", "X5", 1999));
    myCars.add(new Car("Honda", "Accord", 2006));
    myCars.add(new Car("Ford", "Mustang", 1970));

    // Use a comparator to sort the cars
    Comparator myComparator = new SortByYear();
    Collections.sort(myCars, myComparator);

    // Display the cars
    for (Car c : myCars) {
      System.out.println(c.brand + " " + c.model + " " + c.year);
    }


    // can use lambda expression as well
    ArrayList<Car> myCars1 = new ArrayList<>();    
    myCars1.add(new Car("BMW", "X5", 1999));
    myCars1.add(new Car("Honda", "Accord", 2006));
    myCars1.add(new Car("Ford", "Mustang", 1970));

    Collections.sort(myCars1, (o1,o2)-> {
        Car a =(Car) o1;
        Car b = (Car) o2;
        if (a.year<b.year) return -1; //front first
        if (a.year>b.year) return 1; // late first
        return 0;
    });

    // even shorter lambda
    Collections.sort(myCars1, (a, b) -> Integer.compare(a.year, b.year));

    // this works too
    myCars.sort(Comparator.comparingInt(c -> c.year));


    // sort to make even first
    ArrayList<Integer> myNumbers = new ArrayList<>(Arrays.asList(33, 15, 20, 34, 8, 12));

    myNumbers.sort((a, b) -> {
      boolean aIsEven = (a % 2) == 0;
      boolean bIsEven = (b % 2) == 0;

      if (aIsEven == bIsEven) {
        return Integer.compare(a, b);
      } else {
        return aIsEven ? -1 : 1;
      }
    });

    myNumbers.forEach(System.out::println);

  } 
}