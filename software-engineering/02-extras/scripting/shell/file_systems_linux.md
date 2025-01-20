# The Python Ecosystem And Toolchain
_______________________________________________________________________________
## How to install Python on Arch Linux

```
sudo pacman -S --needed python
```

This will install the latest version of Python that pacman (the Arch package
manager has available)
_______________________________________________________________________________
## Where is Python installed on Linux systems?

On Linux systems like Arch Linux, the directory `/usr/bin` is where most
programs on the system are installed to.

You can confirm this with the command `which python` 

This will output `/usr/bin/python`

**View a list of the Python versions installed on your system**

```sh
ls -l -a /usr/kbin/python*
```
_______________________________________________________________________________
### An explanation on flags

Note: `-l` and `-a` are called flags, and can actually be combined
like this `-la`. 

```sh
ls -la /usr/bin/python*
```

Flags are added after a command like `ls` to modify its
behavior or output.

`ls` is a built-in shell command that means give me 
a list of the file names and sub-directories inside a directory.

`ls -l` means long format. This means show all the details and not just the
name of each file and sub-directory

`ls -a` means show me all files an sub-directories,
including the hidden ones. Hidden files and sub-directories have a `.`
in front of their name. E.g. `.zshrc` or `.dezly-kingdom`

```sh
ls -la /usr/bin/python*
```

`*` The asterisk after the file path is a wild card. 
It is used for pattern matching. Essentially you are saying this:

"I want a detailed list, 
of all files and sub-directories (including hidden files),
that start with `/usr/bin/python`"
_______________________________________________________________________________
### How to read the output of `ls -la`

```
lrwxrwxrwx root root   7 B  Wed Dec  4 20:05:56 2024  /usr/bin/python ⇒ python3
lrwxrwxrwx root root  14 B  Wed Dec  4 20:05:56 2024  /usr/bin/python-config ⇒ python3-config
lrwxrwxrwx root root  10 B  Wed Dec  4 20:05:56 2024  /usr/bin/python3 ⇒ python3.13
lrwxrwxrwx root root  17 B  Wed Dec  4 20:05:56 2024  /usr/bin/python3-config ⇒ python3.13-config
.rwxr-xr-x root root  14 KB Wed Dec  4 20:05:56 2024  /usr/bin/python3.13
.rwxr-xr-x root root 3.4 KB Wed Dec  4 20:05:56 2024  /usr/bin/python3.13-config
```
_______________________________________________________________________________
`lrwxrwxrwx` The l at the start of this line means this is a `symlink`. 

NOTE: There are no dictories in this output 
but if a line started with a `d` then that means that it is a directory.
_______________________________________________________________________________

`.rwxr-xr-x` The `.` at the start of this line means that this is a file.

Depending on your setup the `.` may be a different colour to show that it is
a binary executable (A file that is actually a program that can be executed).

_______________________________________________________________________________
### A breakdown of the output

`python3.13` and `python3.13-config` are binary executables, 
specifically for python3.13

```
.rwxr-xr-x root root  14 KB Wed Dec  4 20:05:56 2024  /usr/bin/python3.13
.rwxr-xr-x root root 3.4 KB Wed Dec  4 20:05:56 2024  /usr/bin/python3.13-config
```
_______________________________________________________________________________
`python` and `python-config` are symlinks (think of these as shortcuts,
or references that will point to whatever the latest version of python 
on your system is)

```
lrwxrwxrwx root root   7 B  Wed Dec  4 20:05:56 2024  /usr/bin/python ⇒ python3
lrwxrwxrwx root root  14 B  Wed Dec  4 20:05:56 2024  /usr/bin/python-config ⇒ python3-config
```
_______________________________________________________________________________
## Python Versions

```
python --version
```

Python 3.13.1

_______________________________________________________________________________

The python command that is available on your system is actually a symlink.

symlink means that it is not an actual program but a reference that points
to the location of the actual python program that is 
installed on your system.

```
which python
```

To check what version it is pointing to: 

```
python --version
```

Python 3.13.1

_______________________________________________________________________________
