SQLite only supports one concurrent writer

So only person can update the database at a time.

Multiple people can read from the SQLite database at once.

Locking states:
1. Unlocked
2. Shared Lock
3. Reservef Lock
4. Pending
5. Exclusive Lock
