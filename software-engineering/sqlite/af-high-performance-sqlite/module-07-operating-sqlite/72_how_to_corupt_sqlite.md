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
