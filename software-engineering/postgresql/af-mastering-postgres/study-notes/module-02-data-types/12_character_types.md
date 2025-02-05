Use `text` For storing string data. 

```sql
create table example_table(
    last_name text,
    -- `varchar` is short for variable-length character
    -- It means the currency symbol can only be three characters
    currency_symbol varchar(3),
);
```
