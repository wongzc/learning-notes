1. a modern cli tool for interacting with MySQL
2. support 3 modes: SQL, Javascript, Python
3. to call it
    - in cmd: `mysqlsh`
4. then, select mode:
    - `\sql`
    - `\js`
    - `\py`
5. connect root
    - `\connect root@localhost`
        - then enter password
6. do whatever like `CREATE DATABASE testdb;`

7. js example
    ```javascript
    \js
    var session = mysql.getSession({user: 'root', password: 'yourpass', host: 'localhost'});
    session.createSchema('shop');
    ```

8. python example
    ```python
    \py
    session = mysql.get_session({"user": "root", "password": "yourpass", "host": "localhost"})
    session.create_schema("shop")
    ```