The boolean type in PostgreSQL is one byte.

```sql
create table player_status(
    name text,
    is_online  boolean
);
```

```sql
insert into player_status (name, is_online) values
('Dave', true),
('Jin', false)
;
```
