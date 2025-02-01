# Aaron Francis - Mastering PostgreSQL
_______________________________________________________________________________
## Phase 1 - PostgreSQL Setup

_______________________________________________________________________________
### Install PostgreSQL

This is for Arch Linux:
```
sudo pacman -S --needed postgresql
```

_______________________________________________________________________________
### Iniatialize PostgreSQL and start the service

PostgreSQL needs to set up its database files before you can use it:

Run this command from your home directory:
```sh
sudo -u postgres initdb -D /var/lib/postgres/data \
--locale=C.UTF-8 \
--encoding=UTF8 \
--data-checksums \
--auth-local=peer \
--auth-host=scram-sha-256

```
_______________________________________________________________________________
### Start the PostgreSQL service

Unlike SQLite which is an embedded database, 
PostgreSQL is a client-server database. 

SQLite is an embedded database, meaning the database engine is integrated into the application that uses it. There's no need for a separate server process. Instead, SQLite reads and writes directly to a database file on disk, making it lightweight and self-contained.

On the other hand, PostgreSQL is a client-server database. It requires a server process to run separately and handle database operations. Applications (clients) connect to this server using a network protocol, even if the client and server are on the same machine.

**Run this command to check if the PostgreSQL server is running**
```
sudo systemctl status postgresql
```

To start the service / server
```
sudo systemctl start postgresql
```

To stop the service / server
```
sudo systemctl stop postgresql
```
_______________________________________________________________________________
## Phase 2 -  Use the `psql` cli to create a PostgreSQL user and database

_______________________________________________________________________________
Log into the `psql` the default PostgreSQL command line tool as the the 
default user. The default user is called `postgres`

Run this command. The `u` stands for user
```sh
sudo -u postgres psql
```

Your prompt will now look like this:
```
postgres=#
```

This is the prompt that you will use for admin tasks like creating users,
and creating databases.
_______________________________________________________________________________
### Create a PostgreSQL user

**Run these commands while logged into the `postgres=#` prompt**

Create your username and password:

NOTE: The username should not contain captal letter, hyphens, spaces, or
any other characters.

```sql
CREATE USER dezly_macauley WITH SUPERUSER; 
```

To check that the user was created run this command:
```sql
\du
```

The output
```sql
                                        List of roles
   Role name    |                         Attributes
----------------+------------------------------------------------------------
 dezly_macauley | Superuser
 postgres       | Superuser, Create role, Create DB, Replication, Bypass RLS
```

_______________________________________________________________________________
### Use the interactive command to set the password for the user

This will allow you to set the password interactively:
```sql
\password dezly_macauley
```
_______________________________________________________________________________
### Create a practice database

```sh
CREATE DATABASE practice_db;
```
_______________________________________________________________________________
### Give your user permission to interact with the database

```sql
GRANT ALL PRIVILEGES ON DATABASE practice_db 
TO dezly_macauley;
```

To view the list of databases:
```sql
\l
```
_______________________________________________________________________________

To exit the prompt:
```sql
\q
```
_______________________________________________________________________________
## Phase 3 - Setting up your PostgreSQL workspace

**NOTE: Run these commands from your normal shell**
_______________________________________________________________________________
### If you have cloned this repo run the following commands

**NOTE: If you are building this repo from scratch, go to the next step**

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

```
dezly-saga/software-engineering/postgresql/af-mastering-postgres
```

_______________________________________________________________________________
### Create the project directory and enter it
```
mkdir af-mastering-postgres
cd af-mastering-postgres
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
### Create a `pgcli configuration` file

I will be using `pgcli` for my setup. 

It allows you to interact with a PostgreSQL database from the terminal,
while having autocompletion and syntax highlighting.

```
touch pgcli_config
```
_______________________________________________________________________________
### Configure `pgcli`

Open the `pgcli_config` file and add replace the contents with this:
```ini
[main]

