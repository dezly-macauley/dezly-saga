# Check Constraints

```sql
create table restaurant_sales(
    name 
        text 
        not null
        -- Think of the constraint as a custom error message
        -- if the check fails
        constraint name_be_20_characters_or_less
        check (length(name) <= 20)
    ,
    price 
        numeric
        not null
        constraint price_must_be_a_positive_number 
        check(price > 0)
);
```

```sql
insert into restaurant_sales(name, price) values
    ('pizza', 122.68),
    ('hot chocolate', 128.68)
;
```

PostgreSQL follows an `all or nothing` approach to data integrity
when it comes to multi-inserts. 

This means if one row fails a check, the entire insert fails.
Even if that insert contained rows that did meet the check.
```sql
insert into restaurant_sales(name, price) values
    ('apple pie', 522.17),
    ('dark chocolate', -128.68),
    ('pizza epic sauce extra dip', 122.68)
;
```
