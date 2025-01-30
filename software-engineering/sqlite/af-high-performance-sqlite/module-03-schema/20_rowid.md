The row ID is a secret primary key for your table that SQLite keeps track of

You should not rely on it

```sql
create table rowid_example(
    user_id text primary key, 
    cool_number integer
);
```

You can insert data into multiple rows but only in specific columns by doing
the following:
```sql
-- I am inserting data into only multiple rows,
-- but I am only filling in the `cool_number` column
insert into rowid_example (cool_number) values
(28), (217), (23);
```

```sql
select rowid, * from rowid_example;
```

The output will be:
```sql
+---------+---------+-------------+
| user_id | user_id | cool_number |
+---------+---------+-------------+
| 1       | 1       | 28          |
| 2       | 2       | 217         |
| 3       | 3       | 23          |
+---------+---------+-------------+
```
