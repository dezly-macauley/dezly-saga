# Dezly Macauley 
### Smart Contract Hacking Process
_______________________________________________________________________________

#### Step 0: Get the tools needed to perform the audit 

_______________________________________________________________________________

#### Step 1: Get the onboarding documentation from the client 

The client must fill in an onboarding documentation and send it back to me.
This document must contain the repo url and the exact commit hash I need to
work on.

This is important because I don't want to be working on something that the
client has already fixed.
_______________________________________________________________________________

#### Step 2: Download the repo and switch to the correct commit hash

First get the url of correct repo:
```
https://github.com/Cyfrin/3-passwordstore-audit
```

Next get the commit hash:

```
7d55682ddc4301a7b13ae9413095feffd9924566
```

Download the clients code base and enter it:

```
git clone https://github.com/Cyfrin/3-passwordstore-audit
cd 3-passwordstore-audit
```

Next switch to the commit hash:
```
git checkout 7d55682ddc4301a7b13ae9413095feffd9924566
```

To confirm that you are viewing the code for this commit hash run this
command:

```
git rev-parse HEAD
```

You can also see what branch of the repo this commit hash is part of
by running this command:

```
git branch --contains 7d55682ddc4301a7b13ae9413095feffd9924566
```

You will get an output like this:

```
3-passwordstore-audit on  HEAD (7d55682)
❯ git branch --contains 7d55682ddc4301a7b13ae9413095feffd9924566

* (HEAD detached at 7d55682)
  main
  onboarded
```

Now it's import to note that this commit hash is "detached head".
What this basically means in simple terms is that it is read-only.

So you can edit the files but your changes and notes won't be saved.

_______________________________________________________________________________

#### Step 3: Create your own branch 

Make sure that you are still in the directory with the commit hash

Then run this

```
git switch -c dezly-security-audit
```

This will do the following:
1. Create a new branch in the repo called `dezly-security-audit`
2. It will switch to this branch
3. The first commit of this branch will be the commit hash:
7d55682ddc4301a7b13ae9413095feffd9924566
because this is the agreed starting point of your audit

4. This branch and commit are not a "detached head", so you can actually
save your changes. And push the repo to github as a private repo

To confim that you are on the correct commit hash:
You can use this command again: 
`git rev-parse HEAD` and `git branch`

```
3-passwordstore-audit on  dezly-security-audit
❯ git rev-parse HEAD
7d55682ddc4301a7b13ae9413095feffd9924566
```

```
3-passwordstore-audit on  dezly-security-audit
❯ git branch
* dezly-security-audit
  main
  onboarded
```
_______________________________________________________________________________

#### Step 4: Create a private github repo so that you can save your work

Update the remote url of the private github repo 
you created on Github for the audit.

```
git remote set-url origin https://github.com/dezly-macauley/3-passwordstore-audit.git
```

Then push you changes to the security audit branch
```
git push -u origin dezly-security-audit
```

Then confirm that you are still on the correct commit hash:


```
3-passwordstore-audit on  dezly-security-audit
❯ git rev-parse HEAD

7d55682ddc4301a7b13ae9413095feffd9924566

```

_______________________________________________________________________________

### NOTE: Whenever you push to this repo, these are the steps:

```
git add .
```

```
git commit -m "A message explaining what you changed"
```

```
git push -u origin dezly-security-audit
```

_______________________________________________________________________________

#### Step 5: Find out how many lines of source code are in the scope

Based on the scope, the client only want to have the PasswordStore.sol
contract in the src directory audited.

```
In scope vs out of scope contracts

./src/
└── PasswordStore.sol

```

This is how I get the lines of code:
```
cd 3-passwordstore-audit
cloc ./src/
```

Example of the output:

```
3-passwordstore-audit on  dezly-security-audit 
❯ cloc ./src
       1 text file.
       1 unique file.                              
       0 files ignored.

github.com/AlDanial/cloc v 2.00  T=0.07 s (13.4 files/s, 549.4 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Solidity                         1              6             15             20
-------------------------------------------------------------------------------
```

So this means that for the onboarding form, the `nSLOC` (number of source
of lines of code) is 20

_______________________________________________________________________________

#### Step 5: Figure out the complexity of the code base:

I will be using `bun` instead of npm because it's faster

```
bun init
```

Install the Solidity Metrics tool from Consensys
```
bun add solidity-code-metrics@0.0.27
```

To use it:

```
bunx solidity-code-metrics ./src/PasswordStore.sol > solidity_metrics.md
```

you can also get an html output:

```
bunx solidity-code-metrics ./src/PasswordStore.sol --html > solidity_metrics.html
```

_______________________________________________________________________________

### Read README.md to figure out what the smart contract is supposed to do

E.g. From the 3-passwordstore-audit
```
A smart contract application for storing a password. 
Users should be able to store a password and then retrieve it later. 
Others should not be able to access the password.
```

From this start thinking of attack vectors (potential problems) 

E.g. 
- Can I prevent the user password from being retrieved later?
- Can I prevent the users password from being stored?
- How exactly is their password stored? Is it encrypted on both ends and in
transit?
- Can I access the password?

_______________________________________________________________________________

### Run Solidity Metrics

Look at the html or markdown output.

Go to the section that says `Source Units in Scope` also go to the 
`contract Summary` section to get a good look at the functions that are in
scope.

_______________________________________________________________________________
