1. creating variable and naming them
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
        - must be begin with letter, $ or _
            - but avoid using $ and _
        - can contain letter, number, _ or $
        - camelCase

2. primitive types
    - java static typed, all variable must be declared before used
    - 7 primitive data type
        1. byte: 8 bit signed integer, 2 complement (-128~127)
            - 2 complement: for -ve, invert all bits and +1
                - allow positive and negative number use same binary addition logic
        2. short:  16 bit signed 2 complement (-2^15 ~ 2^15-1)
        3. int: 32 bit signed 2 complement (-2^31 ~2^31-1)
        4. long: 64 bit signed 2 complement (-2^63~2^63-1)
        5. float: single precision 32-bit IEEE 754 floating point
            - can use if want to save memory
            - dont use for precise value like currency
                - use `java.math.BigDecimal` instead
        6. double: double precision 64-bit IEEE 754 floating point
        7. boolean: true or false 
        8. char: single 16-bit Unicode character, from `\u0000` to `\uffff`

3. initialize variable with default value
    - field that declared but not initialized will be set to a default value
        - byte/short/int: 0
        - long: 0L
        - float: 0.0f
        - double: 0.0d
        - char: \u0000
        - string/object: null
        - boolean: false
    - local variable will not be assigned default value, must always manually assign

4. creating value with literals
    - for primitive type can direct assign literals to variable, no need `new`

5. Integer Literals
    - `long`: integer that end with `L` or `l` ( use `L` better)

6. floating point literal
    - `float`: end with `f` or `F`
        - with scientific notation: `float f1 = 1.23e2f`
    - `double`: optionally end with `d` or `D`

7. character and string literals
    - 

        