SQLite has 5 type affinities

1. integer - For numbers without any decimal points
2. real - For numbers with decimal points
3. text - For text content

4. null - You don't need to set this. 
This is what SQLite uses for fields that are blank.

5. blob - binary large object.
This is used to store files like images, audio, videos, and other binary
data directly in the database.
