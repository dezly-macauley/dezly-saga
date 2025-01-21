# SQLite Guide
_______________________________________________________________________________

To check if you have SQLite installed
```
sqlite3 --version
```

To check if you have litecli installed
```
litecli --version
```

_______________________________________________________________________________

### How to create and save a new database:

A database is simply a file that contains a collection of tables.
Each of those tables has its own columns and rows of data.

First create a directory for your database
```
mkdir my-sqlite-databases
```

Enter the directory and create a database file.
```
cd my-sqlite-databases
touch simple_database.sqlite
```

Opent the database using `litecli`. 

NOTE: The default command line tool that SQLite has is called `sqlite3`

The reason I will be using litecli is because litecli has a way better 
user experience. It has autocomplete, syntax highlighting and your can
use `ctrl + l` to tidy up the terminal.
```
litecli simple_database.sqlite
```

simple_database.sqlite>

_______________________________________________________________________________

If you do this:
``` 
litecli
```

litecli will only create a database that exists in the shell memory.
As soon as you exit the prompt, the database will be lost.

to exit
```
exit
```
_______________________________________________________________________________
### How to create a table with specific data types

#### Sqlite data types: 
1. integer
2. real
3. text
4. blob
5. null

_______________________________________________________________________________
Create a table and set its data types
```
create table superheroes(
id integer,
hero_name text) strict;
```

The strict keyword is to ensure that the table is strictly typed.
I.e. You can't insert text data into a column that is declared 
for integer values.

How to delete a table:

```
drop table superheroes;
```

_______________________________________________________________________________
### To check what tables you have in your database:
```
.tables;
```
_______________________________________________________________________________
### How to get high level information about a table


To check how many columns a table has:
```
pragma table_info(superheroes);
```

You will get an output that looks like this:

+-----+-----------+---------+---------+------------+----+
| cid | name      | type    | notnull | dflt_value | pk |
+-----+-----------+---------+---------+------------+----+
| 0   | id        | INTEGER | 0       | <null>     | 0  |
| 1   | hero_name | TEXT    | 0       | <null>     | 0  |
+-----+-----------+---------+---------+------------+----+

_______________________________________________________________________________

You can also run `.schema superheroes;` to see how to re-create your table
using SQL commands

You will get an ouput like this:

+---------------------------+
| sql                       |
+---------------------------+
| CREATE TABLE superheroes( |
| id integer,               |
| hero_name text)           |
+---------------------------+

Just make sure to remember the `;` at the end if you ever
want to use this query.

_______________________________________________________________________________

To check how many rows of data a table has:
E.g. a table called users
```
select count(*) from superheroes;
```

If you want to count how many rows of data a table has that
contains columns do the following:

_______________________________________________________________________________
### How to insert values into a table

```
insert into superheroes values(67, 'Spiderman');
```

#### Warning for strings
Make sure to use single qoutes for strings and not double qoutes 
to avoid errors.

If you need to enter a value that contains an apostrophe, 
the correct way to enter it is like this:

`Mike's Kitchen` should be entered as `'Mike''s Kitchen' `
_______________________________________________________________________________
### How to view a table using `select` and `mode`

This will select all columns from the table called `superheroes`, 
but only the first 5 rows of data will be displayed.
```
select * from superheroes limit 5;
```
_______________________________________________________________________________
### Strict table

create table my_strict_table(
fav_num integer,
user_name text) strict;


_______________________________________________________________________________
### Working with dates

```
select date();
```
+------------+
| date()     |
+------------+
| 2024-12-16 |
+------------+

year-month-day

```
select datetime();
```
+---------------------+
| datetime()          |
+---------------------+
| 2024-12-16 10:23:35 |
+---------------------+

```
select time();
```
+----------+
| time()   |
+----------+
| 10:24:46 |
+----------+

_______________________________________________________________________________
### Working with booleans

In SQLite, booleans are not a separate data type. 
Instead, SQLite uses integers to represent boolean values:

0 is treated as false.
1 is treated as true.

```
CREATE TABLE active_players(
id integer,                 
user_name text,             
has_premium_account integer);
```

```
insert into active_players values (45, 'Jane', 1);
insert into active_players values (45, 'Patrick', 0);
```

Now when you make a query you can do this:

```
select * from active_players where has_premium_account = 1;
```

or this which is more natural.

```
select * from active_players where has_premium_account = true;
```
_______________________________________________________________________________

### Working with real numbers

SQLite gives you 15 digits of precision:

#### These are valid
123456789012345 | 0 decimals and 15 digits

0.123456789012345 | 15 digits after `.0`. The zero does not count in this case

100.567890123456 | (3 digits before the `.`, 12 digits after

If a value passes this, rounding starts to happen.

_______________________________________________________________________________

### Primary Key autoincrement

This will ensure that each row has a unique marker.

```
create table favourite_songs(
song_id integer primary key autoincrement,
song_name text
);
```

insert into  
_______________________________________________________________________________
