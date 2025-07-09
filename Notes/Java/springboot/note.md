1. to create spring boot:
    1. use spring initializer
        - Project: Maven
        - Language: Java
        - Spring Boot: Choose the latest stable version
        - Group: com.example
        - Artifact: demo
        - Name: demo
        - Dependencies: Add
            - Spring Web
            - Spring Boot DevTools (optional but useful for auto-reload)
        - then generate and unzip
    2. manually create
        ```bash
        mkdir springboot-demo && cd springboot-demo
        mvn archetype:generate -DgroupId=com.example -DartifactId=demo -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
        ```
2. run springboot
    - mvnw spring-boot:run 

3. simple example
    ```java
    package com.example.demo;

    import org.springframework.web.bind.annotation.GetMapping;
    import org.springframework.web.bind.annotation.RestController;

    @RestController
    public class HelloController {
        @GetMapping("/")
        public String hello() {
            return "Hi from Spring Boot!!";
        }
    }
    ```

4. Dependency Injection to implement IoC
    - controller
        - need service object to be able to use it in controller, same for others
    - service
    - repository
    - Inversion of Control (IoC): control of object creation and dependency management is inverted ( manage by external framework)
    - Dependency Injection: Spring help to manage object lifecycle, user focus on logic
    - type of injection:
        - constructor injection: pass dependencies through class `@Autowired`  constructor
        - setter injection: `@Autowired`  use a setter method to set
        - field injection: `@Autowired` directly on a field, not recommended

5. Apache Tomcat
    - to run web application
