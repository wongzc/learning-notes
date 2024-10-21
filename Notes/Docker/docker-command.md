1. docker run msql
    - docker pull mysql:latest
    - docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=strong_password -d mysql
    - docker exec -it test-mysql bash