# Multi-line mode allows breaking up the sql statements into multiple lines. If
# this is set to True, then the end of the statements must have a semi-colon.
# If this is set to False then sql statements can't be split into multiple
# lines. End of line (return) is considered as the end of the statement.
multi_line = True 

# Destructive warning will alert you before executing a sql statement
# that may cause harm to the database such as "drop table", "drop database",
# "shutdown", "delete", or "update".
# You can pass a list of destructive commands or leave it empty if you want to skip all warnings.
# "unconditional_update" will warn you of update statements that don't have a where clause
destructive_warning = drop, shutdown, delete, truncate, alter, update, unconditional_update

# Default log level. Possible values: "CRITICAL", "ERROR", "WARNING", "INFO"
# and "DEBUG". "NONE" disables logging.
log_level = NONE 

# Table format. Possible values: psql, plain, simple, grid, fancy_grid, pipe,
# ascii, double, github, orgtbl, rst, mediawiki, html, latex, latex_booktabs,
# textile, moinmoin, jira, vertical, tsv, csv, sql-insert, sql-update,
# sql-update-1, sql-update-2 (formatter with sql-* prefix can format query
# output to executable insertion or updating sql).
# Recommended: psql, fancy_grid and grid.
table_format = psql

# Syntax Style. Possible values: manni, igor, xcode, vim, autumn, vs, rrt,
# native, perldoc, borland, tango, emacs, friendly, monokai, paraiso-dark,
# colorful, murphy, bw, pastie, paraiso-light, trac, default, fruity
syntax_style = fruity 

# Enables context sensitive auto-completion. If this is disabled, all
# possible completions will be listed.
smart_completion = True

# Display the completions in several columns. (More completions will be
# visible.)
wider_completion_menu = False

# Do not create new connections for refreshing completions; Equivalent to
# always running with the --single-connection flag.
always_use_single_connection = False

# If multi_line_mode is set to "psql", in multi-line mode, [Enter] will execute
# the current input if the input ends in a semicolon.
# If multi_line_mode is set to "safe", in multi-line mode, [Enter] will always
# insert a newline, and [Esc] [Enter] or [Alt]-[Enter] must be used to execute
# a command.
multi_line_mode = psql

# When `destructive_warning` is on and the user declines to proceed with a
# destructive statement, the current transaction (if any) is left untouched,
# by default. When setting `destructive_warning_restarts_connection` to
# "True", the connection to the server is restarted. In that case, the
# transaction (if any) is rolled back.
destructive_warning_restarts_connection = False

# When this option is on (and if `destructive_warning` is not empty),
# destructive statements are not executed when outside of a transaction.
destructive_statements_require_transaction = False

# Enables expand mode, which is similar to `\x` in psql.
expand = False

# Enables auto expand mode, which is similar to `\x auto` in psql.
auto_expand = False

# Auto-retry queries on connection failures and other operational errors. If
# False, will prompt to rerun the failed query instead of auto-retrying.
auto_retry_closed_connection = True

# If set to True, table suggestions will include a table alias
generate_aliases = False

# Path to a json file that specifies specific table aliases to use when generate_aliases is set to True
# the format for this file should be:
# {
#     "some_table_name": "desired_alias",
#     "some_other_table_name": "another_alias"
# }
alias_map_file =

# log_file location.
# In Unix/Linux: ~/.config/pgcli/log
# In Windows: %USERPROFILE%\AppData\Local\dbcli\pgcli\log
# %USERPROFILE% is typically C:\Users\{username}
log_file = default

# keyword casing preference. Possible values: "lower", "upper", "auto"
keyword_casing = auto

# casing_file location.
# In Unix/Linux: ~/.config/pgcli/casing
# In Windows: %USERPROFILE%\AppData\Local\dbcli\pgcli\casing
# %USERPROFILE% is typically C:\Users\{username}
casing_file = default

