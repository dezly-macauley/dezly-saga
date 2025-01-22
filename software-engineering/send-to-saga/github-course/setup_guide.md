### Install Git
Refer to DezlyOS

___

Note: 
Whenever you see something like this "[your-username]"

Enter it without the the square brackets

E.g. 
git config --global user.name "[your-username]"

means

git config --global user.name "dezly-macauley"

___

git config --global user.name "[your-username]"
git config --global user.email "[your-email-address]"

___

.gitignore
___

git init

git status

git clone [the-url-of-the-git-repository]

This will download all of the files, branches, and commits

___

### Git Branches

#### How to view available branches
git branch

#### How to create a new branch
git branch [branch-name]

#### How to delete a new branch
git branch -d [branch-name]

#### How to switch to a branch
git checkout [branch-name]

#### How to create a new branch and switch to it
git checkout -b feature/my-new-feature

___

### Making changes

#### Add all files to stagging area
git add .

git commit -m "[your-message-explaining-what-you-changed]"

___

### How to undo a `git add` commmand 

Unstage all files
git restore --stage .


___

### Cloning

git clone https://github.com/dezly-macauley/meta-vc-practice.git

___

### How to view the remote of a git directory

The remote is the url that the directory is being push up to, or that
you pull down changes frome

git remote -v

___

### How to push to github

It is good practice to pull down the latest changes from the repo first.

git pull

git push -u origin [name of branch]

___

### Workflows

Feature branching

___

### How to get the commit hash of the branch

cat .git/refs/heads/main

meta-vc-practice on  main
❯ cat .git/refs/heads/main
49a1751051c7fea1b5f5b0b9ea99049f1a247257

___

### To check what branch we are currently on

meta-vc-practice on  main
❯ cat .git/HEAD

___

git log with nicer output

❯ git log --pretty=onelinee

___

