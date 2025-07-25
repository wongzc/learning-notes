# python_study_note

python important learning points

1. class: use when need attribute+method, center of OOP
    - `encapsulation`: data/attribute + method
    - `inheritance`, cleaner, lesser repeat code
        - MRO
            - if multiple class which all have same final ancestor
            - and all have super().xx() to use ancestor method
            - then will resolve following MRO
            ```python
            class A:
                def greet(self):
                    print("A")

            class E(A):
                def greet(self):
                    print("E")
                    super().greet()

            class B(E):
                def greet(self):
                    print("B")
                    super().greet()

            class C(A):
                def greet(self):
                    print("C")
                    super().greet()

            class D(B,C):
                def greet(self):
                    print("D")
                    super().greet()

            d = D()
            d.greet()

            print(D.mro())
            # will be D B E C A, it wait if other same level will go to that ancestor
            ```
    - `polymorphism`: different class, but if have same method name, can polymorphism call. treat different class as same instance of same class
        ```python
        class Dog:
            def speak(self):
                return "Woof!"

        class Cat:
            def speak(self):
                return "Meow!"

        animals = [Dog(), Cat()]

        for animal in animals:
            print(animal.speak())  # Same method name, different behavior
        ```
    - `abstract`: create a template
        defined ( with pass only). it is also data abstraction
        ```python
        from abc import ABC, abstractmethod

        class template(ABC): 
            @abstractmethod
            def method(self): 
                pass
        ```
    - can multi inherit, just `class A(inher1, inher2)`

2. python is compiled to bytecode, and run by python-virtual-machine (PVM, interpreter), while c++ is compiled to machine code, run by CPU

3. mutable datatype can be changed after create, list/dict/set
    - immutable can only be reassigned, memory address changed

4. argument pass by object-reference
    - mutable data: passed by reference, i.e., can be updated by function
        - (list, dict, set, bytearray)
    - immutable data: passed by value ( int, float, str, tuple, bool, frozenset)

5. list comprehension:
    - `[a for a in range(10)]`
    - `[a if a==b else aa for a in range(10)]`
    - `[<expression_if_true> if <condition> else <expression_if_false> for <item> in <iterable> if <condition>]`

6. lambda:
    - `y=lambda x:x*k `#k can be variable form global, x is the argument
    - `y=lambda a,b=0:a*b` #can set default value
    - `y=lambda a,b=None: a*(b if b is not None else 0)` # use if else

7. exception handling:
    - try: xxx 
    - except Exception as e: yyy ( or just except: yyy)
    - finally: zzz

8. string can use .swapcase() to change AbCDe to aBcdE

9. for vs while:
    - use for when have start and stop condition
    - use while when only have stop

10. pass func as argument: 
    ```python
    def a(text): 
        return text.upper()
    def b(func): 
        return func('haha')
    ```
    - b is taking func as input, and use func to process the text 'haha'

11. use of asterisk *
    - unpack:
        ```python
        def add(a, b, c):
            print(a + b + c)

        nums = [1, 2, 3]
        add(*nums) 

        d={'a':3,'b':4,'c':8}
        add(**d) 
        ```  

    - merge list
        - `c=[*a,*b]`
    - function argument:
        - will get list of args and dict of kwargs
        ```python
        def func(*args, **kwargs): 
            print(args,kwargs)
        func(1,'a',kk=11,ll='asa')
        #will be (1,'a') {'kk'=11,ll='asa'}
        ```
    - variable length:
        - a,*b=[1,2,3,4,5,6,7,8]
            - a=1
            - b=[2,3,4,5,6,7,8]

12. decorators:
    - make change to function input/output, add logging
        - existing decorator
            - `@staticmethod`: for method that not going to access state of class, can skip self. i.e.,` def func(a)` instead of `def func(self,a)`
            - `@classmethod`: 
                - to access class level variable
                ```python
                class Student:
                    total_students = 0

                    def __init__(self, name):
                        self.name = name
                        Student.total_students += 1

                    @classmethod
                    def get_total_students(cls):
                        return cls.total_students

                s1 = Student("Alice")
                s2 = Student("Bob")

                print(Student.get_total_students())  # 2

                ```
                - alternative constructor
                ```python
                class Person:
                    def __init__(self, name, age):
                        self.name = name
                        self.age = age

                    @classmethod
                    def from_birth_year(cls, name, birth_year):
                        return cls(name, 2025 - birth_year)

                p1 = Person("Alice", 30)  
                p2 = Person.from_birth_year("Bob", 1995)

                print(p1.name, p1.age)  # Alice 30
                print(p2.name, p2.age)  # Bob 30
                ```
            - `@property`: change method to attribute, so can call without ()
            - `@abstractmethod`
            - `@lrucache`
                - store input output mapping, return output is input is cached
                - fix sized cache, evicted based on LRU
    - can access control as well (** jwt wrapper?)
    - use for caching: create a cache={}, if arg in cache, return cache[arg], else store cache[arg]=result before return back
    - timing
    - add metadata
        ```python
        def add_metadata(author=None, version=None):
            def decorator(func):
                func.author = author
                func.version = version
                return func
            return decorator

        @add_metadata(author="Zhao Cai", version="1.0")
        def my_function():
            return "Hello"

        print(my_function())            # Hello
        print(my_function.author)       # Zhao Cai
        print(my_function.version)      # 1.0
        ```
    - Enforcing Functionality
        ``` python
        def my_decorator(func):
            def wrapper(*args, **kwargs):
                print("Before function call")
                result = func(*args, **kwargs)
                print("After function call")
                return result
            return wrapper
        

        @my_decorator
        def say_hello(name):
            print(f"Hello, {name}!")
            return f'nice to meet you, {name}'

        print(say_hello("Alice"))
        ```
    - decorator style:
        ``` python
        outer(x,y,z): # to pass param into decorator, like metadata, repeat func etc
            decorator(func): # take in func as param
                wrapper(*args, **kwargs): # if set metadata, can skip this and just func.xx=yy, then return func
                    # no return, do wrapper thing
                return wrapper
            return decorator
        ```
