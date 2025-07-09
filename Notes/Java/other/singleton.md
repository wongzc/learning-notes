# Java Singleton Example

## 1. Eager Initialization
```Java
public class Singleton {
    private static final Singleton instance = new Singleton(); 
    // private so that not accessible outside
    // static so that static method can access
    // final so cannot be changed

    private Singleton() {} // private constructor, so cannot call "new" from outside

    public static Singleton getInstance() {
        return instance;
    }
}
```
✅ Thread safe
❌ not lazy, instance created when class initialize ( even instance unused)
    - when call `getInstance()`
    - when use `Class.forName("Singleton")`
    - other class init the class as part of its own static init

Usage
```Java
Singleton single = Singleton.getInstance();
single.otheMethod();
```

## 2. Lazy Initialization
```Java
public class Singleton {
    private static Singleton instance; // not final, so that can be changed by getInstance

    private Singleton() {}

    public static Singleton getInstance() {
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
}
```
✅ lazy loaded: only create instance when getInstance is called
❌ not thread-safe, can create multiple instance (if 2 thread entered `instance==null`)


## 3. Synchronized Method
```Java
public class Singleton {
    private static Singleton instance;

    private Singleton() {}

    public static synchronized Singleton getInstance() {
        // improve lazy when synchronized, so only 1 thread can use getInstance at a time
        // avoid situation where multiple enter `instance==null`
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
}
```
✅ thread-safe, lazy
❌ slow, as method-level locking


## 4. Double-Checked Locking
```Java
public class Singleton {
    private static volatile Singleton instance;
    // volatile here to ensure other thread see instance!=null when it updated
    // because we need it to be seem outside of synchronized!

    private Singleton() {}

    public static Singleton getInstance() {
        if (instance == null) {
            // thread only enter if instance==null ( multi thread can enter)
            synchronized (Singleton.class) {
                // move synchornized to inisde of method, only 1 thread can at a time
                if (instance == null) {
                    // check again, incase another thread enter after instance created
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```
✅ thread-safe, lazy, efficient
❌ complex code, need to use `volatile` ( bug for pre-java 5)


## 5.Static Inner Class
```Java
public class Singleton {
    private Singleton() {}

    private static class Holder {
        private static final Singleton instance = new Singleton();
        // ensure thread safe by static inner class hold the instance
    }

    public static Singleton getInstance() {
        return Holder.instance;
        // ensure lazy by only init Holder when getInstance called
    }
}
```
✅ thread-safe, lazy, efficient, modern
❌ not suitable for dependency injection


## 6. Enum Singleton
```Java
public enum Singleton {
    INSTANCE;

    public void doSomething() {
        // ...
    }
}
```
✅ thread-safe, lazy, efficient
❌ cant extend class, cant pass argument