1. GIL=global interpreter lock  in CPython
    - process wide mutex, only allow 1 thread excute python bytecode

2. why?
    - Cpython memory management use reference counting
        - updating of reference count is shared mutable operation
        - multiple thread can update same value and cause race condition
    - not atomic operation, 2 thread may update same time and corrupt memory
    - use GIL to restrict
        - avoid fine grained lock
        - keep interpreter simple and fast

3. how?
    - thread need to acquire GIL to run
    - after run, release GIL
    - cooperative preemption:
        - thread run 5ms and switch ( default 5ms)
        - control by `sys.setswitchinterval(0.005)`
    - blocking I/O
        - when thread blocking I/O, CPython release GIL, other thread run

4. other
    - use multi process to bypass GIL, for true parallell