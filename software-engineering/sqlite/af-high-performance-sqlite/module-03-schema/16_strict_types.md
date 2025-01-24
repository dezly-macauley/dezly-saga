# Strict Types
_______________________________________________________________________________

How to create a table with strict types.

Remember to add `strict` at the end.
```sql
create table player_wins(wins integer, player_name text) strict;
insert into player_wins values(20, 'Sasuke');
select * from player_wins;
```

If you try to enter value into a column that is not the correct variable type,
SQLite will reject the value.

```sql
insert into player_wins values('Vegeta', 9000);
select * from player_wins;
```

cannot store TEXT value in INTEGER column player_wins.wins

_______________________________________________________________________________
