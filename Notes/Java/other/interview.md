1. Java use javac to compile into bytecode .class (platform independent), then run with JVM ( dependent )

2. JVM
    - Java interpreter
    - load, verify, execute bytecode
    - platform dependent

3. JIT
    - Just in Time
    - part of JRE ( Java Runtime Env)
    - JIT is part of JVM, to compile bytecode to machine code

4. Memory storage
    - Class
        - store class data
        - constant, field, method data and code
    - Heap
        - allocate memory to object during run time
        - store object ( class instance), instance field, arrays
        - manage by GC
    - Stack
        - data and partial result which needed for method and performing dynamic linking
        - store local variable, parameter, reference
    - Program Counter Register
        - store address of JVM instruction currently executing
    - Native Method Stack
        - all native method
    - MetaSpace
        - store class metadata

5. classlodaer
    - part of JRE, when execute or create .class, classloader load java class and interface to JVM.
    - so JVM dont need to know about the file and file system

6. JVM vs JRE vs JDK
    - JVM:
        - part of JRE
        - convert bytecode to machine code
    - JRE:
        - package that provide environment to run java program
    - JDK:
        - environment to develop and execute Java program
        - with Dev too + JRE

7. public static void main(String args[])
    - the main function
    - public:
        - Access specifier
        - public, so that JVM can invoke from outside of class
    - static
        - keyword
        - so that can use element without init the class
        - without this, it run fine but it will not be entry point
    - void
        - return type
        - dont return anything
    - main
        - method name
    - (String args[])
        - store Java command -line arguments
        - array of String

8. Java String Pool
    - place in heap memory that store all String defined
    - in stack store the variable that storing the string
    - when creating new string, JVM check if the string alr in pool, if yes then reuse
        - note: String is immutable!

9. Package
    - grouping of related type of class, interface etc
    - provide access to protection and namespace management
    - prevent name conflict
    - control access
        - public: can access outside of package
        - protected: can outside of package only if subclass ( extend)
        - default: same package only
        - private: same class only
    - provide hidden class
    - easier locate related class

10. type of package:
    - user defined package
    - build in package

11. data types in java
    - Primitive: value only
        - boolean
        - byte: 8 bit signed
        - char
        - short: 16 bit signed
        - int: 32 bit signed
        - long: 64 bit unsigned
        - float: 0.0f
        - double: 0.0d
    - Non-Primitive: with memory address of its value, as cannot store value directly in memory
        - Strings
        - Array
        - Class
        - Object
        - Interface

12. Pointer
    - Java dont allow
    - Security: direct memory access is dangerous, memory overflow, corrup
    - Simplicity: Java has garbage collector
    - Portability: pointer depends on machine memory layout

13. wrapper
    - object class that encapsulate primitive data types
    - Boolean, Byte, Short, Integer, Character, Long, Float, and Double
    - can create custom wrapper
    - need wrapper due to:
        - provide methods
        - autoboxing unboxing
        - final: cannot extend, ensure consistency and security
        - immutable: value cannot be change after create, when we change the value, actually recreate a new wrapper, thread-safe

14. Instance vs local vs class variable
    - Instance: in class, with default value, unique to instance
    - Local: in method, no default value
    - Class: in class, but with static, share by all instance of class

15.  System.out, System.err, and System.in
    - System.out
        - normal message output
    - System.err
        - error message output
    - System.in
        - input

16. IO Stream
    - use to write to or read from file

17. Stream
    - InputStream / OutputStream (Byte Streams)
        - handle raw binary data ( images, audio, PDFs)
        - deal with bytes ( 8 bit)
        - java.io.InputStream
        - java.io.OutputStream
    - Reader / Writer (Character Streams)
        - textual data ( character, strings, text files)
        - deal with character ( 16 bit unicode)
        - java.io.Reader
        - java.io.Writer
    - BufferedInputStream & BufferedOutputStream
        - use buffer to improve I/O when read/write binary data
        - Wrap InputStream/ OutputStream
            - read chunks data into buffer, faster than byte by byte read
            - write in buffer then write in chunks 
    - Filter streams
        - subclass of FilterInputStream, FilterOutputStream, FilterReader, or FilterWriter

