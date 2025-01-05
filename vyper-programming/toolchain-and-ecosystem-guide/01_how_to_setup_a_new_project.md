# How to setup a new Python project using `uv`
_______________________________________________________________________________
## First make sure that you have `uv` installed

Use this command to install uv:
```
sudo pacman -S --needed uv
```
_______________________________________________________________________________

Check if uv is installed: 
```
which uv             
```

This will output:

```
/usr/bin/uv
```
_______________________________________________________________________________
When programs are installed on a system-wide level 
`/usr/bin/` is the directory where the program is placed.

_______________________________________________________________________________
## Install Python versions using `uv install`

E.g. I want to install `Python 3.13` and `Python 3.12`
```
uv install python3.13
uv install python3.12
```
_______________________________________________________________________________
View **only uv installed Python version**

```
ls $HOME/.local/share/uv/python
```

This will output:
```
cpython-3.12.8-linux-x86_64-gnu
cpython-3.13.1-linux-x86_64-gnu
```
_______________________________________________________________________________
View **all the available Python versions** on your system

This includes uv installs, and installs made by your package manager
```
uv python list
```
_______________________________________________________________________________
## Create a new project with a specific version of Python

This will create a new Python project
```
mkdir example-project
cd example-project
uv init
```

The project structure will look like this:

```
.
├── .git
├── .gitignore
├── hello.py
├── pyproject.toml
├── .python-version
└── README.md
```
_______________________________________________________________________________
It should already have a `.gitignore` template.

If it doesn't, create one:
```
touch .gitignore
```

Put this inside the file
```gitignore
# Python-generated files
__pycache__/
*.py[oc]
build/
dist/
wheels/
*.egg-info

# Virtual environments
.venv
```
_______________________________________________________________________________
### Create a virtual environment for the project

Without doing anything else, 
if you run the command `which python`

You will see that it is currently using the system-wide version of Python
for your project. 

```
usr/bin/python
```

This is the version of Python that was installed by `pacman`, 
the package manager on Arch Linux.

You do not want your project to use this version.

_______________________________________________________________________________
To make the project use a local Python version that was installed by uv,
you have to create a virtual environment and use the `-p` flag 
to specify the version of Python that you want to use in your project.

E.g. To use the `Python3.13` that was installed by uv, 
use the following command:

```
cd example-project
uv venv -p python3.13
```
_______________________________________________________________________________
## Activate the virtual environment 

```
source .venv/bin/activate
```

Now if you run the `which python` command,
it will point to the virtual environment:

```
/home/dezly-macauley/example-project/.venv/bin/python
```

To deactivate the virtual environment, simply type `deactivate`

_______________________________________________________________________________
**Note:**: Closing the terminal will automatically deactivate the environment.

This means that each time you open up the directory, 
you will have to run this command


```
source .venv/bin/activate
```

_______________________________________________________________________________
## How to automatically load the virtual environment



_______________________________________________________________________________
