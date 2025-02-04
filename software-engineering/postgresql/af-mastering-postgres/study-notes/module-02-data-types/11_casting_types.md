To convert one type into another do this:

Let's say that there is a value which is a text value.

E.g. '512' this is the number '512' as a string as you can see from the 
single qoutes.

To convert it to an integer, do this:

### This is the syntax in PostgreSQL
'512'::integer

### The generic way to do this in SQL is this
cast('512' as money)

_______________________________________________________________________________
## When would you use this?

E.g. E.g. you have read the contents of a json file and stored the entire
json object inside as a string. Then you would store the string as a jsonb
in your table.
_______________________________________________________________________________
