Create a database per tenant

So when someone logs and selects from their tenant database, 
you are garanteed that all the data belongs to them because they are the only
tenant in that database.

So you will have as many databases as you have users.
_______________________________________________________________________________
I have created two databases:

`practice_db.sqlite` and `central_db.sqlite`

First use litecli to login to `practice_db.sqlite`

Then run this command:
```
.databases;
```

It will show that it is the `main` database I am connected to.
_______________________________________________________________________________

Then run this command:

_______________________________________________________________________________
