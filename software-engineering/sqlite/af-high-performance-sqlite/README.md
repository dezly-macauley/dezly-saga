# Aaron Francis - High Performance SQLite
_______________________________________________________________________________
## Setup Instructions

If you have cloned this repo run the following commands

Install all of the project dependencies.
```
uv sync
```

Give direnv permission to setup the virtual environment
```
direnv allow
```
_______________________________________________________________________________
## How to create this directory from scratch

### Install a version of Python for `uv` to use. 

NOTE: Even if your system has this Python version, 
you should still run this command.
```
uv python install 3.13.1
```

To see what versions of Python were installed by uv:
```
cd $HOME/.local/share/uv/python/ && ls
```

You will get an output that looks like this
```
cpython-3.12.8-linux-x86_64-gnu
cpython-3.13.1-linux-x86_64-gnu
```
_______________________________________________________________________________
### Choose a location where you want to create the directory:

I have chosen this location:

`dezly-saga/software-engineering/sqlite/af-high-performance-sqlite`
_______________________________________________________________________________
### Create the project directory and enter it
```
mkdir af-high-performance-sqlite
cd af-high-performance-sqlite
```
_______________________________________________________________________________
### Create a new Python project with `uv` inside the directory
```
uv init
```

Yout can delete the `hello.py` file

_______________________________________________________________________________
### Make sure that the repo is using the correct version of Python

E.g. I want Python version 3.13

Open `.python-version` and make sure the version is set to 3.13
```
3.13
```

Open `pyproject.toml` and make sure that the range 
contains your Python version.

(This means any version of Python that is 3.13 or higher)
```
requires-python = ">=3.13"
```

NOTE: If you wanted to specify a range then you can do the following:

`requires-python = ">=3.9, <3.13"`

_______________________________________________________________________________
### Create a virtual environment that contains the version of Python that you want
```
uv venv -p 3.13
```
_______________________________________________________________________________
### Create a `.gitignore` file in the root of your project

Add the following to the file:

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
### Create a practice database in the root of your project

```
practice_db.sqlite
```
_______________________________________________________________________________
### Create a directory for `litecli` configuration files:

```
mkdir litecli-configs
```

Add the following file to `litecli-configs`

`litecli-configs/default_db_config`

