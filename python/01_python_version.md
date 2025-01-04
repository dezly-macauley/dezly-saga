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
