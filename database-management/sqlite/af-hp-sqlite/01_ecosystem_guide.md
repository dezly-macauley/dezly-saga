How to check what version of SQLite you have:

```
sqlite3 -version
```

To enter an in-memory database you would just use the default cli
```
sqlite3
```

To exit:
```
.exit
```
_______________________________________________________________________________
## A better way to interact with your database

First make sure that you have the following tools installed on your machine:

`uv` A Rust-powered Python package manager, and Python environment manager. 

`direnv` A program that will automatically set up a directory for you
when you enter that directory, based on what you put inside a `.envrc` file

_______________________________________________________________________________

Use uv to install the version of Python you want:
```
uv python install 3.13
```

To see what versions of Python were installed by uv:
```
cd $HOME/.local/share/uv/python/ && ls
```

Create a repo to save your SQL queries
```
mkdir name-of-repo
```

initialize the repo:
```
cd name-of-repo
uv init
```

Create a virtual environment that contains the version of Python that you want
```
uv venv -p python3.13
```
_______________________________________________________________________________

### Additional commands: 

#### uv sync

If you have just download a uv project (E.g. Using `git clone` ),
then use the command:
```
uv sync
```
This will look at the `pyproject.toml` file, 
and install all of the dependencies needed by the project.

_______________________________________________________________________________

#### uv tool install

Use this command as a way to install Python tools globally ( E.g. When you 
want certain commands to be available outside of your repo, 
such as language support protocols).

_______________________________________________________________________________
To activate the virtual environment, 
you will need to enter this command everytime you enter the directory:

```
source .venv/bin/activate
```

and you will also need to run this command to deactivate 
the virtual environment.
```
deactivate
```

This is quite tedious and so that is where `direnv` comes in.

_______________________________________________________________________________
Open up your `.zshrc` file. It is usually located at `$HOME`

Add the following lines to the file:
```sh

# This is required for `direnv` to work
# direnv is a program that will automatically load the `.env` file in a
# directory when you enter that directory
eval "$(direnv hook zsh)"

# NOTE: This line below will disable the messages that direnv displays when 
# you enter a directory 

export DIRENV_LOG_FORMAT=""
```
_______________________________________________________________________________

Open the project directory and create a `.envrc` file inside the repository.

Add this line inside the `.envrc` file
```
source .venv/bin/activate
```

Close the terminal and open the directory again.

You will see a message that says the the `.envrc` file is blocked:
```
direnv: error 
/home/dezly-macauley/dezly-saga/database-management/sqlite/af-hp-sqlite/.envrc 
is blocked. Run `direnv allow` to approve its content
```

To fix this, run the command:
```
direnv allow
```

Now everytime you enter this directory, your Python setup will be loaded.

To confirm that this is working, run the command:
```
which python
```

If you have done everything correctly, 
it should be using the version of Python from the virtual environment in
your project.
```
/home/dezly-macauley/dezly-saga/database-management/sqlite/af-hp-sqlite/.venv/bin/python
```
_______________________________________________________________________________
## Install `litecli`

Now you can install `litecli` in your project. 

This is a better version of the default cli `sqlite3` that is part of SQLite.

litecli has syntax-highlighting and autocomplete.

To add `litecli` or any other Python package to your repo,
just go to `https://pypi.org/`
```
cd name-of-repo
uv add litecli
```

_______________________________________________________________________________

To open `litecli`, open a new terminal and type this:
```
litecli
```

To exit:
```
exit
```


_______________________________________________________________________________
