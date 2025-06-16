1. java:
    - platform independent
    - object oriented 
    - garbage collector ( auto memory management)
    - strong typing
    - rich standard library

2. primitive data type vs objects
    - storage:
        - primitive store actual value
        - object store reference
    - memory
        - less for primitive
        - more for object
    - built-in operation
        - less for primitive
        - object more
    - null
        - no for primitive
        - object ok

3. String vs StringBuilder vs StringBuffer
    - String: 
        - immutable 
        - change on it is slow due to new object created and old object will be in memory until clean by GC.
    - StringBuilder:
        - mutable
        - fast in change
        - but not thread safe
    - StringBuffer:
        - mutable
        - slower change
        - thread safe
    - common method for StringBuffer and StringBuilder
        - add text at end: `sb.append("World")`
        - insert text: `sb.insert(5,"lol")`
        - delete (include start, exclude end): `sb.delete(5,10)`
        - delete char at index: `sb.deleteCharAt(2)`
        - replace text: `sb.replace(0,5,"lol")`
        - reverse: `sb.reverse()`
        - change char at index: `sb.setCharAt(0.'J')`
        - trim or pad with null char (u\0000): `sb.setLength(10)`

4. modifiers
    - class
        - `public`: access from any class ( if different package need import)
        - default ( no modifier):  access from same package only ( no need import)
        - `final`: cannot be inherit (subclassed by other)
        - `abstract`: must be subclassed, cannot directly instantiated
    - inner class
        - all from outer class
        - `private`: access from enclosing class
            - if outer class imported or create instance, still cant call
        - default ( no modifier): access within same package
        - `protected`: access from package + subclass
            - subclass means extends
        - `public`: access from anywhere ( if outer accessible)
        - `static`: 
            - tied to class, not instance
                `Outer.Inner innerInstance = new Outer.Inner();`
            - without `static`: tied to instance, not class
                `Outer.Inner innerInstance = outerInstance.new Inner();`

5. `==` vs `.equals()`
    - `==`: 
        - Primitive: compare value ( int, double, char)
        - Objects: compare object reference (memory address)
    - `.equals()`: compare content of strings

6. Array vs ArrayList
    0. difference
        - Array:
            - fixed size
            - primitive or objects
            - built-in, no need import
            - faster access
            - only length and index access method
        - ArrayList
            - sizable
            - only object
            - need `import java.util.ArrayList;`
            - slower
            - can add, remove, contains etc
    1. create:
        - `int[] arr = new int[3];`
        - `ArrayList<Integer> list = new ArrayList<>();`
    2. assign:
        - `arr[0]=10;`
        - `list.add(10);` or `list.set(0,10);`
    3. access:
        - `arr[0]`
        - `list.get(0)`
    4. size/length
        - `arr.length`
        - `list.size()`
    5. both can use `for (int i: xx)`
    6. remove
        - not for Array
        - `list.remove(index)`
    7. other:
        - `List<Integer> list = Arrays.asList("a", "b", "c")`
        - this is list, similar to array, but implement List interface
        - so can use get, set, forEach etc
        - cannot resize

7. constructor
    - special method that same name as class
    - no return type even `void`
    - can be overloaded with different param
        - can call one from another using `this(param1, param2);`
    - identify 1 from another using method signature, which has:
        1. param type
        2. param order
        3. param count

8. method overloading
    - consider method signature only
        - does not consider return type!

9. convert string to integer
    - `int num = Integer.parseInt(str);`

10. enum
    - can have field, constructors, method like class
    - enum constructor are implicitly private
    ```Java
    public enum Day {
        MONDAY("Weekday"), // use constructor to set type for each constant
        SATURDAY("Weekend"),

        private final String type;  // field

        // constructor
        Day(String type) {
            this.type = type;
        }

        // method
        public String getType() {
            return type;
        }
    }

    public class TestEnum {
        public static void main(String[] args) {
            Day d = Day.MONDAY;
            System.out.println(d);               // Output: MONDAY
            System.out.println(d.getType());     // Output: Weekday
        }
    }
    ```

