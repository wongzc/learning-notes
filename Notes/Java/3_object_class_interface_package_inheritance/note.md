https://dev.java/learn/oop/

1. Object
    - software bundle of related state and behavior
    - 2 characteristic of object
        1. state: store in fields
        2. behavior: expose through methods
            - operate on state 
            - primary mechanism for object-object communication
    - hide state and requiring all interaction through method is `data encapsulation`
    - benefit of object
        1. modularity: code for object can be written and maintained independently
        2. information-hiding: interact only with methods, internal implementation is hidden
        3. code re-use
        4. pluggability and debugging

2. Class
    - blueprint of object
    - if we define a bicycle class, the class wont have `main()` method
        - because it is not a complete application, it is just blueprint for bicycles
        - responsibility of creating and using the class belongs to other class

3. inheritance
    - inherit commonly used state and behavior from other class
    - each class only 1 direct superclass ( single inheritance) 
    - use `extends` to inherit

4. Interface
    - interface is a group of related methods with empty bodies
    - use  `implements` to implement an interface
    - can multiple inheritance
    - use to define a contract ( set of methods,  like python abstract class)
    - diff with abstract class in python:
        - cannot have field, except `public static final`
            - to define constant `public static final double PI = 3.14159;`
                - `final`: value cant be change after init
        - cannot have constructor

5. package
    - namespace that organize set of related class and interface
    - class library provided by Java ( known as API)
    