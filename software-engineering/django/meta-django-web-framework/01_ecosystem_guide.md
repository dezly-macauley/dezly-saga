# How to setup a new Django project
_______________________________________________________________________________

NOTE: I use a package manager called `uv`
And I use a tool called `direnv` to automatically 
activate my virtual environment.
_______________________________________________________________________________
## Use uv to install a version if Python

```
uv python install 3.13
```

NOTE: Even if you already have this version of Python, you want to install
it using uv. This is because you want to ensure that your project is using
a specific version of Python and not the one used by your system which can
change whenever you updated your system.

_______________________________________________________________________________
## Create the project directory

```
mkdir name-of-project
```
_______________________________________________________________________________
## Initialize the directory with `uv`

```
cd name-of-project
uv init 
```

_______________________________________________________________________________
## Create a `.gitignore`

NOTE: This is normally created for you but if if isn't, create a file called
`.gitignore`

Then open the file and put this inside:
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
## Create the virtual environment

Create a Python virtual environment using the version of Python that you
installed with uv.
```
uv venv -p python3.13
```
_______________________________________________________________________________
## Setup `direnv` to automatically load the virtual environment

Open the directory and add a `.envrc` file

Then add this line inside the file:

```
 source .venv/bin/activate
```

Exit the directory, enter the directory again and run this command:
```
direnv allow
```

_______________________________________________________________________________
## Create the Django project

First add Django as a dependency:
```
cd name-of-project
uv add django
```

Create the Django project structure inside the current directory
```
django-admin startproject name_of_project .
```

NOTE:
1. Don't forget the `.` at the end. That means this directory
2. Convert any hyphens to underscores when you use this command.

_______________________________________________________________________________
## Run the development server

Open a seperate terminal and run this there:
```
python manage.py runserver
```

You will get a message like this:

```
January 15, 2025 - 12:18:21
Django version 5.1.5, using settings 'meta_django_web_framework.settings'
Starting development server at http://127.0.0.1:8000/
Quit the server with CONTROL-C.
```

To view your Python project open a web browser like firefox and put that
link in the search bar:

```
http://127.0.0.1:8000
```

_______________________________________________________________________________
