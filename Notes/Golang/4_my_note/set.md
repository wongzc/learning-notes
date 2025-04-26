no built in set type in go, use map instead
- create empty: `set := make(map[int]struct{})`
- create with value: `set:= map[int]struct{}{1:{}, 2{},3:{}}`
- add to: `set[4]=struct{}{}`
- remove: `delete(set, 2)`
- check if in : `if _,exist :=set[4];exist {...}`