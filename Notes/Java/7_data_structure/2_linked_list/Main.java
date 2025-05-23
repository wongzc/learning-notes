
import java.util.LinkedList;

public class Main {

    public static void main(String[] args) {
        LinkedList<String> cars = new LinkedList<>(); // can just skip the type, it will infer from left 
        cars.add("Volvo");
        cars.add("BMW");
        cars.add("Ford");
        cars.add("Mazda");
        System.out.println(cars);

        // linkedlist vs arraylist
        // linkedlist has all method same as array list
        // as both implement list interface
        // ArrayList: if array is not big enough, new larger array will be created
        // linkedList: multiple container linked, each with a link to first and next container
            // when add, new element put into new container and linked to other

        // ArrayList for storing data, linkedlist to manipulate data
        // ArrayList: fast access by index, less memory
        // linkedList: fast data removal from center

        // Linked list also implement Deque
        // unique method for linked list
            // addFirst(E e)
            // addLast(E e)
            // removeFirst()
            // removeLast()
            // peekFirst()
            // peekLast()
    }
}
