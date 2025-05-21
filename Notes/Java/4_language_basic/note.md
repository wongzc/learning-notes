1.  creating variable and naming them

    - 4 type of variable
      1. instance variable( non-static variable)
         - field declared without `static` inside class
      2. class variable ( static variable)
         - with `static`
         - all instance share the same copy of variable
         - if dont want change can add `final`
      3. local variable
         - no special keyword
         - only exist between opening and closing brace of method
      4. Parameter
         - parameters are variables, not field
         - passed into class, method
    - naming variable
      - variable case sensitive
      - must be begin with letter, $ or \_
        - but avoid using $ and \_
      - can contain letter, number, \_ or $
      - camelCase

2.  primitive types

    - java static typed, all variable must be declared before used
    - 7 primitive data type
      1. byte: 8 bit signed integer, 2 complement (-128~127)
         - 2 complement: for -ve, invert all bits and +1
           - allow positive and negative number use same binary addition logic
      2. short: 16 bit signed 2 complement (-2^15 ~ 2^15-1)
      3. int: 32 bit signed 2 complement (-2^31 ~2^31-1)
      4. long: 64 bit signed 2 complement (-2^63~2^63-1)
      5. float: single precision 32-bit IEEE 754 floating point
         - can use if want to save memory
         - dont use for precise value like currency
           - use `java.math.BigDecimal` instead
      6. double: double precision 64-bit IEEE 754 floating point
      7. boolean: true or false
      8. char: single 16-bit Unicode character, from `\u0000` to `\uffff`

3.  initialize variable with default value

    - field that declared but not initialized will be set to a default value
      - byte/short/int: 0
      - long: 0L
      - float: 0.0f
      - double: 0.0d
      - char: \u0000
      - string/object: null
      - boolean: false
    - local variable will not be assigned default value, must always manually assign

4.  creating value with literals

    - for primitive type can direct assign literals to variable, no need `new`

5.  Integer Literals

    - `long`: integer that end with `L` or `l` ( use `L` better)

6.  floating point literal

    - `float`: end with `f` or `F`
      - with scientific notation: `float f1 = 1.23e2f`
    - `double`: optionally end with `d` or `D`

7.  character and string literals

    - char and string can contain any Unicode character (UTF-16)
    - can use unicode escape also , example: `\u0108`
    - support other escape: `\b (backspace), \t (tab), \n (line feed), \f (form feed), \r (carriage return), \" (double quote), \' (single quote), and \\ (backslash)`
    - `null`

8.  SE7 onward, can use `_` between digits in numerical literal

    - `float pi =  3.14_15F`
    - cannot be beginning or end of number
    - cannot be adjacent to decimal point or `L` or `F`

9.  Array

    - length of array is fixed
      - item in array is element
    - create array
      - declare with `int[] anArray;`
        - like this also can `int anArray[];`
      - allocate memory `anArray = new int[10];`
      - then use index to get or assign value `anArray[0]=123;`
      - create and init `int[] anArray = {1,2,3,4}`
      - array of array `String[][]`
      - declare and allocate `int[] anArray = new int[10]`
    - copy array
      - `System.arraycopy(copyFromArray, start, copyToArray, destArrayPos, length)`

10. var type identifier - start from SE10, can use `var` for local variable 
    - compiler decide type of variable 
    - `var message = "hii";`

    ```Java
    var list = List.of("one", "two", "three", "four");

    for (var element: list) {
    System.out.println(element);
    }

    ```
    - can only use for local variable declared in methods, constructor, initializer block
    - not for field
    - cannot `var xxx=null` as null dont have type!

11. Operator
    - can use `+` to concate string
    - `x++` or `x--`
    - `&&` `||`
    - conditional: `xx =yy? aa:bb;`
    - check if instance of a class `xxx instanceof YYY`
        - null is not an instance of anything
    
12. Expression, statement, block
    - floating point: avoid using `==` for double as rounding may make the value not equal
    - statement:
        - assignment expression
        - ++ or --
        - method invocations
        - object creation
        - declaration statement
        - control flow statement
    - block:
        - 0 or more statement between braces

13. Control flow statement
    - `if (xx) {yy;} else if (zz) {aa;} else {bb;}`
    - `while (xx) {yy;}`
    - ` do {yy;} while (xx);` : at least execute do once
    - `for(int i=0; i<10,i++) {yy;}`
        - 3 expression optional, can be like `for ( ; ; )` for infinite loop
    - `for ( int item: anArray) {yy;}` to loop through array/ collections
    - `break`
        - unlabeled: break innermost switch, for, while, do-while
        - labeled: break the labeled statement
    ```Java
        search:
            for (i = 0; i < arrayOfInts.length; i++) {
                for (j = 0; j < arrayOfInts[i].length;
                    j++) {
                    if (arrayOfInts[i][j] == searchfor) {
                        foundIt = true;
                        break search;
                    }
                }
            }
    ```
    - `continue`
        - unlabeled: skip current iteration of for, while, do-while
        - labeled: skip current iteration of an outer loop

    - `return`
    - `yield`

14. switch
    - type of selector can only be
        1. byte, short, char, and int
        2. Character, Byte, Short, and Integer wrapper types
        3. enumerated type
        4. String type
    - boolean, long, float, and double cannot used for switch
    ```Java
    int quarter = ...; // any value

    String quarterLabel = null;
    switch (quarter) {
        case 0: quarterLabel = "Q1 - Winter"; 
                break;
        case 1: quarterLabel = "Q2 - Spring"; 
                break;
        case 2: quarterLabel = "Q3 - Summer"; 
                break;
        case 3: quarterLabel = "Q3 - Summer"; 
                break;
        default: quarterLabel = "Unknown quarter";
    };
    ```
    - use `break` in switch, else it will fall through to next case
    - the switch cannot be null else will have error

15. switch expression (SE14)
    ```Java
    Day day = ...; // any day
    int len =
        switch (day) {
            case MONDAY, FRIDAY, SUNDAY -> 6;
            case TUESDAY                -> 7;
            case THURSDAY, SATURDAY     -> 8;
            case WEDNESDAY              -> 9;
        };
    System.out.println("len = " + len);
    ```
    ```Java
    public String convertToLabel(int quarter) {
        String quarterLabel =
            switch (quarter) {
                case 0  -> {
                    System.out.println("Q1 - Winter");
                    yield "Q1 - Winter"; // use yield, not return here to avoid ambiguity
                }
                default -> "Unknown quarter";
            };
        return quarterLabel;
    }
    ```
    - switch expression can use normal case block also
    ```Java
    int quarter = ...; // any value

    String quarterLabel =
        switch (quarter) {
            case 0 :  yield "Q1 - Winter";
            case 1 :  yield "Q2 - Spring";
            case 2 :  yield "Q3 - Summer";
            case 3 :  yield "Q3 - Summer";
            default: System.out.println("Unknown quarter");
                    yield "Unknown quarter";
        };
    ```







