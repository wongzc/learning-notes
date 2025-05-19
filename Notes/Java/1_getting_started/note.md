https://dev.java/learn/getting-started/

1. Java use compiler to transform code to binary ( executable)
    - binary contain bytecode
2. after compile can execute
3. compile and execute use 2 specific software from Java development Kit ( JDK)
    - from Java SE11 onwards can merger compile and execute together by executing .java file
4. a Java Class must be saved in a file that has the same name as class, with extension .java
5. class can be any name that not start with number
    - convention: start java class with capital letter
6. Compilation
    - transform xx.java into xx.class
    - download java means download JDK
        - officially distributed by OpenJDK project & Oracle
    - JRE: Java Runtime Environment
        - subset of JDK, not distributed by OpenJDK/ Oracle anymore
        - only contain tool needed to run Java app
        - cannot compile code
    - J2EE/ Java EE/ Jakarta EE
        - acronym to Java Enterprise Edition
        - set of tools/ library to create enterprise-class applications
7. setting up JDK
    - https://jdk.java.net/
    - then unzip the downloaded file
    - and in cmd type `set JAVA_HOME=C:\.... ( the path)`
    - then update `PATH` environment with `set PATH=%JAVA_HOME%\bin;%PATH%`
8. compile class
    - check java version `java -version`
        - if error, need to check `JAVA_HOME` and `PATH`
    - run `javac MyFirstClass.java` to compile
9. adding code to class to run
    - after compile, run with  `java MyFirstClass`
10. from Java SE11 onward can run with: `java MyFirstClass.java`
    - as long as it is single file
11. Common Problem and solution
    - `javac is not recognized as an internal or external command, operable program or batch file`
        - compiler problem, cant find javac
        - use full path like `C:\jdk22\bin\javac HelloWorldApp.java`
    - `Class names, HelloWorldApp, are only accepted if annotation processing is explicitly requested`
        - forgot to include `.java` when compile