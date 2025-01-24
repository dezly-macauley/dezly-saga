To check that you are in WAL mode:

```sql
pragma journal_mode;
```

If it isn't then you can set it like this:
```sql
pragma journal_mode = wal;
```

Write Ahead Log mode

```
wal
```

This is the best way to make SQLite performant.
