# Mysql B+ Tree
### source: xiaolincoding
https://xiaolincoding.com/mysql

1. MySQL support multiple storage engine, default is innoDB
    - when read, read unit is page ( row by row too much I/O)
    - page is 16KB ( data page)
    - consist of:
        1. File Header (38 byte)
            - file information
            - point to next & previous data page
        2. Page Header (56 byte)
            - file status information
        3. Infimun + SuperMUM (26 byte) -> row record
            - 2 pseudo record, largest & smallest record in page
        4. User Record (variable size) -> row record
            - row data
            - use primary index to singly linked
                - singly linked: easy insert, delete, but search slow O(n)
                - so need to use page directory to locate!
        5. Free Space (variable size)
            - unused space
        6. Page Directory ( variable size)
            - save user relative space, for record index purpose
            - how directory created:
                1. group records and max & min into different group
                    - not inlcude marked 'deleted' record
                    - 1st group: 1 record
                    - last group: 1~8 record
                    - middle group: 4~8 record
                2. the last one in each group is the latest, it will store infor of record counts in group
                3. in page directory, it store the address of last record in each group, called 'slot'
        8. File Tailer (8 byte)
            - to check if page complete

2. B+ tree in innoDB
    - short & wide tree is better
    - each node in B+ tree is a data page
    <br>
    <img src="https://cdn.xiaolincoding.com//mysql/other/7c635d682bd3cdc421bb9eea33a5a413.png" width="500">
    - only leave node store data, branch node is just directory record
    - each node ( in the same layer) is double linked
    - use binary search to go layer by layer down to locate

3. primary vs secondary index
    - primary index: leaf node store data
        - clustered index
        - innonDB always create a clustered index when create table, index select by:
            1. primary index if it exist
            2. else, unique column that doesnt contain `NULL`
            3. else, create a hidden self increase id as index
    - secondary inex: leaf node store primary index
        - non-clustered index

4. why B+ Tee
    1. database is write to disk, disk is slow for read/write
        - memory: ns, disk: ms, 10^6 times!
        - smallest unit fo disk to I/O is `sector` (512B), os I/O `block`, which is multiple `sector`
        - when search, first read index from disk to memory, and use index to serach from disk

    2. binary search: from a sorted list, get middle, compare middle and decide to go left half or right half, repeat untill found ot no more.
    3. binary search tree: 
        - using a list fot BS, problem:
            - when insert new, need to sort again
            - when search, need to keep calculate  middle
        - use binary search tree, no need cal middle, just compare and decide go left or right
            - when insert also no need sort
        - problem with bs tree: if all value inserted are keep increasing, become a list, search time O(n)!!
    4. self balance tree:
        - to avoid BS tree become list
        - AVL tree: left & right height diff maximum 1, keep complexity as O(logn)
        - red-black tree
        - problem: more element, higher tree, I/O larger. the real problem: it is binary tree! only connect to 2 node!
    5. B tree
        - each node have M child ( M>2)
        - if M=3, each node have maximum 2 data and 3 child
            - if more, will split
            <br><img src="https://cdn.xiaolincoding.com//mysql/other/9a96956de3be0614f7ec2344741b4dcc.gif" width="500">
        - better than binary tree, shorter tree, less I/O
        - problem:
            - each node have data+index, data size maybe large and need more I/O to read useful index
            - when read to memory, data that doesnt need may load to memory together with the index.
            - if need range query, need inorder transversal, many I/O
                - inorder: visit left, visit root, then visit right
    6. B+ tree
        - only store data in leaf node ( index + record)
        - branch node: only index
        - all parent index in child index as well
        - all index in leaf node, connected with each other
        - index in branch node is biggest ( or smallest) for the child node
        <br><img src="https://cdn.xiaolincoding.com//mysql/other/b6678c667053a356f46fc5691d2f5878.png" width="500">
        - compare with B tree
            1. single query:
                - B tree faster, but fluctuation bigger as well, data may in branch or leaf
                - B+ tree in branch only index, can have more index and shorter & fatter, less I/O
            2. insert & delete:
                - B+ tree less change, as it has many redundant node
                <br><img src="https://cdn.xiaolincoding.com//mysql/other/23730b5af987480fabff0f1d142a2b6c.gif" width="500">
                - B tree big change
                <br><img src="https://cdn.xiaolincoding.com//mysql/other/7552002f9b8195ab650d431bfe66cce2.gif" width="500">
                - similar for insert, B+ tree may split if full, but only affect 1 route, no need complex algorithm
            3. range query:
                - B+ tree leaf node connected with link, can locate start condition then move
                - B tree must use transvesral to get, many I/O
                    - can use for nosql or mongoDB, which less range query
        - B+ tree in innoDB:
            - double linked leaf node

    