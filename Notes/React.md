1. useState
    - result in count=1, as React batches the state updates
    ```
    const [count, setCount] = useState(0);

    const handleClick = () => {
    setCount(count + 1); // This will use the initial value of count (0)
    setCount(count + 1); // This will still use the initial value of count (0)
    setCount(count + 1); // Again, this uses the initial value of count (0)
    };
    ```
    - result in count=3
    ```
    const [count, setCount] = useState(0);

    const handleClick = () => {
    setCount((prevCount) => prevCount + 1);
    setCount((prevCount) => prevCount + 1);
    setCount((prevCount) => prevCount + 1);
    };
    ```