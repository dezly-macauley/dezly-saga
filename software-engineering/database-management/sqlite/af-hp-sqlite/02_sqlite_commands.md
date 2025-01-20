### How to open a database

```
litecli --liteclirc $litecli_config_path database_file.sqlite
```

### How to exit litecli
```
exit
```

### How to view all the tables in a database
```
.table;
```

### How to select all fields from a table
```
select * from table_name;
```
_______________________________________________________________________________

### How to create a table

SQLite uses a dynamic typing system, meaning it does not strictly enforce the data types of a column. Instead, it uses type affinities, which are general types associated with the column definitions.

#### SQLite Type Affinities and Variable Types

1. **VARCHAR(255)** → **TEXT** affinity
2. **INT** → **INTEGER** affinity
3. **CHAR, CLOB, TEXT** → **TEXT** affinity
4. **BLOB, [not specified]** → **BLOB** affinity
5. **REAL, FLOAT, DOUBLE** → **REAL** affinity
6. **Anything else** → **NUMERIC** affinity


## Explanation of SQLite Type Affinities
SQLite uses a dynamic typing system and type affinities, not strict data types.

### 1. TEXT Affinity
- For columns with type declarations containing `CHAR`, `CLOB`, `TEXT`, or unspecified.
- Examples: `VARCHAR`, `TEXT`

### 2. INTEGER Affinity
- For columns with type declarations containing `INT`.
- Examples: `INTEGER`, `BIGINT`

### 3. REAL Affinity
- For columns with type declarations containing `REAL`, `FLOA`, `DOUB`.
- Examples: `REAL`, `DOUBLE`

### 4. BLOB Affinity
- For columns with type declarations containing `BLOB` or no type specified.
- Examples: `BLOB`, `none`

### 5. NUMERIC Affinity
- For columns with other type declarations like `NUM`, `DEC`, `FIX`, `MON`, `VAR`.
- Examples: `NUMERIC`, `DECIMAL`

## SQLite Variable Types
SQLite recognizes and stores the following basic data types:

1. **NULL**: Represents a null value.
2. **INTEGER**: Represents a signed integer.
3. **REAL**: Represents a floating-point number.
4. **TEXT**: Represents a text string.
5. **BLOB**: Represents a binary large object, stored exactly as input.

_______________________________________________________________________________

NOTE: Remember to add to table modifier at the end to eforce strict types

```
CREATE TABLE fight_records (
    player_name TEXT,
    total_wins INTEGER,
    total_losses INTEGER
) strict;
```

_______________________________________________________________________________
### How to delete a table

```
drop table if exists name_of_table;
```

_______________________________________________________________________________
### How to insert values into a table

```
INSERT INTO fight_records (player_name, total_wins, total_losses) 
VALUES ("Dezly", 20, 9);
```
_______________________________________________________________________________
Get today's date:

```
select datetime("now");
```
year-month-day hour(24 hour format):minutes:seconds

E.g.
2025-01-16 13:20:51

_______________________________________________________________________________
### Boolean

Boolean values are stored as `INTEGER` in your database because SQLite does
not have a boolean type.

true = 1
false = 0

But you can query tables using true and false.  

```
select * from name_of_table where bool = true;
```

Create the table:
```
CREATE premium_status (
    player_name TEXT,
    has_premium INTEGER,
) strict;
```

Insert into table:
```
INSERT INTO premium_status (player_name, has_premium) 
VALUES ("Dezly", 20, 9);
```

_______________________________________________________________________________

FTS - Full text search

_______________________________________________________________________________
