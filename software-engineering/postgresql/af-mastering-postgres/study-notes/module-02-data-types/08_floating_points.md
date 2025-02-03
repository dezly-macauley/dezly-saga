Floating points are like numerics but they can be accessed much faster.

The trade of is that floating points tend to be less precise if the number
contains lots of decimals.

Use floats for when you are dealing with non-financial data and you want the

_______________________________________________________________________________
### Types of floating points in PostgreSQL

1. real or float4: 4 bytes
- 1E-37 - 1E37 (6 digits after the decimal point)

2. double precision or float8: 8 bytes
- 1E-307 - 1E+308 (15 digits after the decimal point)


_______________________________________________________________________________
