The file that has your SQLite database will usually look like this.

```
name_of_database.sqlite
```

Next to there will be two files:
```
name_of_database.sqlite-shm
name_of_database.sqlite-wal
```

# Don't delete these!

name_of_database.sqlite-wal:

The `.sqlite-wal` contains data that has not been commited to the main
database.

The .sqlite-shm is a shared memory file

_______________________________________________________________________________
### Tip 2
Never try to edit the .sqlite file directly!

Always use something like `litecli` to edit your database.

_______________________________________________________________________________
### Tip 3

Do not try to make a manual file backup of your database by simply copying
the file. Use one of the built-in commands in litecli like `vacuum into`

_______________________________________________________________________________