18. to get input from console
    - Command line argument
    - Buffered Reader Class
    - Console Class
    - Scanner Class

19. `>>` vs `>>>`
    - `>>`: shift with sign, 1 for -ve, 0 for +ve
    - `>>>`: shift with 0

20. covariant return type
    - when extend a class and override the method, the child can return a more specific ( subclass) or original return type

21. transient
    - to skip serialize a variable, like pw, Thread

22. sleep() vs wait():
    - sleep()
        - Thread class
        - pause thread for specific time
        - not release lock/monitor
        - used for delay and throttling
        - static
        - can call anywhere
    - wait()
        - object class
        - make thread wait until another thread call notify(), notifyAll()
        - release lock/monitor
        - for inter-thread communication
        - not static
        - only call in synchronized block/method ( need monitor)

23. String vs StringBuffer vs StringBuilder
    - String: immutable, thread safe
    - StringBuffer: mutable, thread safe
    - StringBuilder: mutable, fast, not thread safe

24. Create String
    - `String s = "hello";`: create from literal, store in string pool and will resuse if needed
    - `String s = new String("hello");`: create from new(), store in Heap, slower and consume memory

25. Array:
    - store fix sized sequence in java
    - create in heap memory
    - can use either
        - `int arr[]`: c style
        - `int[] arr`: java style
    - copy:
        - shallow: use arr_from.clone() to copy
        - deep: use System.arraycopy()
    - jagged array: array of array with different length
    - array cannot be volatile

26. OOP
    - Inheritance
        - subclass inherit superclass
    - Polymorphism
        - ability to take more than 1 form
        - compile time polymorphism (method overloading)
            - define method with different input type
        - run time polymorphism (method overriding )
            - extend a class and override method
            - cannot override:
                - final: cannot change
                - static: no use. override in child without error, but when call it is still no change
                - private: no access
            - scope (access modifier) can be change
                - only same or wider
                - public -> public
                - protected -> protected, public
                - default -> default, protected, public
                - private: cannot inherit!!
    - Abstraction
        - act of represent essential features without including background details
    - Encapsulation
        - data and method bundle together
        - hide data to restrict access

27. Class
    - create:
        - new: use to create object
        - .newInstance(): `xxx.class.getDeclaredConstructor().newInstance();`
            - when dont know the class until runtime
            - dependency injection framework like Spring, Hibernate

    - static vs instance method
        - static call without instance 
        - static no access to `this` 
            - keyword to reference a variable that refer to current object
        - static only can access static member of class, instance can access both
        - static method cannot be overriden ( as they resolved at compile time, not run time)

28. Object
    - entity with properties and method
    - instance of class

29. constructor
    - special method to initialize object
    - if dont provide, compiler will auto generate default constructor
    - copy constructor
        - Constructor that initialize object as a copy of existing object
        - parameter is object
    - private constructor
        - mark with private
        - for singleton, class that only provide static method

30. Interface
    - collection of static final variable and abstract method
    - can implement multiple Interface

31. Marker Interface
    - empty interface without method or field

32. Abstract class vs Interface
    - Abstract class:
        - can have abstract or non abstract method
        - can have final method
        - only single extend
        - can have protected, private etc
        - inherit everything except constructor
    - Interface
        - only abstract method
        - not allow final method
        - allow multiple implement
        - public by default
    - for IS-A ( inheritance)

33. aggregation
    - 1 class contains reference to another class as a field
    - life cycle of the container object is independent of container object
        - contained object can exist even container destroyed
    - unidirectional ( 1 way relation)
    - part of association

34. composition
    - 1 class container ref to another class as field
    - but the inner class is own by container, and create inside using new()
        - child cannot exist without parent
    - loose coupling compared to inheritance

