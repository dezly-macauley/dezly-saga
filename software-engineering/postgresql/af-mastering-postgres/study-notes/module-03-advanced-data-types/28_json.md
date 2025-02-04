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
_______________________________________________________________________________
### How to convert a `text` value to `jasonb` value in PostgreSQL

Use this `::jsonb`  

This is called type casting. 
The syntax is:

'the contents of jsonb object in single qoutes'::jsonb


```sql
insert into my_jsonb_table (file_name, file_data) values
    ('dezly_macauley.json', '{"user_handle":"dezly_macauley","daily_post_count":225,"yearly_spending":52751.23,"is_premium_user":true,"iphone_version":null,"favorite_tags":["programming","pizza"],"contact_info":{"email":"dezly@example.com","phone":"+1234567890","address":{"street":"123 Tech St","city":"Techville","zip":"12345"}}}'::jsonb)
;
```
_______________________________________________________________________________
### How to run shell commands while logged into pgcli

Just use `\!` and the command that you want to execute

I can use the program `lsd`
```sql
\! lsd
```

I can use the program`bat` to open the file with syntax highlighting:
```sql
\! bat json-examples/tsunade_senju.json
```

I can even edit the file in Neovim with the `nvim` command:
```sql
\! nvim json-examples/tsunade_senju.json
```
_______________________________________________________________________________
### How to to insert a `json` file into `PostgreSQL` 

- There is a directory in the root of this project called, `json-examples`
- There is a file inside called `tsunade_senju.json` that contains data.

_______________________________________________________________________________
#### Step 1 - Read the contents of the file and store it into a variable

(NOT WORKING)
The contents of the json file will be read and stored 
into a variable called `file_content` as text.

NOTE: Use the `cat` command (don't use bat).
This is because the cat command is a raw unformatted output.

```sql
\! file_content=$(cat json-examples/tsunade_senju.json)
```
_______________________________________________________________________________
#### Step 2 - Typecast the text variable into a `jsonb` and insert it

NOTE: To use a text variable in PostgreSQL the syntax is :'variable_name'

```sql
insert into my_jsonb_table (file_name, file_data) values
    ('tsunade_senju.json', :'file_content'::jsonb)
;
```

INSERT INTO my_jsonb_table (file_name, file_data)
VALUES
    ('tsunade_senju.json', '\! cat json-examples/tsunade_senju.json'::jsonb);

_______________________________________________________________________________
### How to view the values of each field in the json object

NOTE: There is a difference between `->` and `->>`
1. `->`JSON/JSONB Object Access Operator
- This will return the field in `json` format
- Use this to display nested objects and arrays

2. `->>` JSON/JSONB Text Access Operator
- This will return the field in `text` format
- Use this to display primative values like `string`, `number`, 
`boolean` and `null`

_______________________________________________________________________________
#### Viewing primative values

NOTE: `as user_handle` tells PostgreSQL 
what to disply ass the column name when returning the values.

```sql
select
    file_data->>'user_handle' as user_handle,
    file_data->>'daily_post_count' as daily_post_count,
    file_data->>'yearly_spending' as yearly_spending,
    file_data->>'is_premium_user' as is_premium_user
from my_jsonb_table 
where file_name = 'dezly_macauley.json';
```

The output looks like this:
```sql
+----------------+------------------+-----------------+-----------------+
| user_handle    | daily_post_count | yearly_spending | is_premium_user |
|----------------+------------------+-----------------+-----------------|
| dezly_macauley | 225              | 52751.23        | true            |
+----------------+------------------+-----------------+-----------------+
```
_______________________________________________________________________________
#### Viewing compoung types (Arrays)

```sql
select
    file_data->'favorite_tags' as favorite_tags
from my_jsonb_table 
where file_name = 'dezly_macauley.json';
```

The output:
```sql
+--------------------------+
| favorite_tags            |
|--------------------------|
| ["programming", "pizza"] |
+--------------------------+
```
_______________________________________________________________________________
To get a specfic value from the array.

NOTE: The arrows are important
- `->` means give me the array in the jason object as a data structure
- `->>` means display the value at that index in the array as text

```sql
select
    file_data->'favorite_tags'->>0 as first_tag,
    file_data->'favorite_tags'->>1 as second_tag 
from my_jsonb_table 
where file_name = 'dezly_macauley.json';
```

The output:
```sql
+-------------+------------+
| first_tag   | second_tag |
|-------------+------------|
| programming | pizza      |
+-------------+------------+
```
_______________________________________________________________________________
If you had a nested array like this:

```json
{
  "favorite_tags": [
    ["programming", "coding"], 
    ["pizza", "food"]
  ]
}
```

```sql
SELECT
    file_data->'favorite_tags'->0->>0 AS first_tag_in_array_one,
    file_data->'favorite_tags'->0->>1 AS second_tag_in_array_one,
    file_data->'favorite_tags'->1->>0 AS first_tag_in_array_two,
    file_data->'favorite_tags'->1->>1 AS second_tag_in_array_two
FROM my_jsonb_table 
WHERE file_name = 'dezly_macauley.json';
```

```sql
+------------------------+------------------------+-----------------------+-----------------------+
| first_tag_in_array_one | second_tag_in_array_one | first_tag_in_array_two | second_tag_in_array_two |
|------------------------|------------------------|------------------------|-------------------------|
| programming            | coding                 | pizza                  | food                    |
+------------------------+------------------------+------------------------+-------------------------+
```
_______________________________________________________________________________
#### Viewing compoung types (Nested json object)

```sql
select
    file_data->'contact_info'->>'email' as email,
    file_data->'contact_info'->>'phone' as phone,
    file_data->'contact_info'->'address'->>'street' as street,
    file_data->'contact_info'->'address'->>'city' as city,
    file_data->'contact_info'->'address'->>'zip' as zip
from my_jsonb_table 
where file_name = 'dezly_macauley.json';
```

```sql
+-------------------+-------------+-------------+-----------+-------+
| email             | phone       | street      | city      | zip   |
|-------------------+-------------+-------------+-----------+-------|
| dezly@example.com | +1234567890 | 123 Tech St | Techville | 12345 |
+-------------------+-------------+-------------+-----------+-------+
```
_______________________________________________________________________________
#### The JSON/JSONB Path Operator, and Path Extractor

1. `#>` Use this when you want to get a compound type back
- E.g. arrays, or a nested object.
2. `#>>` Use this when you want to get a primative type back.
- E.g. A number, string, boolean, or null

This is a better way to access json fileds

```sql
select
    file_data #>> '{contact_info, email}' as email,
    file_data #>> '{contact_info, phone}' as phone,
    file_data #>> '{contact_info, address, street}' as street,
    file_data #>> '{contact_info, address, city}' as city,
    file_data #>> '{contact_info, address, zip}' as zip
from my_jsonb_table
where file_name = 'dezly_macauley.json';
```

The output:
```sql
+---------------------+---------------+---------------+-------------+---------+
| email               | phone         | street        | city        | zip     |
|---------------------+---------------+---------------+-------------+---------|
| "dezly@example.com" | "+1234567890" | "123 Tech St" | "Techville" | "12345" |
+---------------------+---------------+---------------+-------------+---------+
```
_______________________________________________________________________________
