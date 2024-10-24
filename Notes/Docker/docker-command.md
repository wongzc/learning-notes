1. docker run msql
    - `docker pull mysql:latest`
    - `docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=strong_password -d mysql`
    - run mysql container `docker exec -it test-sql mysql -u root -p`
        - then enter password
2. docker container
    -   `docker ps`