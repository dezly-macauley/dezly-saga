SQLite does not have a `boolean` variable type, so you have to store boolean
values as integers.

- The value `true` will be stored as `1`
- The value `false` will be stored as `0`

A simple table:
```sql
create table character_info(
    name text,
    has_powers integer
);
```

Some rows of data:
```sql
insert into character_info values
    ('Batman', false),
    ('Superman', true),
    ('Flash', true),
    ('Robin', false);
```
_______________________________________________________________________________

The output will look like this:
```sql
+----------+------------+
| name     | has_powers |
+----------+------------+
| Batman   | 0          |
| Superman | 1          |
| Flash    | 1          |
| Robin    | 0          |
+----------+------------+
```
So as you can see, SQLite it kind of weird when it comes to booleans.

1. SQLite understands what `true` and `false` are but it will store them
as integers.
_______________________________________________________________________________

```sql
select * from character_info where has_powers = true;
```

The output
```sql
+----------+------------+
| name     | has_powers |
+----------+------------+
| Superman | 1          |
| Flash    | 1          |
+----------+------------+
```
_______________________________________________________________________________