11. abstract class vs interface
    - abstract:  ( use as common base)
        - can have concrete and abstract method
            - leave it no body to be abstract
        - single inherit only
        - constructor allowed
        - any field ok
    - interface: (use as common properties)
        - method implicitly `abstract` and `public`
            - but can have:
                - `default`: with body ( class can override)
                - `static`: with body (call directly on interface)
                - `private`: only internal use in interface!
        - multiple inheritance
        - no constructor allowed
        - field is implicit `public static final`

12. autoboxing and unboxing
    - autoboxing: convert primitive type to wrapper class
    - unboxing: object to primitive
    ```Java
    int primitive = 5;
    Integer wrapped = primitive;  // Autoboxing
    int unboxed = wrapped;        // Unboxing
    ```

13. runnable vs callable interface in concurrency
    - runnable:
        - don't return result, don't throw exception, just run
    - callable:
        - return result, throw exception

14. ArrayList vs LinkedList
    - ArrayList
        - fast random access ( direct access in O(n))
        - slow insert/ delete
    - LinkedList
        - fast insert/ delete
        - slow random access ( traverse node by node)

15. singleton pattern and thread safe implementation in java
    1. enum singleton
    - handle thread-safe, serialize, reflection attack
    ```java
    public enum MySingleTon {
        INSTANCE;
        public void doSomething() {
            ...
        }
    }
    ```
    ```java
    MySingleTon.INSTANCE.doSomething(); //usage
    // no need instantiate, directly use
    ```

    2. Eager initialization
    - create instance when class loaded
    - thread safe
    - simple implementation
    ```java
    public class MySingleton {
        private static final MySingleton instance = new MySingleTon();

        private MySingleton() {
            // constructor
        }
        public static MySingleton getInstance() {
            return instance;
        }
    }
    ```
    ```java
    MySingleton singleton = MySingleton.getInstance();
    // never call new MySingleton()!
    singleton.xxx();
    ```

    3. Lazy intialization
    - instantiated only when needed ( save resource)
    - volatile+ synchronized to ensure thread safe
        - volatile: ensure visibility, other thread can immediate see when this is updated
    ```java
    public class MySingleton {
        private static volatile MySingleton instance;

        private MySingleton() {}

        public static MySingleton getInstance() {
            // double checked locking (DCL)
            if ( instance == null) {
                // only enter if instance null
                // multiple can enter
                synchronized (MySingleton.class) {
                    // lock with synchronized
                    // so only 1 thread can enter here
                    if (instance == null) {
                        // check again incase mutiple thread enter the outer instance==null
                        instance = new MySingleton();
                    }
                }
            }
            return instance;
        }
    }
    ```
    ```java
    MySingleton singleton = MySingleton.getInstance();
    // never call new MySingleton()!
    singleton.xxx();
    ```

    4. enum is simple and good, use non-enum when need lazy loading, dependencies injection, flexibility

16. reflection
    - feature to inspect and manipulate class, method, field, constructor at run time, even they are private.
    - risk
        - break encapsulation and force creation of multi instance of singleton by access private constructor.
        ```Java
        Constructor<MySingleton> ctor = MySingleton.class.getDeclaredConstructor(); // get the constructor using reflection API
        ctor.setAccessible(true); // bypass private access
        MySingleton another = ctor.newInstance(); // creates another instance!
        ```


17. Serialization
    - converting java object to byte stream
        - to save to disk or send over network
    - risk:
        - a new object create when serialize/ deserialize
        - violate singleton pattern ( if it is singleton)
    - use readResolve() to solve
        - during deserialize, if readResolve() is present, java replace the object with return value of readResolve()
        - problem only when deserialize, when object created form bytestream
    ```java
    protected Object readResolve() {
        return getInstance();  // return the original instance
    }
    ```

18. `volatile`
    - ensure write by 1 thread are immediately visible to others
    - prevent instruction reordering during object creation

19. fail-fast vs fail-safe iterators
    - fail-fast: throw `ConcurrentModificationException` if collection modified while iterating.
        - ArrayList
        - HashMap
        - HashSet
        - ensure data consistency
    - fail-safe: work on clone of iterators
        - ConcurrentHashMap
        - CopyOnWriteArrayList
        - allow collection modification when iterate

