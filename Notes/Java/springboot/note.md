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