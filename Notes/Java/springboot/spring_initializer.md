1. Project
    - Gradle
        - Gradle for build/ dependency management
    - Maven
        - use Apache maven for build and dependency management

2. Language
    - Java/ Groovy/ Kotlin

3. SpringBoot version
    - snapshot

4. Project Metadata
    - group: package namespace, more of convention to use reverse of domain like "com.wongzhaocai" etc, to create folder structure for code
    - artifact: name for JAR file, "backend-service.jar"
    - name: display name, human readable way "My Backend Service"
    - description
    - package name
    - packaging
        - jar
        - war: only used for servelet container like Tomcat
            - use in old java server
    - java version

5. dependencies
    - modules that want to include
        - spring web
        - spring reactive web: to better handle concurrency
        - Spring Data JPA: ORM
        - Spring Boot Actuator: expose endpoint to monitor app
        - H2 Database: in memory DB
        - Spring boot DevTools: dev util