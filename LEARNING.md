## Learning path

1. In-memory store
   - Store databases, tables, and rows only in Go structs/maps/slices.
   - Learn: structs, pointers, maps, slices, method receivers, copying vs sharing data.
   - Goal: understand the shape of a database without worrying about files yet.

2. JSON file-backed store
   - Save table data to normal JSON files and load it back on startup.
   - Learn: `os.ReadFile`, `os.WriteFile`, `os.MkdirAll`, `encoding/json`, file paths.
   - Goal: understand persistence: data survives after the program exits.

3. Binary file-backed row store
   - Replace JSON with your own binary format for rows.
   - Learn: bytes, fixed-size vs variable-size fields, encoding integers/strings, offsets.
   - Goal: understand that databases usually store compact binary data, not JSON.

4. Page-based storage
   - Store rows inside fixed-size pages, for example 4KB blocks.
   - Learn: pages, page headers, free space, row slots, page IDs.
   - Goal: understand how databases avoid treating the whole table as one giant file.

5. B-tree indexes
   - Add a structure that helps find rows without scanning the whole table.
   - Learn: sorted keys, tree nodes, leaf/internal pages, search/insert/split.
   - Goal: understand why indexed queries are faster than full table scans.

6. Transactions + WAL
   - Add a write-ahead log before changing table files.
   - Learn: atomicity, durability, commit/rollback, crash recovery.
   - Goal: understand how databases avoid corruption when a crash happens mid-write.

7. Buffer pool
   - Cache recently used pages in memory instead of reading from disk every time.
   - Learn: page cache, dirty pages, flushing, eviction policies like LRU.
   - Goal: understand how databases manage memory vs disk access.

8. Concurrency control
   - Allow multiple operations/users safely at the same time.
   - Learn: locks, isolation levels, read/write conflicts, MVCC basics.
   - Goal: understand how databases prevent users from corrupting each other's data.

## Current project focus

You are currently around step 2: a JSON file-backed store.

The immediate things to learn/research are:

- The difference between a parser and an executor.
- The difference between a statement and an expression.
- How to represent SQL commands as AST structs.
- How to keep executor state, such as the current selected database.
- How `CREATE DATABASE` differs from `CREATE TABLE`.
- How rows should be represented before you build real binary storage.

## Useful search terms

- `Go os.MkdirAll example`
- `Go json Marshal Unmarshal struct`
- `database storage engine architecture`
- `SQL parser AST statements expressions`
- `database row store vs column store`
- `write ahead log explained`
- `database buffer pool explained`
- `B-tree database index explained`