# If generate_casing_file is set to True and there is no file in the above
# location, one will be generated based on usage in SQL/PLPGSQL functions.
generate_casing_file = False

# Casing of column headers based on the casing_file described above
case_column_headers = True

# history_file location.
# In Unix/Linux: ~/.config/pgcli/history
# In Windows: %USERPROFILE%\AppData\Local\dbcli\pgcli\history
# %USERPROFILE% is typically C:\Users\{username}
history_file = default

# Order of columns when expanding * to column list
# Possible values: "table_order" and "alphabetic"
asterisk_column_order = table_order

# Whether to qualify with table alias/name when suggesting columns
# Possible values: "always", "never" and "if_more_than_one_table"
qualify_columns = if_more_than_one_table

# When no schema is entered, only suggest objects in search_path
search_path_filter = False

# Default pager. See https://www.pgcli.com/pager for more information on settings.
# By default 'PAGER' environment variable is used. If the pager is less, and the 'LESS'
# environment variable is not set, then LESS='-SRXF' will be automatically set.
# pager = less

# Timing of sql statements and table rendering.
timing = True

# Show/hide the informational toolbar with function keymap at the footer.
show_bottom_toolbar = False 

# Keybindings:
# When Vi mode is enabled you can use modal editing features offered by Vi in the REPL.
# When Vi mode is disabled emacs keybindings such as Ctrl-A for home and Ctrl-E
# for end are available in the REPL.
vi = False

# Error handling
# When one of multiple SQL statements causes an error, choose to either
# continue executing the remaining statements, or stopping
# Possible values "STOP" or "RESUME"
on_error = STOP

# Set threshold for row limit. Use 0 to disable limiting.
row_limit = 1000

# Truncate long text fields to this value for tabular display (does not apply to csv).
# Leave unset to disable truncation. Example: "max_field_width = "
# Be aware that formatting might get slow with values larger than 500 and tables with
# lots of records.
max_field_width = 500

# Skip intro on startup and goodbye on exit
less_chatty = True 

# Show all Postgres error fields (as listed in
# https://www.postgresql.org/docs/current/protocol-error-fields.html).
# Can be toggled with \v.
verbose_errors = False

# Postgres prompt
# \t - Current date and time
# \u - Username
# \h - Short hostname of the server (up to first '.')
# \H - Hostname of the server
# \d - Database name
# \p - Database port
# \i - Postgres PID
# \# - "@" sign if logged in as superuser, '>' in other case
# \n - Newline
# \dsn_alias - name of dsn connection string alias if -D option is used (empty otherwise)
# \x1b[...m - insert ANSI escape sequence
# eg: prompt = '\x1b[35m\u@\x1b[32m\h:\x1b[36m\d>'
prompt = '\x1b[1;38;5;134m\d\x1b[0m\n🔎 '
prompt_continuation = '-> '

# Number of lines to reserve for the suggestion menu
min_num_menu_lines = 4

# Character used to left pad multi-line queries to match the prompt size.
multiline_continuation_char = ''

# The string used in place of a null value.
null_string = '<null>'

# manage pager on startup
enable_pager = True

# Use keyring to automatically save and load password in a secure manner
keyring = False 

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
# These three values can be used to further refine the syntax highlighting.
# They are commented out by default, since they have priority over the theme set
# with the `syntax_style` setting and overriding its behavior can be confusing.
# literal.string = '#ba2121'
# literal.number = '#666666'
# keyword = 'bold #008000'

# style classes for colored table output
output.header = "#00ff5f bold"
output.odd-row = ""
output.even-row = ""
output.null = "#808080"

# Named queries are queries you can execute by name.
[named queries]

# Here's where you can provide a list of connection string aliases.
# You can use it by passing the -D option. `pgcli -D example_dsn`
[alias_dsn]
# example_dsn = postgresql://[user[:password]@][netloc][:port][/dbname]