20. Java Memory Model (JMM)
    - rule for thread interact through memory
    - tell JVM how to handle memory visibility & ordering of variable
    - ensure
        - visibility: when 1 thread update variable, another eventually see the update, through `volatile`
        - Atomicity: operations like `volatile`, `synchronized` prevent partial update
        - Ordering: define what order of operations is allowed to be observed by different threads
    - JMM tools:
        - `volatile`: ensure visibility
        - `synchronized`: atomicity+ visibility+ordering
        - `final`: help to create immutable objects
        - `java.util.concurrent` tools: `AtomicInteger`, `ReentrantLock`

21. hashCode() relation with equals()
    - if `a.equals(b)` returns `true`
        - then `a.hashCode() == b.hashCode()`
    - but if `a.hashCode() == b.hashCode()` doesn't means a==b
        - hash collisions can happen
    - if override `.equals()`, need to override `.hashCode()` as well
        - `.contains()`, `.get()`, `.remove()` all use `.hashCode()`
        - override happen when want to have specific way of equal of object, like name equal etc, then need to set hashCode to just return name.hashCode()!

22. all java object implicitly extend `java.lang.Object`
    - it has method like `equals()`, `hashCode()`, `toString()`

23. `@Override`
    - override class method: optional but recommended
    - override interface method: must

24. method references
    - shorthand for lambda
    - 4 type:
        1. static method
        2. instance method of a object
        3. instance method of arbitrary object of specific type
        4. constructor
    - example
    ```java
    List<String> names = List.of("a","b","c");

    names.forEach(name -> System.out.println(name)); //lambda
    names.forEach(System.out.println); // method reference
    ```

    static method
    ```java
    Function<Double, Double> absFunc = x -> Math.abs(x); //lambda
    Function<Double, Double> absFunc = Math::abs; //method ref
    absFunc.apply(-5.0);
    ```
25. Observer pattern
    - when object change state, all dependent objects will be notified
    ```java
    NewsAgency agency = new NewsAgency();
    Observer reader1 = new NewsReader("Alice");
    Observer reader2 = new NewsReader("Bob");

    agency.registerObserver(reader1);
    agency.registerObserver(reader2);

    agency.setNews("New Java version released!");
    ```
    - implement using event bus

26. comparable vs comparator 
    - comparable:
        - implement inside class to define default order
        - need to override compareTo
        ```java
        class Person implements Comparable<Person> {
            String name;
            int age;

            public Person(String name, int age) {
                this.name = name;
                this.age = age;
            }

            @Override
            public int compareTo(Person other) {
                return this.age - other.age; // sort by age
            }
        }

        List<Person> people = List.of(new Person("Alice", 30), new Person("Bob", 25));
        Collections.sort(people); // Uses compareTo()
        ```
    - comparator:
        - define comparator outside of class
        ```java
        Comparator<Person> nameComparator = (p1, p2) -> p1.name.compareTo(p2.name);
        Comparator<Person> ageComparator = (p1, p2) -> Integer.compare(p1.age, p2.age);

        List<Person> people = new ArrayList<>();
        people.add(new Person("Alice", 30));
        people.add(new Person("Bob", 25));

        Collections.sort(people, nameComparator); // Sort by name
        ```
27. try-with resource
    - auto close resource when done using, even exception is thrown
    - resource is any object that implement `java.lang.AutoCloseable` or `java.io.Closeable`
        - `FileInputStream`, `BufferedReader`, `Scanner`, `Connection`
    ```java
    try (BufferedReader reader = new BufferedReader(new FileReader("data.txt"))) {
        // instead of try{ reader=...}, it use try(reader=...) {}
        System.out.println(reader.readLine());
    } catch (IOException e) {
        e.printStackTrace();
    }
    ```
28. try-catch-finally
    - any variable use in `finally` must be declared outside of `try` block

