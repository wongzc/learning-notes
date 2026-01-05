1. type
```TypeScript

// let: variable can be reassigned
// const: cannot reassigned, but content cant be mutate
let a: number = 1;
let b: string = "hi";
let c: boolean = true;

//array
let xs: number[] = [1, 2, 3];
let ys: Array<string> = ["a", "b"];

let t: [number, string] = [200, "OK"];

//dict (Record)
const dict: Record<string, number> = {
  apple: 1,
  banana: 2,
};
dict["orange"] = 3;
const n = dict["apple"]; // number

//dict (Map), use when many insert/delete, need iteration order, key not string
const m = new Map<string, number>();

m.set("apple", 1);
m.set("banana", 2);

console.log(m.get("apple"));      // number | undefined
console.log(m.has("banana"));     // boolean
m.delete("banana");

//set
const s = new Set<number>();

s.add(1);
s.add(2);
s.add(2);           // no effect (unique)
console.log(s.has(2)); // true
s.delete(1);

for (const x of s) console.log(x);

```

2. any, unknown, never
```TypeScript
//any: disable type checking
let vAny: any = JSON.parse("{}");           // avoid if possible


//unknown: need to use if or typeof ==="string" to check
let vUnknown: unknown = JSON.parse("{}");   // safer
if (typeof vUnknown === "object" && vUnknown !== null) {
  // still needs narrowing for properties
}

function fail(msg: string): never {
  throw new Error(msg);
}

```

3. Union, intersection
```TypeScript
type Id = string | number;

type HasId = { id: string };
type HasName = { name: string };
type User = HasId & HasName; // must have both id and name

const user: User = {
  id: "u123",
  name: "Alice",
};
```

4. const
```TypeScript
// as const: prevent mutation through variable, but mutate through underlying is allow
const ENV = "prod" as const;        // type is "prod" not string
const ROLES = ["admin", "user"] as const; // readonly ["admin","user"]
type Role = typeof ROLES[number];   // "admin" | "user"

```

5. interface & type
```TypeScript
//interface:
//  - allow declaration merging
interface Person {
  name: string;
  age?: number; // optional
}
// declaration merging, interface define twice
interface User {
  id: string;
}

interface User {
  name: string;
}

const u: User = { id: "1", name: "Alice" };

//extension
interface User2 extends User {
    name2: string
}

// type
//  - support union, intersections, primitive, tuple, mapped type
type Person2 = {
  name: string;
  age?: number;
};
type Id = string | number;
type Pair = [number, number];
type Result<T> =
  | { ok: true; value: T }
  | { ok: false; error: string };

//extension
type Id2= Id & {id2: string};

```

6. index signature
```TypeScript
type StringMap = Record<string, string>;
const headers: StringMap = { "x-trace-id": "abc" };

type Counter = Record<string, number>;
const counts: Counter = {};
counts["k"] = (counts["k"] ?? 0) + 1;
```

7. optional chaining
```TypeScript
const city = user.profile?.address?.city ?? "unknown";

```

8. function types
```TypeScript
function add(a: number, b: number): number {
  return a + b;
}

const mul = (a: number, b: number): number => a * b;
```

9. default param
```TypeScript
function greet(name = "world") {
  return `hello ${name}`;
}

function sum(...nums: number[]) {
  return nums.reduce((acc, n) => acc + n, 0);
}
```

10. generics
```TypeScript
function identity<T>(x: T): T {
  return x;
}

function first<T>(arr: T[]): T | undefined {
  return arr[0];
}
```

11. constraint generics
```TypeScript
function getId<T extends { id: string }>(x: T) {
  return x.id;
}
```
12. narrowing ( to avoid bug)
```TypeScript
type ApiErr = { error: string };
type ApiOk = { data: unknown };

function isApiErr(x: any): x is ApiErr {
  return x && typeof x === "object" && typeof x.error === "string";
}

function handle(res: ApiErr | ApiOk) {
  if ("error" in res) return res.error;
  return res.data;
}

```

13. discriminated unions
```TypeScript
type Result =
  | { ok: true; value: number }
  | { ok: false; error: string };

function f(r: Result) {
  if (!r.ok) return r.error;
  return r.value;
}
```

14. async/ promise
```TypeScript
async function fetchJson<T>(url: string): Promise<T> {
  const r = await fetch(url);
  if (!r.ok) throw new Error(`HTTP ${r.status}`);
  return (await r.json()) as T;
}
```

15. classes
```TypeScript
class Queue<T> {
  private items: T[] = [];
  push(x: T) { this.items.push(x); }
  pop(): T | undefined { return this.items.shift(); }
}
```

16. Common TS util
```TypeScript
type User = { id: string; name: string; age?: number };

type UserId = Pick<User, "id">;           // { id: string }
type UserNoAge = Omit<User, "age">;       // { id, name }
type UserReq = Required<User>;            // makes age required
type UserRO = Readonly<User>;             // all fields readonly
type UserPartial = Partial<User>;         // all optional
```

17. props typing
```TypeScript
type ButtonProps = {
  label: string;
  onClick: () => void;
  disabled?: boolean;
};

export function Button({ label, onClick, disabled }: ButtonProps) {
  return <button disabled={disabled} onClick={onClick}>{label}</button>;
}
```

18. state typing
```TypeScript
const [count, setCount] = useState<number>(0);
const [user, setUser] = useState<User | null>(null);

```

19. event typing
```TypeScript
function onInput(e: React.ChangeEvent<HTMLInputElement>) {
  console.log(e.target.value);
}
```

20. for loop
```TypeScript
// for range
for (let i = 0; i < 5; i++) {
  console.log(i);
}

// for ...of
const nums = [10, 20, 30];
for (const n of nums) {
  // n: number
  console.log(n);
}

//object entries
const obj = { a: 1, b: 2 };
for (const [key, value] of Object.entries(obj)) {
  // key: string
  // value: number
}

```

21. length
```TypeScript

//array
const arr = [1, 2, 3];
arr.length;

//set
const s = new Set<number>([1, 2, 3]);
s.size;

//map
const m = new Map<string, number>([
  ["a", 1],
  ["b", 2],
]);
m.size;

//object
const obj = { a: 1, b: 2 };
Object.keys(obj).length; // 2

```

```TypeScript
```

```TypeScript
```