# Nonce
_______________________________________________________________________________
Every transaction has a `nonce` field.

This nonce field is a number that is specific to the account 
that is making the current transaction.
_______________________________________________________________________________

#### What is that value of the number in the `nonce` field?

1. This value is the number of successful transactions that 
the account making the transaction has already made.

2. This number does NOT include the current transaction because it is being
processed. The nonce is not about how many transactions that account has made,
but instead, how many successful transactions that account has made in the
past.

3. So if you are using a brand new account and you make a transaction.
The nonce will be 0 because that account has never made any successful
transactions and the current transaction is still pending.
_______________________________________________________________________________
