# python_study_note

python important learning points

1. class: use when need attribute+method, center of OOP
    - encapsulatiopn: data/attribute + method
    - inheritance, cleaner, lesser repeat code
    - polymorphism: diffeernt class, but if have same method name, can polymorphism call. treat differetn class as same instance of same class
    - abstract: create a template ( use abc.ABC & abc.abstarctmethod) with essentiate method defined ( with pass only). it is also data abstraction
        ```
        class template(ABC): 
            @abstractmethod
            def method(self): pass
        ```
    - can multi inherit, just class A(inher1, inher2)

2. python is compiled to byttecode, and run by python-virtual-machine (PVM, interpreter), while c++ is compiled to machine code, run by CPU

3. mutable datatype can be changed after create, list/dict/set
    - immutable can only be reassigned, memory address changed

4. argument pass by object-reference
    - muttable data: passed by reference, i.e., can be updated by function
    - immutable data: passed by value
        - for function to call var in outer fucntion (nested function): use nonlocal
        - for fucntion to call var in global (non-fucntion): use global

5. list comprehension:
    - `[a for a in range(10)]`
    - `[a if a==b else aa for a in range(10)]`
    - `[<expression_if_true> if <condition> else <expression_if_false> for <item> in <iterable> if <condition>]`

6. lambda:
    - `y=lambda x:x*k `#k can be vairable form global, x is the argument
    - `y=lambda a,b=0:a*b` #can set default value
    - `y=lambda a,b=None: a*(b if b is not None else 0)` # use if else

7. exception haddling:
    - try: xxx 
    - excpet Exception as e: yyy ( or just except: yyy)
    - finally: zzz

8. string can use .swapcase() to change AbCDe to aBcdE

9. use for when have start and stop condition, use while when only have stop

10. pass func as argument: 
    ```
    def a(text): 
        return text.upper()
    def b(func): 
        return func('haha')
    ```
    - b is taking func as input, and use func to process the text 'haha'

11. use of asterisk *
    - unpack:
        - a=[1,2,3], b=(4,5,6)
            - def add(x,y,z): return x+y+z
            - add(*a)
            - add(*b)
        -d={'x':3,'y':4,'z':8}
            - add(**d)

    - merge list
        - c=[*a,*b]
    - function argument:
        - def func(*args, **kwargs): print(args,kwargs)
            - will get list of args and dict of kwargs
    - variable length:
        - a,*b=[1,2,3,4,5,6,7,8]
            - a=1
            - b=[2,3,4,5,6,7,8]

12. decorators:
    - make change to function input/output, add logging
        - existing decorator
            - @staticmethod: for method that not going to access state of class, can skip self. i.e., def func(a) instead of def func(self,a)
            - @classmethod: to access class variable, that is not like self.xxx. usually direct under a class.
                - we create a method, @classmethod; def func(cls): return cls.xxx
                - we can access this value in init by the classname.variablename
            - @property: change method to attribute, so can call without ()
            - @abstractmethod
            - @lrucache
    - can access control as well (** jwt wrapper?)
    - use for caching: create a cache={}, if arg in cache, return cache[arg], else store cache[arg]=result before return back
    - timing
    - add metadata, like func.var=''
    - Enforcing Functionality
        ``` 
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
13. backslash: \
    - use for line continuation
    - use for escape char like
        - \\ -> \
        - \n -> newline
        - \' -> '

14. foward slash:
    - divie and path
    - note, path can eb;
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
    - use help(function/class/module ** no need () for func)
    - or just function/class/module.__doc__

17. __init__.py:
    - can use to import function from others .py (from .other_file import func1) ** need to have the "."
    - the folder that contain init+ other can be a module
    - we can import it to other python file. from module import func1/ import module.func1

18. dynamically typed, but can use type hint

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
    - ```
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

26. memory management: heap to manage, garabage collector, recyle unused memeroy and freee it out

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
    ```
    currentfile = os.path.realpath(__file__)
    current=os.path.dirname(currentfile)
    parent=os.path.dirname(current)
    sys.path.append(parent)
    ```
