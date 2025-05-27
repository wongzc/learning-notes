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

7. create array
```Java
```