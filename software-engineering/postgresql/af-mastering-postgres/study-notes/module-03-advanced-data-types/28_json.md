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


_______________________________________________________________________________