29. HashMap vs LinkedHashMap vs TreeMap
    - HashMap
        - unordered
        - fast lookup,O(1)
        - `Map<String, Integer> map = new HashMap<>();`
        - null key ok
    - LinkedHashMap (use double linked list with hash table)
        - ordered
        - more memory usage, O(1), but need extra time to update linkedlist
        - `Map<String, Integer> map = new LinkedHashMap<>();`
        - null key ok
    - TreeMap ( red black tree)
        - sorted by key ( can pass comparator if need)
        - `Map<String, Integer> map = new TreeMap<>();`
        - not null key

30. `transient`
    - to indicate field that shouldn't be serialized

31. method overloading and method overriding
    - both are polymorphism

32. checked vs unchecked exception in Java
    - checked:
        - compiler forces you to handle
        - `try-catch` or `throws`
        - extend `Exception`
    - unchecked:
        - exception not checked at compile time
        - extend `RuntimeException`

33. `java.util.concurrent`
    - dedicated package for concurrent programming in Java
    - thread-safe, high-performance
    - `ExecutorService`: manage and execute task asynchronously
        ```java
        ExecutorService executor = Executors.newFixedThreadPool(5);
        executor.submit(() -> System.out.println("Task executed"));
        ```
    - `ConcurrentHashMap`: thread-safe hashMap
        ```java
        Map<String, Integer> map = new ConcurrentHashMap<>();
        ```
    - `CountDownLatch`: for synchronizing threads waiting for operation to be complete
        ```java
        CountDownLatch latch = new CountDownLatch(3);
        latch.countDown(); // use this to subtract 1 from count
        latch.await(); // this will block further action until latch=0
        ```
    - `AtomicInteger`: lock-free atomic operation on integer
        ```java
        AtomicInteger counter = new AtomicInteger(0);
        counter.incrementAndGet();
        ```
34. synchronized collections vs concurrent collections
    - synchronized collections
        - example:
            - `Collections.synchronizedList(new ArrayList<>())`
            - `Collections.synchronizedMap(new HashMap<>())`
        - normal collection wrap with synchronized wrapper
        - use single lock to guard all operations
        - legacy
        - 1 thread can access method at a time
            - simple, but low performance under high concurrency
    - concurrent collections
        - for high-concurrency scenario
        - fine grained locking
        - some use lock-reads
        - example:
            - `ConcurrentHashMap`
            - `CopyOnWriteArrayList`
            - `ConcurrentLinkedQueue`
            - `BlockingQueue`

35. Fork/ Join framework
    - Fork: split task into smaller subtask, concurrent execute
    - Join: combine the result

36. lock-free programming
    - thread safe without using locks
        - use low level atomic operations to avoid blocking threads
    - benefit:
        - no thread block
        - no deadlock
        - minimal context switching
        - scalable
    - use `java.util.concurrent.atomic`
        - `AtomicInteger`, `AtomicLong`, `AtomicReference`
    - some Scenarios may not be lock-free
        1. complex multi step operations
            - read, validate and write multiple variable
        2. fairness/ ordering
            - lock-free cause starvation if 1 thread keeps getting ahead
            - lock-free use CAS ( compare and swap), fast thread may keep win again and again, slow thread may starve
        3. wait-free guarantee
            - wait-free means thread is not block by other thread or retry because of other
                - lock: may block
                - lock-free: may cause some thread fail and retry at CAS

37. Compare and Swap (CAS)
    - compare the expected ( thread read ) value with current value of a variable
        - if expected = current, swap the current with new value
        - else, retry

38. ABA problems
    - when value appears unchanged, but actually changed and changed back (A -> B -> A)
    - happen in lock-free, as CAS only see values, not history
    - to solve:
        - stamping with version

39. garbage collection process tuning
    - for high throughput, low latency, memory intensive applications
    - to solve
        1. long GC pause: all thread stopped to reclaim memory
        2. frequent GC
        3. OutOfMemoryError
        4. poor CPU util: thread stop during stop-the-world (STW ) GC, so CPU util low
    - use JVM options to control GC
        - choose GC algorithm
        - heap size tuning
        - check GC logs