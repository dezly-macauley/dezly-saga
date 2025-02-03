### What does json stand for?
_______________________________________________________________________________
JSON stands for JavaScript Object Notation, 
which is a lightweight data-interchange format 
that's easy to read and write for humans and machines.

1. json (JavaScript Object Notation)
- Data is stored as text
- This variable type is smaller than jsonb
- white spaces are retained
- Parsing happens at query time, which can be slower for complex operations.
- No indexing support

2. jsonb (JavaScript Object Notation Binary)
- A variant of json
- Data is stored in binary format
- This variable type is larger than json
- white spaces are stripped for efficiency
- Parsing happens at storage time, which makes queries faster.
- Supports indexing, allowing efficient lookups and searches.
_______________________________________________________________________________
### How to create a table that stores data in `jsonb` format

```sql

```
_______________________________________________________________________________