# Format for number representation
# for decimal "d" - 12345678, ",d" - 12,345,678
# for float "g" - 123456.78, ",g" - 123,456.78
[data_formats]
decimal = ""
float = ""

# vi: ft=ini
```
_______________________________________________________________________________
### Ensure that you have `direnv` setup corectly on your system

NOTE: For this to work you need the following:
- The program `direnv` is installed on your system.
- Also you should have the following lines added to your `.zshrc`

**~/.zshrc**
```sh

# This will allow direnv to automatically load the environment variables in
# a `.envrc` file when you enter your project directory.
eval "$(direnv hook zsh)"

# This line below will disable the messages that direnv displays when 
# you enter a directory 
export DIRENV_LOG_FORMAT=""
```
_______________________________________________________________________________
### Create a `.envrc` file in the root of your project

Add the following lines to the file:
```bash
# Activat the Python virtual environment to ensure that this project is
# using the correct version of Python,
# and also that it using the version installed by `uv` and not the version
# of Python that is installed system-wide
source .venv/bin/activate

# The location of the configuration file the `pgcli` tool in this project
export pgcli_config="./pgcli_config"

# The line below has nothing to do with `direnv`
# This is simply to tell Neovim to treat this a bash file, 
# .envrc is not a recognized file type.
# vi: ft=bash
```
_______________________________________________________________________________
### Create a `Makefile` in the root directory of the project

```
touch Makefile
```

Add this to the `Makefile`

```Makefile
# NOTE: To use these commands:
# `make name-of-command`
# E.g. 
# `make status`
# `make start`, 

# Check if the PostgreSQL server is working
status:
	@sudo systemctl status postgresql

# Start the PostgreSQL server
start:
	@sudo systemctl start postgresql

# Stop the PostgreSQL server
stop:
	@sudo systemctl stop postgresql

# Interact with your database using `pgcli`
pgcli:
	@pgcli \
		--dbname practice_db \
        --host 127.0.0.1 \
        --port 5432 \
		--username dezly_macauley \
        --password \
        --pgclirc $(pgcli_config) 
```

`127.0.01` is the same as `localhost`. 
This means the database server is running on your local machine.
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

...af-mastering-postgres/.venv/bin/python
_______________________________________________________________________________
### Install pgcli (Command Line Tool)

```
uv add pgcli
```
_______________________________________________________________________________
### To start pgcli using the Makefile

```sh
make sql
```

_______________________________________________________________________________
### To exit

```
exit
```
_______________________________________________________________________________
## Phase 4 (Optional) - Setup pgAdmin

From the root directory of your project, run this command:
```
uv add pgadmin4 
```
_______________________________________________________________________________
### Create the directories that pgadmin4 needs for its own internal data 

These will be created outside of your project directory:
```sh
sudo mkdir /var/lib/pgadmin
sudo mkdir /var/log/pgadmin
```

Give yourself permission to access these files:
```sh
sudo chown $USER /var/lib/pgadmin
sudo chown $USER /var/log/pgadmin
```
_______________________________________________________________________________
### From your project directory run this command

You need to run this command when you start using pgadmin4:
```sh
pgadmin4
```

NOTE: The first time you log in you will be asked to setup an email and
a password to login to pgadmin4

```sh
NOTE: Configuring authentication for SERVER mode.

Enter the email address and password to use for the initial pgAdmin user account:

Email address: dezlymacauley@proton.me
Password:
Retype password:
Starting pgAdmin 4. Please navigate to http://127.0.0.1:5050 in your browser.
 * Serving Flask app 'pgadmin'
 * Debug mode: off
```

Keep this running and open the link in your browser
```
http://127.0.0.1:5050
```
_______________________________________________________________________________
### To set pgadmin4 to dark mode

Go to `File`, then click `Preferences`

Then scroll down to `Miscellaneous`, then click `Themes`

Select `Dark` and click `Save`

_______________________________________________________________________________
