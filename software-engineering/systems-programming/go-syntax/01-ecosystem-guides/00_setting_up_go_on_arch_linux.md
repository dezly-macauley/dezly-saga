# Setting up Go on Arch Linux

## Step 1: Install the compiler and language server

```
sudo pacman -S go 

```

To check which version of the Go compiler you have, run this command:

```
go version
```

Create your GOPATH and add the path to your .zshrc file

Run these commands

```
cd ~

mkdir .go
```

Then add this line to the end of your .zshrc file
```
export GOPATH="$HOME/.go/"
```

---

Install the language server for for Go  

```
sudo pacman -S gopls

```

To make sure that it is installed run the command:

```
which gopls
```

---
