NOTE: I am using `pgcli` and not `psql`

### To display an overview of all the tables in your database.
```sql
\d
```
_______________________________________________________________________________

### To get an overview of a specific table:

```sql
\d name_of_table
```

_______________________________________________________________________________
### How to view only specific columns in a table

This will display only data from the `id` and `first_name` columns 
of the table `player_info` and it will only show the first 10 rows.
```sql
select id, first_name from player_info limit 10;
```

_______________________________________________________________________________