35. association
    - HAS-A

36. throw-clause
    - checked exception ( subclass of `Exception` )
        - check by compiler if we handle it
        - for recoverable situation
        - if method throw `IOException`, `SQLException`, need to have throw outside
        ```Java
        void readFile() throws IOException {// need to throw here
            throw new IOException("File not found");
        }
        ```
    - unchecked exception ( subclass of `RuntimeException` or `Error`)
        - compiler wont check if we handle
        - unrecoverable 
        - dont need declare throw
        - `NullPointerException`, `ArithmeticException`, `ArrayIndexOutOfBoundsException`

37. virtual function
    - java method that not final, static, private are virtual
    - method that can be override


38. serialization
    - need to `implements Serializable` to be able to serialize
    - if subclass dont want, need to override method `writeObject` by throwing `NotSerializableException`

39. Collection Framework
    - classes 
        - ArrayList
            - dynamic resize
            - fast random access
            - slow insert delete
            - synchronize
                - synchronizedList
                ```Java
                List<String> list = Collections.synchronizedList(new ArrayList<>());
                synchronized (list) { // need synchronized for safe iteration
                    for (String s : list) {
                        System.out.println(s);
                    }
                }
                ```
                - CopyOnWriteArrayList
                ```Java
                CopyOnWriteArrayList<String> list = new CopyOnWriteArrayList<>();
                for (String s : list) { // dont need manual synchronize
                    System.out.println(s);
                }
                ```
        - Vector
            - thread-safe ver of ArrayList
            - slower than ArrayList
            - for backward compatibility of old Java Ver
            - all method lock, slow, **recommend to just synchronize ArrayList**
        - LinkedList
            - doubly linked list
            - fast insert delete
            - slow random access
        - PriorityQueue
            - order by natural order (asc) or custom comparator
            - no null
            - not thread safe
        - TreeSet
            - sort in Asc by default
            - no dups
            - implement using red-black tree
        - ArrayDeque
            - for stack
            - `Deque<Integer> stack = new ArrayDeque<>();`
        - HashSet
            - unique value only
        - LinkedHashSet
            - ordered version of HashSet
    - Interface 
        - Collection
            - base interface, extended by List, Set, Queue
            - add(), remove(), size(), contains(), iterator()
        - Set (unordered, no dups)
            - implemented by HashSet, LinkedHashSet, TreeSet
            - same method as collection
        - List (ordered, allow dups)
            - implemented by ArrayList, Vector, LinkedList
            - get(int index), set(int index, E element), indexOf()
        - Queue (FIFO)
            - implement by PriorityQueue, LinkedList
            - offer() //insert , poll()//get head with remove , peek() // get head without remove
        - Deque
            - implement by ArrayDeque, LinkedList
            - offerFirst(), offerLast(), pollFirst(), pollLast()
        - Map
            - implement by HashMap, TreeMap, LinkedHashMap
            - put(), get(), containsKey(), keySet()

40. Type casting
    - Polymorphism downcasting
        - when object type declare as parent
        - can cast it to child to use child method 
    - numeric type conversion
        - int, double etc
    - generics

41. Generics:
    - class, interface, method with type parameter to work with different type while ensure compile time safety
    - why need?
        - catch error at compile time ( not run time)
        - no need manual casting
        - code reuse for many types
    - without generic
        ```Java
        List list = new ArrayList();
        list.add("Hello");
        String s = (String) list.get(0);  // ✅ Need manual casting
        // we can list.add(123), which can cause ClassCastException at runtime
        ```
    - with generic
        - class
        ```Java
        class Box<T> {        // T is a type parameter
            private T value;
            public void set(T value) { this.value = value; }
            public T get() { return value; }
        }

        Box<Integer> intBox = new Box<>();
        intBox.set(10);
        Integer val = intBox.get();  // ✅ No casting
        ```

        - method
        ```Java
        public <T> void print(T data) {
            System.out.println(data);
        }

        print("Hello");   // T = String
        print(123);       // T = Integer
        ```

    - Convention
        - T: type
        - E: Element
        ```Java
        class MyList<E> {
            private List<E> list = new ArrayList<>();
            public void add(E element) { list.add(element); }
            public E get(int index) { return list.get(index); }
        }
        ```
        - K: key
        - V: value
        ```Java
        class MyMap<K, V> {
            private Map<K, V> map = new HashMap<>();
            public void put(K key, V value) { map.put(key, value); }
            public V get(K key) { return map.get(key); }
        }
        ```
        - ?: wildcard, use when dont know and dont care

