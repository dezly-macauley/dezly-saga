# How to create and backup a datbase
_______________________________________________________________________________
## Creating a datbase
Open litecli and then run this command:
```sql
.open name_of_new_database.sqlite;
```
This will create the file in same directory that the command was run from,
and then open it in litecli.

If the file exists, it will simply be opened.

_______________________________________________________________________________
## Creating a table

```sql
create table player_info(
    player_name text, 
    wins integer, 
    losses integer) 
strict;

-- Make sure to use single qoutes for text
insert into player_info values('Dezly', 40, 27);
```
_______________________________________________________________________________
### To view the list of tables in the database

```sql
.tables;
```
_______________________________________________________________________________
### To view the structure of the table 

```sql
pragma table_info(player_info);
```

The output will look like this:
```sql
+-----+-------------+---------+---------+------------+----+
| cid | name        | type    | notnull | dflt_value | pk |
+-----+-------------+---------+---------+------------+----+
| 0   | player_name | TEXT    | 0       | <null>     | 0  |
| 1   | wins        | INTEGER | 0       | <null>     | 0  |
| 2   | losses      | INTEGER | 0       | <null>     | 0  |
+-----+-------------+---------+---------+------------+----+
```
_______________________________________________________________________________
### To view all of the rows and columns inside a table

```
select * from player_info;
```

The output will look like this:
```sql
+-------------+------+--------+
| player_name | wins | losses |
+-------------+------+--------+
| Dezly       | 40   | 27     |
+-------------+------+--------+
```

_______________________________________________________________________________
### How to make an exact backup of the database

Open litecli and run this command:
```bash
vacuum into 'backup_practice_db.sqlite';
```

This will make a graceful backup. In simple terms a backup that is at a point
in time, and that avoids any data loss. 

Trying to manually copy a .sqlite file is risky as sqlite could be trying
to write data to a database will you are manually making a copy.

vacuum into is much safer because it is built-into sqlite.

_______________________________________________________________________________