13. backslash: \
    - use for line continuation
    - use for escape char like
        - \\ -> \
        - \n -> newline
        - \' -> '

14. forward slash: /
    - dir and path
    - note, path can be;
        - path = "C:\\Users\\Name\\file.txt"
            or path = r"C:\Users\Name\file.txt"
        - path = "C:/Users/Name/file.txt"

15. string prefix
    - f: format string with {var}
    - r: raw string
    - b: byte string
    - u: unicode string

16. docstring
    - define with 'or" or triple ' or ", for class & function/method & module
    ```python
    def add(a, b):
        """
        xxx
        """
        return a + b
    ```
    - call with
    ```python
    print(add.__doc__)
    help(add)
    ```

17. __init__.py:
    - to mark dir as package instead of namespace package
        - namespace package: `import mypackage.module1`
        - normal: can use above or `from mypackage import ...`
    - package init code
        ```python
        # in mypackage/__ini__.py
        print("mypackage is being imported")
        ```
    - control what get import from normal package `from mypackage import *`
        ```python
        # mypackage/__init__.py
        from .module1 import func1
        from .module2 import func2

        __all__ = ["func1", "func2"] # only these 2 imported with *
        ```
    - put versioning data inside or shorten import path
    ```python
    __version__ = "1.0.0"
    from .utils.helper import useful_func # then other package can "from mypackage import useful_func"
    ```

18. dynamically typed, but can use type hint
    ```python
    def greet(name: str, age: int) -> str:
    ```

19. floor/int:
    - math.floor(): toward small, 3.7->3, -3.7->-4
    - int(): keep only int part

***20. import file/ path and relative

21. deep copy:
    - copy.deepcopy for list/set/dict
    - list[:] for list (but list in list will still be shallow! take note)

22. python use Tim Sort: merge + insertion sort
    * other sorting method

23. python debug: python -m pdb
    - Set Breakpoint: b <function_name> (e.g., b factorial)
    - Continue Execution: c
    - Step Through Code: s
    - Inspect Variables: p <variable_name> (e.g., p n)
    - Exit Debugger: q

24. generator:
    - (x for x in y)
    - ```python
        def func(y): 
            for i in range(y): 
                yield i
      ```
    - zip(x,y)
    - map lambda to list
        names = ['Alice', 'Bob', 'Charlie']
        n=map(lambda x: x*2, names)

25. iterator: all generator are iterator
    - any thing with next is iterator

26. memory management: heap to manage, garbage collector, recycle unused memory and free it out

27. monkey patching: dynamically modify class at run time, like change the class.func to other func

28. __init__ is constructor, to initialize value
    
29. pickling & unpickling: change data to byte stream, for network transfer
    - import pickle
    - with open('abc.pickle','wb') as file: pickle.dump(data,file)
    - with open('abc.pickle','rb') as file: data=pickle.load(file)

30. access specifier:
    - attribute/method in class start with __: cant access outside
    - start with _: proected, protected concept is just infromational only

31. unit test
    import unittest
    from math_operations import divide

    class TestMathOperations(unittest.TestCase):
        def test_divide(self):
            self.assertEqual(divide(9, 3), 3)
            # Testing for division by zero
            with self.assertRaises(ValueError):
                divide(10, 0)

    if __name__ == '__main__':
        unittest.main()

32. GIL restricted python multi threading, means 1 python process, only 1 thread at a time
    to overcome:
        - multiprocessing: using multiple python process to run
        - multithreading: use threading module, can thread to other task when 1 is waiting I/O related response
        - asyncio: similar with threading, but more suit for faster and more frequent switch, control with async, await

33. python switch
    - match term: case 'a':'aa' case 'b':'bb'

34. Walrus Operator:
    - assign value within expression
    - if (n:=len(list))>3:
    - while (user_input := input('abc'))!='haha'
    - lengths = [n for word in words if (n := len(word)) > 3]
        lengths = [len(word) for word in words if len(word) > 3]

35. when import a file in python, code of imported file will be always executed, unless wrap it with if __name__=="__main__"

36. when say import A.py into B.py, if A.py reading a file using relative path, may causing problem if B.py is in differnt folder ( as run from B will use B as currntdir)
    better to use full path to avoid such problem

37. to import a .py in parent dir. just use: sys.path.append(parent), then can import the filename
    ```python
    currentfile = os.path.realpath(__file__)
    current=os.path.dirname(currentfile)
    parent=os.path.dirname(current)
    sys.path.append(parent)
    ```
38. nonlocal & global
    ```python
    def outer():
        x = "outer value"

        def inner():
            nonlocal x  # Refers to x in the nearest enclosing (non-global) scope
            # for nested function to call outer function variable
            x = "modified by inner"
            print("Inner:", x)

        inner()
        print("Outer after inner:", x)
    ```

    ```python
    x = "global value"

    def change_global():
        global x  # to call var in global, not function
        x = "modified globally"
        print("Inside function:", x)

    change_global()
    print("Outside function:", x)

    ```