42. why cannot create generic array
    - array have runtime check, but generic Erasure remove type
    - causing error
    - but we can have generic arrayList
    ```Java
    List<String> list = new ArrayList<>(); //ok, cause dont check at runtime
    List<String>[] arr = new List<String>[10]; //nok, when erasure it become List[] only
    ```

43. in memory storage
    - arrays
        - store in contiguous memory
        - easier access by calculate address using base + offset
    - ArrayLists
        - store in contiguous
        - when created, default 10-16 elements ( depends on java ver)
        - resizing: when reached capacity, create larger array and move old to new

44. convert between Array & ArrayList
    - to ArrayList
        - `Arrays.asList(the_array)` fast, but cannot change arrayList size!!
        - `List<String> list = new ArrayList<>(Arrays.asList(the_array));`
        - 
        ```Java
        List<String> list = new ArrayList<>();
        Collections.addAll(list, arr);
        ```
    - to array
        - `Object[] arr = list.toArray();`

45. FailFast vs FailSafe
    - HashMap: map modified while iterating, throw `ConcurrentModificationException`
    - ConcurrentHashMap: map modified no impact to iteration. it iterate over a snapshot

46. Thread
    - extends Thread, with run method
    ```java
    class MyThread extends Thread {
        public void run(){ 
            ///do something
        }
    }
    ```
    - advantage of multithreading
        - responsiveness
        - resource sharing
        - economy
        - scalability
        - better communication
    - 2 way to create thread
        - extends `Thread` class, override `run()` method
        - implements `Runnable` interface, override `run()` method
    - thread has its own program counter, execution stack, and local variables
        - but share same memory space with other threads in same process
    - lifecycle of thread ( use `.getState()`)
        - New: created but not started
        - Runnable: running
        - blocked: temporarily suspended, waiting for a resource or event
            - by `suspend()`
            - use `resume()` to unlock
        - waiting: wait for another thread to perform task, without timeout
            - Sleep()
            - Wait(): wait thread until other thread signal it to wake up
                - need `notify()` or `notifyAll()`
            - Join(): wait for thread to finish, auto resume
            - waiting for I/O operation
            - Synchronization issue
        - timed_waiting: wait for specific time
        - terminated: finish execution
    - main thread
        - parent thread of all other thread
        - auto create when program start
    - daemon thread
        - for background operation that need to perform continuously
            - garbage collector
        - lower priority than user thread
            - only execute when no user required background task
    - how java achieve multi-threading
        - through time-sharing/ time-slicing
        - CPU switching between active threads
        - os in charge of allocating CPU time to each thread
    - thread priority
        - MIN_PRIORITY
        - MAX_PRIORITY
        - NORM_PRIORITY

47. Garbage collector
    - to avoid memory leak
    - cannot avoid GC in java
    - free up unused memory in application periodically
    - Heap memory
        - object created with `new`
        - if no longer reachable, GC will free up
    - Stack memory
        - local variable, cleared when method exit
        - not manage by GC
    - draw back
        - may cause pause in application when clearing memory
        - clearing process not deterministic, unpredictable
        - may still temporarily increase memory usage if program create and discard many short-lived object
    - type of garbage collector
        - Minor
        - Major
        - Full