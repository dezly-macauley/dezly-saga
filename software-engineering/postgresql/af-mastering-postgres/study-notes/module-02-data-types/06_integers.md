### How to delete a table from your database

```sql
drop table if exists ninja_info;
```
_______________________________________________________________________________
### The main integer types in PostgreSQL

1. smallint or int2 (2 bytes):
- range: -32,768 to 32,767
- Use this for smaller numbers like an age column. 

2. integer or int4 (4 bytes):
- Range: -2,147,483,648 to 2,147,483,647
- Stick to this type for most cases

3. bigint or int8 (8 bytes):
- Range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
- Probably a good idea to make an id column this type

_______________________________________________________________________________

How to create the table:
```sql
create table player_info(
    player_name text,
    age smallint
);
```

How to check that the table was created:
```sql
\d
```

The output:
```
+--------+-------------+-------+----------------+
| Schema | Name        | Type  | Owner          |
|--------+-------------+-------+----------------|
| public | player_info | table | dezly_macauley |
+--------+-------------+-------+----------------+
```

How to view information about the `player_info` table

```sql
\d player_info
```

The output:
```
+-------------+----------+-----------+
| Column      | Type     | Modifiers |
|-------------+----------+-----------|
| player_name | text     |           |
| age         | smallint |           |
+-------------+----------+-----------+
```
_______________________________________________________________________________
How to insert data into the table:
```sql
insert into player_info (player_name, age) values
('Alice', 25),
('Bob', 42),
('Charlie', 18),
('Diana', 60);
```

NOTE: 
- Always add `text` values in single qoutes.
- When inserting values remember to add the `;` after the last insert.

_______________________________________________________________________________
