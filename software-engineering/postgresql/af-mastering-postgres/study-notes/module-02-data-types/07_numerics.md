# `numeric` 

This is for numbers that have decimals, 
and you want a high level of precision for decimals stored. 

```sql
create table financial_info(
    description text,

    -- NOTE: precision and scale

    -- The first number in the brackets is the precision
    -- The second number is the scale
    interest_rate numeric(5, 3)
);
```

```sql
insert into financial_info (description, interest_rate) values
('normal example', 18.621);
```

```sql
insert into financial_info (description, interest_rate) values
-- This will fail because the precision of the interest_rate column was set
-- to 5 digits and this number is actually 6 digits
('This should fail', 718.621);
```

```sql
insert into financial_info (description, interest_rate) values
-- The precision was set to 5, and the scale was set to 3 
-- This value will be stored because it is 5 digits 
-- However the decimals will be rounded to 3 decimals
('This will be rounded down', 5.6472);
```
This number will be stored as 5.647

_______________________________________________________________________________
```sql
insert into financial_info (description, interest_rate) values
-- The precision was set to 5, and the scale was set to 3 
-- This value will be stored because it is 5 digits 
-- However the decimals will be rounded to 3 decimals
('This will be rounded up', 5.6475);
```
This number will be stored as 5.648

_______________________________________________________________________________
### Precision and scale 
E.g. 18.621
- Precision is the total number of digits. So the precision is 5
- Scale is the number of digits after the decimal point. So the scale is 3

_______________________________________________________________________________
