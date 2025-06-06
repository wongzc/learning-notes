1. java:
    - platform independent
    - object oriented 
    - garbage collector ( auto memory management)
    - strong typing
    - rich standard library

2. primitive data type vs objects
    - storage:
        - primitive store actual value
        - object store reference
    - memory
        - less for primitive
        - more for object
    - built-in operation
        - less for primitive
        - object more
    - null
        - no for primitive
        - object ok

3. String vs StringBuilder vs StringBuffer
    - String: 
        - immutable 
        - change on it is slow due to new object created and old object will be in memory until clean by GC.
    - StringBuilder:
        - mutable
        - fast in change
        - but not thread safe
    - StringBuffer:
        - mutable
        - slower change
        - thread safe
    - common method for StringBuffer and StringBuilder
        - add text at end: `sb.append("World")`
        - insert text: `sb.insert(5,"lol")`
        - delete (include start, exclude end): `sb.delete(5,10)`
        - delete char at index: `sb.deleteCharAt(2)`
        - replace text: `sb.replace(0,5,"lol")`
        - reverse: `sb.reverse()`
        - change char at index: `sb.setCharAt(0.'J')`
        - trim or pad with null char (u\0000): `sb.setLength(10)`

4. modifiers
    - class
        - `public`: access from any class ( if different package need import)
        - default ( no modifier):  access from same package only ( no need import)
        - `final`: cannot be inherit (subclassed by other)
        - `abstract`: must be subclassed, cannot directly instantiated
    - inner class
        - all from outer class
        - `private`: access from enclosing class
            - if outer class imported or create instance, still cant call
        - default ( no modifier): access within same package
        - `protected`: access from package + subclass
            - subclass means extends
        - `public`: access from anywhere ( if outer accessible)
        - `static`: 
            - tied to class, not instance
                `Outer.Inner innerInstance = new Outer.Inner();`
            - without `static`: tied to instance, not class
                `Outer.Inner innerInstance = outerInstance.new Inner();`

5. `==` vs `.equals()`
    - `==`: 
        - Primitive: compare value ( int, double, char)
        - Objects: compare object reference (memory address)
    - `.equals()`: compare content of strings

6. Array vs ArrayList
    0. difference
        - Array:
            - fixed size
            - primitive or objects
            - built-in, no need import
            - faster access
            - only length and index access method
        - ArrayList
            - sizable
            - only object
            - need `import java.util.ArrayList;`
            - slower
            - can add, remove, contains etc
    1. create:
        - `int[] arr = new int[3];`
        - `ArrayList<Integer> list = new ArrayList<>();`
    2. assign:
        - `arr[0]=10;`
        - `list.add(10);` or `list.set(0,10);`
    3. access:
        - `arr[0]`
        - `list.get(0)`
    4. size/length
        - `arr.length`
        - `list.size()`
    5. both can use `for (int i: xx)`
    6. remove
        - not for Array
        - `list.remove(index)`
    7. other:
        - `List<Integer> list = Arrays.asList("a", "b", "c")`
        - this is list, similar to array, but implement List interface
        - so can use get, set, forEach etc
        - cannot resize

7. constructor
    - special method that same name as class
    - no return type even `void`
    - can be overloaded with different param
        - can call one from another using `this(param1, param2);`
    - identify 1 from another using method signature, which has:
        1. param type
        2. param order
        3. param count

8. method overloading
    - consider method signature only
        - does not consider return type!

9. convert string to integer
    - `int num = Integer.parseInt(str);`

10. enum
    - can have field, constructors, method like class
    - enum constructor are implicitly private
    ```Java
    public enum Day {
        MONDAY("Weekday"), // use constructor to set type for each constant
        SATURDAY("Weekend"),

        private final String type;  // field

        // constructor
        Day(String type) {
            this.type = type;
        }

        // method
        public String getType() {
            return type;
        }
    }

    public class TestEnum {
        public static void main(String[] args) {
            Day d = Day.MONDAY;
            System.out.println(d);               // Output: MONDAY
            System.out.println(d.getType());     // Output: Weekday
        }
    }
    ```