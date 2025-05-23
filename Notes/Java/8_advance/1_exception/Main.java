
public class Main {

    public static void main(String[] args) {
        // use try...catch(Exception e) {} finally{}
        try {
            int[] myNumbers = {1, 2, 3};
            System.out.println(myNumbers[10]);
        } catch (Exception e) {
            System.out.println("Something went wrong.");
        } finally {
            System.out.println("The 'try catch' is finished.");
        }
    }

    static void checkAge(int age) {
        if (age < 18) {
            // use throw to create custom error
                // ArithmeticException, FileNotFoundException, ArrayIndexOutOfBoundsException, SecurityException
            throw new ArithmeticException("Access denied - You must be at least 18 years old.");
        } else {
            System.out.println("Access granted - You are old enough!");
        }
    }
}
