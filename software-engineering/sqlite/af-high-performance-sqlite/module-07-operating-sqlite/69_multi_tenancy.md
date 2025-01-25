Create a database per tenant

So when someone logs and selects from their tenant database, 
you are garanteed that all the data belongs to them because they are the only
tenant in that database.

So you will have as many databases as you have users.
_______________________________________________________________________________
I have created two databases:

`practice_db.sqlite` and `central_db.sqlite`
_______________________________________________________________________________
First I choose which database is your main database.

In this example, I want to use `practice_db.sqlite` as my main database.

_______________________________________________________________________________
Once you are logged in, connect to the second database:

```sql
attach database './central_db.sqlite' as central_db;
```
_______________________________________________________________________________

Then run this command:
```
.databases;
```

It will show the databases that you are connected to.
_______________________________________________________________________________
How to connect to a second database

```
attach database './central_db.sqlite' as central_db;
```

To create a table in the `central_db` database while connected to
`practice_db` then do this:

```sql
CREATE TABLE central_db.tenants(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    database TEXT,
    subdomain TEXT
);
```

Insert values into the table:
```sql
insert into central_db.tenants values(
    1, 'a8f4be.sqlite', 'aaron'
);
```

To check that the table was created:
```sql
select * from central_db.tenants;
```

The output will look like this:
```sql
+----+---------------+-----------+
| id | database      | subdomain |
+----+---------------+-----------+
| 1  | a8f4be.sqlite | aaron     |
+----+---------------+-----------+
```
_______________________________________________________________________________
