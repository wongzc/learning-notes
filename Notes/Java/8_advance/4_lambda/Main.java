
import java.util.ArrayList;
import java.util.function.Consumer;

interface StringFunction { // interface with only 1 method
    String run(String str);
}
public class Main {

    public static void main(String[] args) {
        ArrayList<Integer> numbers = new ArrayList<Integer>();
        numbers.add(5);
        numbers.add(9);
        numbers.add(8);
        numbers.add(1);

        // lambda
        // format:
            // parameter -> expression
            // (parameter1, parameter2) -> expression
            // (parameter1, parameter2) -> { code block }
        numbers.forEach((n) -> {System.out.println(n);});
        
        // use Consumer to store lambda
        // lambda can be store in variable if variable type is interface which has only 1 method
            //lambda should have same parameter and return type as the methor
        Consumer<Integer> method = (n) -> { System.out.println(n); };
        numbers.forEach(method);

        // defined lambda variable, type as interface
        StringFunction exclaim = (s) -> s+"!";
        StringFunction ask = (s) -> s+"?";
        printFormatted("Hello", exclaim);
        printFormatted("Hiiii", ask);

        // this also works
        System.out.println(exclaim.run("shubydubi"));

    }

    public static void printFormatted(String str, StringFunction format) {
        String result = format.run(str);
        System.out.println(result);
    }
}
