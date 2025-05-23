
import java.util.ArrayList;
import java.util.Collections;

public class Main {

    public static void main(String[] args) {
        ArrayList<String> cars = new ArrayList<String>();
        cars.add("Volvo");
        cars.add("BMW");
        cars.add("Ford");
        cars.add(0, "Mazda"); // add element at specific position

        System.out.println(cars);

        // access  item
        cars.get(0);

        // change item
        cars.set(0,"dasdas");

        // remove item
        cars.clear();

        //get size
        cars.size();

        //loop through arraylist
        for (int i=0;i<cars.size();i++) {
            System.out.println(cars.get(i));
        }

        // loop through with for each
        for (String car:cars) {
            System.out.println(car);
        }

        // array list with integer
        ArrayList<Integer> myNum = new ArrayList<>();
        myNum.add(123);
        myNum.add(1);
        myNum.add(456);
        
        System.out.println(myNum);

        // sort array list
        Collections.sort(myNum);
        System.out.println(myNum);

        // sort with reverse order
        Collections.sort(myNum, Collections.reverseOrder());
        System.out.println(myNum);

    }
}