Add this inside the file:
```ini
[main]

# Multi-line mode allows breaking up the sql statements into multiple lines. If
# this is set to True, then the end of the statements must have a semi-colon.
# If this is set to False then sql statements can't be split into multiple
# lines. End of line (return) is considered as the end of the statement.
multi_line = True 

# Destructive warning mode will alert you before executing a sql statement
# that may cause harm to the database such as "drop table", "drop database"
# or "shutdown".
destructive_warning = True

# Default log level. Possible values: "CRITICAL", "ERROR", "WARNING", "INFO"
# and "DEBUG". "NONE" disables logging.
log_level = NONE 

# Table format. Possible values:
# ascii, double, github, psql, plain, simple, grid, fancy_grid, pipe, orgtbl,
# rst, mediawiki, html, latex, latex_booktabs, textile, moinmoin, jira,
# vertical, tsv, csv.
# Recommended: ascii
table_format = ascii

# Syntax coloring style. Possible values (many support the "-dark" suffix):
# manni, igor, xcode, vim, autumn, vs, rrt, native, perldoc, borland, tango, emacs,
# friendly, monokai, paraiso, colorful, murphy, bw, pastie, paraiso, trac, default,
# fruity.
# Screenshots at http://mycli.net/syntax
syntax_style = fruity 

# Keybindings: Possible values: emacs, vi.
# Emacs mode: Ctrl-A is home, Ctrl-E is end. All emacs keybindings are available in the REPL.
# When Vi mode is enabled you can use modal editing features offered by Vi in the REPL.
key_bindings = emacs

# Enabling this option will show the suggestions in a wider menu. Thus more items are suggested.
wider_completion_menu = False

# Autocompletion is on by default. This can be truned off by setting this
# option to False. Pressing tab will still trigger completion.
autocompletion = True

prompt = '\x1b[1;38;5;134m\d\x1b[0m\n🔎 '
prompt_continuation = '-> '

# Show/hide the informational toolbar with function keymap at the footer.
show_bottom_toolbar = False 

# Skip intro info on startup and outro info on exit
less_chatty = True 

# Use alias from --login-path instead of host name in prompt
login_path_as_host = False

# Cause result sets to be displayed vertically if they are too wide for the current window,
# and using normal tabular format otherwise. (This applies to statements terminated by ; or \G.)
auto_vertical_output = False

# keyword casing preference. Possible values "lower", "upper", "auto"
keyword_casing = auto

# disabled pager on startup
enable_pager = True

# Custom colors for the completion menu, toolbar, etc.
[colors]
completion-menu.completion.current = 'bg:#ffffff #000000'
completion-menu.completion = 'bg:#008888 #ffffff'
completion-menu.meta.completion.current = 'bg:#44aaaa #000000'
completion-menu.meta.completion = 'bg:#448888 #ffffff'
completion-menu.multi-column-meta = 'bg:#aaffff #000000'
scrollbar.arrow = 'bg:#003333'
scrollbar = 'bg:#00aaaa'
selected = '#ffffff bg:#6666aa'
search = '#ffffff bg:#4444aa'
search.current = '#ffffff bg:#44aa44'
bottom-toolbar = 'bg:#222222 #aaaaaa'
bottom-toolbar.off = 'bg:#222222 #888888'
bottom-toolbar.on = 'bg:#222222 #ffffff'
search-toolbar = 'noinherit bold'
search-toolbar.text = 'nobold'
system-toolbar = 'noinherit bold'
arg-toolbar = 'noinherit bold'
arg-toolbar.text = 'nobold'
bottom-toolbar.transaction.valid = 'bg:#222222 #00ff5f bold'
bottom-toolbar.transaction.failed = 'bg:#222222 #ff005f bold'

# style classes for colored table output
output.header = "#00ff5f bold"
output.odd-row = ""
output.even-row = ""

# Favorite queries.
[favorite_queries]
# example:
# long_query = SELECT name, type \
#             FROM sqlite_master \
#             WHERE type IN ('table', 'view') \
#             AND name NOT LIKE 'sqlite_%';

# Startup commands
# litecli commands or sqlite commands to be executed on startup.
# some of them will require you to have a database attached.
# they will be executed in the same order as they appear in the list.
[startup_commands]
#commands = ".tables", "pragma foreign_keys = ON;"
# vi: ft=ini
```
_______________________________________________________________________________
### Create a `.envrc` file in the root directory of the project

Add the following to the file:
```bash
source .venv/bin/activate
export litecli_configs="./litecli-configs"
export selected_config="default_db_config"
export litecli_config_path="$litecli_configs/$selected_config"

# NOTE: The line below has nothing to do with `direnv`
# This is simply to tell Neovim to treat this a bash file, 
# .envrc is not a recognized file type.
# vi: ft=bash
```
_______________________________________________________________________________
### Create a `Makefile` in the root directory of the project

Add this to the `Makefile`

```Makefile
# NOTE: To use this command:
# `make run`

run:
	 @litecli --liteclirc $(litecli_config_path) practice_db.sqlite
```

_______________________________________________________________________________
### Close the repo and navigate to it via the terminal

You should see this error when you enter

direnv: 
error dezly-saga/software-engineering/sqlite/af-high-performance-sqlite/
.envrc is blocked. Run `direnv allow` to approve its content

This is because you have to give `direnv` permission to run the `.envrc` file

Open the terminal and run this command:
```
direnv allow
```
_______________________________________________________________________________
### While in the directory, check your Python version

```
which python
```

It should point to the version of Python in the `.venv` directory 

...af-high-performance-sqlite/.venv/bin/python
_______________________________________________________________________________
### Install litecli and sqlite-utils

```
uv add litecli
uv add sqlite-utils
```

_______________________________________________________________________________
### To start litecli using the Makefile

```
make run
```

_______________________________________________________________________________
### To exit

```
exit
```
_______________________________________________________________________________
