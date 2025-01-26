```sql
create table user_accounts(
    name text,
    balance real
) strict;
```

```sql
insert into user_accounts values
    ('Seth', 26.3),
    ('Cindy', -10.52),
    ('Liam', -15.78)
;
```

```sql
select * from user_accounts;
```

The output:
```sql

```

### How to update multiple balances at once

```sql
update user_accounts
set balance = case
    when name = 'Cindy' then 250
    when name = 'Liam' then 700
    -- If the name is not 'Cindy' or 'Liam' then the balance will
    -- not be changed.
    else balance
end;
```
