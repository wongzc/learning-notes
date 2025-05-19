https://dev.java/learn/single-file-program/
1. From JDK 11 onward, can launch single file source code program with java launcher 
    - without need of explicit compile
    - compiled code in-memory
2. to execute single file source code program
    - class defined in source file must contain `public static void main(String[])`

3. pass arguments
    - can pass argument
    - run with `java HelloJava.java World!`
4. multiple class in same file
    - can be put in same file for encapsulation purpose

5. Reference JDK Classes and Non-JDK Classes
    - class that is in core JDK not need to be added to classpath to execute, can just import and use
        - like `java.util.Scanner` or `java.util.regex.MatchResult`
    - for non core JDK class:
        - like `org.apache.commons.lang3.RandomUtils`
        - need to have the `commons-lang.jar` added to classpath to launch
        - to compile 
            - `javac -cp /path/to/commons-lang3-3.12.0.jar ReferenceNonJDKClass.java`
            - need to add it as classpath
        - when run
            - `java -cp /path/to/commons-lang3-3.12.0.jar ReferenceNonJDKClass`
