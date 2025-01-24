# Flexible Types

You can create a table without any data types.
```sql
create table ex(a,b,c);
```

Insert data into the table
```sql
insert into ex values(1, 'asdf', 3);
insert into ex values(1, 'asdf', 'asdf');
insert into ex values('asdf', 2, 'asdf');
```

view all rows and columns in the table
```sql
select * from ex;
```

The output:
```sql
+------+------+------+
| a    | b    | c    |
+------+------+------+
| 1    | asdf | 3    |
| 1    | asdf | asdf |
| asdf | 2    | asdf |
+------+------+------+
```
_______________________________________________________________________________
### What happens when you insert data with the wrong variable type

```sql
create table ex2(a integer, b text);
insert into ex2 values(162, 'Tsunade');
-- The value 'Dezly' in the second row do not match the column type of integer
insert into ex2 values('Dezly', 'Macauley');
```

```sql
select * from ex2;
```

The output:
```sql
+-------+---------+
| a     | b       |
+-------+---------+
| 162   | Tsunade |
| Dezly | Macauley|
+-------+---------+
```

SQLite will store the data, despite the column type not matching.

This is because by default, SQLite applies the data type of a value, at
the cell level.

So in row 2, `Dezly` is stored as `text`,
even though column `a` is meant for storing integers.

_______________________________________________________________________________
You can confirm this with the following:

```sql
select typeof(a) from ex2 where a = 'Dezly';
```

Dezly is stored as text
```sql
+-----------+
| typeof(a) |
+-----------+
| text      |
+-----------+
```

The way this works is due to something called `type affinity`

1. Column a is for integers, so `integer` is the type affinity.
2. When you put a value into column a, SQLite will try to store that value
as an integer.
3. If it can't, it will store that value as another type to avoid data loss.

_______________________________________________________________________________
