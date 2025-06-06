1. create a map
```java
Map<Character, Character> pairs = Map.of(
            ')', '(', 
            ']', '[', 
            '}', '{'
        );

Map<Integer, Integer> map = new HashMap<>();
map.put(key, value);
map.get(key);
map.containsKey(key);
map.containsValue(value);
map.remove(key);
map.size();
map.isEmpty();
map.clear();
map.keySet();
map.values();
map.entrySet();
```

2. loop through char
```Java
for (char c : s.toCharArray()) {...}
```

3. create stack
```Java
Deque<Character> stack = new ArrayDeque<>();
stack.push(10);
stack.pop();
stack.isEmpty();
stack.peep(); // get without remove
```

4. create array list (class)
```Java
ArrayList<Integer> list = new ArrayList<>();

list.add(10);
list.get(0);
list.set(1, 30);
list.remove(0); //remove and shift
list.size();
list.isEmpty();
```

5. create list ( interface)
```Java
List<Integer> list = new ArrayList<>();
// same method as array list
```
6. create array
```Java
int[] numbers = new int[5]; // empty array of length 5
int[] numbers = {1, 2, 3, 4, 5}; //init with value
```

6. create hashset
```Java
HashSet<Integer> set = new HashSet<>();
set.add(xxx);
set.remove(xxx);
set.contains(xxx);
set.size();
set.isEmpty();
set.clear();
// can use for each to set also
```

7. using `this`: refer to instance 
    1. when instance variable has same with local var/ param: use `this.xx=xx` to mean that we are assigning the instance var
    2. call another constructor in same class
        ``` Java
        public class Person {
            String name;
            int age;

            public Person(String name) {
                this(name, 0); // Calls the other constructor
            }

            public Person(String name, int age) {
                this.name = name;
                this.age = age;
            }
        }
        ```
    3. pass current object to another method/ constructor
        ```Java
        class Engine {
            Car owner;

            public void setOwner(Car car) {
                this.owner = car;
            }
        }

        class Car {
            Engine engine = new Engine();

            public Car() {
                engine.setOwner(this); // Pass the current Car object to the Engine
            }
        }
        ```
    4. Access member from inner class
        ```Java
        public class Outer {
            int value = 10;

            class Inner {
                int value = 20;

                public void show() {
                    System.out.println(this.value);        // Refers to Inner.value
                    System.out.println(Outer.this.value);  // Refers to Outer.value
                }
            }
        }

        ```
8. Java generics
    - type-safe code that work for multiple type
    - cannot use for primitive type
    - no runtime type info, Java erase type info 
    ``` Java
    public class Box<T> {
        T value;

        public void set(T value) {
            this.value = value;
        }

        public T get() {
            return value;
        }
    }

    Box<String> b = new Box<>(); // type defined here
    b.set("Hello"); 
    String s = b.get(); // no need to cast type
    ```