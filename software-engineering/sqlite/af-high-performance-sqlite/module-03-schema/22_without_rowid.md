You can create a table that does not have a rowid

1. This is a performance optimization because there is one column that is the
id. So lookups of data will be faster.
2. Make sure that you create a column that is the primary key.
3. Add the table modifier `without rowid` 
4. If your primary key is an integer then you don't need to add the table
modifier `without rowid`, because your primary id will become the rowid.
5. If your primary key is not an integer then this will be a performance
optimization because it will get rid of the rowid.

```sql
create table trading_cards(
    card_id text primary key,
    name text
) strict, without rowid;
```

```sql
-- Here I am being explicit about which columns I want to insert data into
insert into trading_cards (card_id, name) values
    ("33fex", "green dragon"),
    ("dsf27", "red dragon")
;
```
