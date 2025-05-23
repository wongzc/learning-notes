
import java.util.HashMap;

public class Main {

    public static void main(String[] args) {
        HashMap<String, String> cities = new HashMap<>();
        // hashmap is key value pair

        // add new item
        cities.put("England", "London");
        cities.put("Germany", "Berlin");
        cities.put("Norway", "Oslo");
        cities.put("USA", "Washington DC");
        System.out.println(cities);

        // access item
        cities.get("England");

        // get size
        cities.size();

        // loop through hashmap key
        for (String k : cities.keySet()) {
            System.out.println(k);
        }

        // loop through hashmap value
        for (String k : cities.values()) {
            System.out.println(k);
        }

        // remove item
        cities.remove("England");

        // remove all item
        cities.clear();
    }
}
