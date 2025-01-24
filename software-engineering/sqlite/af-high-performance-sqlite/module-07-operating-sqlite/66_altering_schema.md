How to rename a column in a table 

```sql
create table weapons(name text, count integer) strict;
insert into weapons values('Katana', 7) ;
select * from weapons;
```

To rename `name` to `weapon_name`
```sql
alter table weapons rename name to weapon_name;
```
