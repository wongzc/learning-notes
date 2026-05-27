1. What
    - Implementation of Python
        - implementation: the actual program that understand pyton and run it
        - GIL is implementation details of CPython 
    - official Python interpreter ( written in C)

2. How
    - Cpython parse python code to bytecode
    - Cpython VM then run the bytecode

3. Other
    - other implementation
        - PyPy:
            - faster
            - direct translate code to machine code
            - still has GIL