### What does json stand for?
_______________________________________________________________________________
JSON stands for JavaScript Object Notation, 
which is a lightweight data-interchange format 
that's easy to read and write for humans and machines.

1. json (JavaScript Object Notation)
- Data is stored as text
- This variable type is smaller than jsonb
- white spaces are retained
- Parsing happens at query time, which can be slower for complex operations.
- No indexing support

2. jsonb (JavaScript Object Notation Binary)
- A variant of json
- Data is stored in binary format
- This variable type is larger than json
- white spaces are stripped for efficiency
- Parsing happens at storage time, which makes queries faster.
- Supports indexing, allowing efficient lookups and searches.

- Duplicate keys are not allowed. 
If you try to store a JSON object with duplicate keys, 
jsonb will only keep the last occurrence of the key.
_______________________________________________________________________________
### How to create a table that stores data in `jsonb` format

```sql
create table my_jsonb_table(
    -- `not null` simply means that each row
    -- must contain data in these columns
    file_name text not null,
    file_data jsonb not null
);
```
_______________________________________________________________________________
### How to store jsonb data into a column

E.g. I have a json object in this format:
```json
{
    "user_handle": "dezly_macauley",
    "daily_post_count": 225,
    "yearly_spending": 52751.23,
    "is_premium_user": true,
    "iphone_version": null,
    "favorite_tags": ["programming", "pizza"],
    "contact_info": {
        "email": "dezly@example.com",
        "phone": "+1234567890",
        "address": {
            "street": "123 Tech St",
            "city": "Techville",
            "zip": "12345"
        }
}
```

```sql
insert into my_jsonb_table (file_name, file_data) values
    (),
;
```

_______________________________________________________________________________
