B trees are the underderlying data structure that makes indexing affective.

This is the way that the data is arranged when you create an index

If you don't have an index and you search for something in a table, 
SQLite will scan every row in the database until it finds what 
you are looking for.

When you use an index, SQLite puts the names in order and then put that
ordered list of names into a datastructure that is easily traversible.

This is more efficient because SQLite does not have to scan every single 
row to find what you are looking at.

_______________________________________________________________________________
I'm going to start with this table:

```sql
create table ninja_emails(
    -- `primary key` means that the column called `id` will 
    -- be used to identify each row in the table.
    --- `autoincrement` means that if a row is inserted without an `id`,
    -- SQLite will generate an id number that is unique for that row.
    id integer primary key autoincrement,
    -- I don't want the name field to be empty
    name text not null,
    -- The email field should not be empty and every email must be unique
    email TEXT not null unique 

    -- strict means that SQLite will not allow a row to be added 
    -- if the variable type of the value does not match the variable type of a
    -- column
) strict;
```

How to insert multiple rows of data into the tables:
```sql
INSERT INTO ninja_emails (name, email) 
VALUES 
    ('Naruto Uzumaki', 'naruto.uzumaki@konoha.com'),
    ('Sasuke Uchiha', 'sasuke.uchiha@konoha.com'),
    ('Sakura Haruno', 'sakura.haruno@konoha.com'),
    ('Kakashi Hatake', 'kakashi.hatake@konoha.com'),
    ('Hinata Hyuga', 'hinata.hyuga@konoha.com'),
    ('Shikamaru Nara', 'shikamaru.nara@konoha.com'),
    ('Temari', 'temari@suna.com'),
    ('Neji Hyuga', 'neji.hyuga@konoha.com'),
    ('Gaara', 'gaara@suna.com'),
    ('Tenten', 'tenten@konoha.com');
```

View the table
```sql
select * from ninja_emails;
```
The output
```sql
+----+----------------+---------------------------+
| id | name           | email                     |
+----+----------------+---------------------------+
| 1  | Naruto Uzumaki | naruto.uzumaki@konoha.com |
| 2  | Sasuke Uchiha  | sasuke.uchiha@konoha.com  |
| 3  | Sakura Haruno  | sakura.haruno@konoha.com  |
| 4  | Kakashi Hatake | kakashi.hatake@konoha.com |
| 5  | Hinata Hyuga   | hinata.hyuga@konoha.com   |
| 6  | Shikamaru Nara | shikamaru.nara@konoha.com |
| 7  | Temari         | temari@suna.com           |
| 8  | Neji Hyuga     | neji.hyuga@konoha.com     |
| 9  | Gaara          | gaara@suna.com            |
| 10 | Tenten         | tenten@konoha.com         |
+----+----------------+---------------------------+

```
_______________________________________________________________________________
### How to create an index on the name column




_______________________________________________________________________________

```sql
select * from ninja_emails where name = 'Gaara';
```
_______________________________________________________________________________
