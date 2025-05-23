// wrapper class: use primitive data type as object
// primitive data type to Wrapper class
    // byte	    Byte
    // short	Short
    // int	    Integer
    // long	    Long
    // float	Float
    // double	Double
    // boolean	Boolean
    // char	    Character
// use when working with Collection objects like ArrayList
    // invalid: ArrayList<int> myNumbers = new ArrayList<>(); 
    // valid:   ArrayList<Integer> myNumbers = new ArrayList<>(); 

// primitive data type: is faster, no object overhead, less memory. but cant be null and use in collection
// wrapper class: can be null, have many methods, can use in collections. but slow and more memory

public class Main {

    public static void main(String[] args) {
        Integer myInt = 5;
        Double myDouble = 5.99;
        Character myChar = 'A';
        System.out.println(myInt);
        System.out.println(myDouble);
        System.out.println(myChar);

        System.out.println(myInt.intValue());
        System.out.println(myDouble.doubleValue());
        System.out.println(myChar.charValue());

        String myString = myInt.toString();
        System.out.println(myString.length());
    }
}
