# How to add packages to your project using`uv`
_______________________________________________________________________________
## Search for the package you want in PyPI (The Python Package Index)
[PyPI](https://pypi.org/)
https://pypi.org/

_______________________________________________________________________________
To add a package do the following:

```
cd your-project-directory
uv add name-of-package
```

To add the latest version
```
uv add titanoboa
```

To add a specific version. 
E.g. titanoboa version 0.2.5
````
uv add titanoboa==0.2.5
````

To remove a package:
```
uv remove name-of-package
```
_______________________________________________________________________________
### uv sync

Use this command if you have downloaded a Python project 
that was setup using uv.
```
uv sync 
```

```toml
[project]
name = "syntax-recap"
version = "0.1.0"
description = "Add your description here"
readme = "README.md"
requires-python = ">=3.13"
dependencies = [
    "titanoboa>=0.2.5",
    "vyper>=0.4.0",
]
```

uv sync will look at the `pyproject.toml` file in the project
and then install or remove any packages, 
until the project environment is in order.

_______________________________________________________________________________
