
import java.util.ArrayList;
import java.util.Iterator;

public class Main {

    public static void main(String[] args) {

        ArrayList<String> cars = new ArrayList<>();
        cars.add("Volvo");
        cars.add("BMW");
        cars.add("Ford");
        cars.add("Mazda");

        // Get the iterator
        Iterator<String> it = cars.iterator();

        // Print the first item
        System.out.println(it.next());

        // loop through iterator
        while (it.hasNext()) {
            System.out.println(it.next());
        }

        // iterator to remove item from collection
        // cannot use for loop to remove item, as collection changing size!
        ArrayList<Integer> numbers = new ArrayList<>();
        numbers.add(12);
        numbers.add(8);
        numbers.add(2);
        numbers.add(23);
        Iterator<Integer> it2 = numbers.iterator();
        while (it2.hasNext()) {
            Integer i = it2.next();
            if (i < 10) {
                it2.remove();
            }
        }
        System.out.println(numbers);
    }
}
