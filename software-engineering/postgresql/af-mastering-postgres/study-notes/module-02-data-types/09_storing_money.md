### Method 1 - Store the money in integer

So you would use the lowest common denomination

E.g. If you are trying to store 215.78 US dollars

- The lowest common denomination of a dollars is cents
- 1 dollar = 100 cents

So to add 215.78 dollars to column in a table, convert the amount to cents

215.78 dollars * 100 = 21578 

Next store 21578 in your table as an integer
E.g. us_dollar_price_in_cents

_______________________________________________________________________________
When you need to retrieve the data then you would convert it back to dollars:

21578 cents / 100 = $215.78 

215 us dollars, and 78 cents

_______________________________________________________________________________

This requires some conversion work but the benefit is that integers as faster
to access.
_______________________________________________________________________________
### Method 2 - Store the money as a numeric

This keeps the precision of the decimals 
and you don't have do any conversions.

The downside to this is that numerics are the slowest number type to access.

_______________________________________________________________________________
## Storing Multiple currencies

Create a separate column that stores the value next to the currency.

E.g. Currency codes have 3 letters like 'USD', 'EUR', 'JPY'
currency CHAR(3)

_______________________________________________________________________________
