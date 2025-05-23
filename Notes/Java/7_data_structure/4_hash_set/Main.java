
import java.util.HashSet;

public class Main {

    public static void main(String[] args) {
        HashSet<String> cars = new HashSet<>();
        cars.add("Volvo");
        cars.add("BMW");
        cars.add("Ford");
        cars.add("BMW");
        cars.add("Mazda");
        System.out.println(cars);

        // check if item exist
        cars.contains("lal");

        // hashset size
        cars.size();

        // loop through hash set
        for (String i: cars) {
            System.out.println(i);
        }

        //remove item
        cars.remove("BMW");

        // remove all
        cars.clear();
    }
}
