### The structure of a PostgreSQL database

A PostgreSQL database, can contain multiple schemas.

And each of those schemas contain multiple tables.

So the structure is:

```
└── practice_db
    ├── schema_1
    │   ├── table_1
    │   └── table_2
    └── schema_2
        ├── table_1
        └── table_2
```

A schema is the structure of the table.
I.e. What variable types should be in each column.

Every PostgreSQL database starts off with a public schema

_______________________________________________________________________________
### How does this compare to other SQL databases?

_______________________________________________________________________________
#### SQLite - A single file contains the database

```
└── practice_db
    ├── table_1
    ├── table_2
    ├── table_3
    └── table_4
```
_______________________________________________________________________________
#### MySQL - single file contains the database

You have the MySQL server, which can have multiple databases

```
└── practice_db
    ├── database_1
    │   ├── table_1
    │   └── table_2
    └── database_2
        ├── table_1
        └── table_2
```

_______________________________________________________________________________
