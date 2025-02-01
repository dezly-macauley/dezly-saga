import os

# Get the absolute path to .pgadmin-setup directory
BASE_DIR = os.path.dirname(os.path.abspath(__file__))

# Data directory
SQLITE_PATH = os.path.join(BASE_DIR, 'data/pgadmin4_sqlite.db')

# Log directory
LOG_FILE = os.path.join(BASE_DIR, 'logs/pgadmin4.log')

# You may also need to override other settings based on your needs
# For example, enabling server mode if not already set
SERVER_MODE = True

# NOTE on `pgadmin_sqlite.db`
# ------------------------------------------------------------------
# 1. The file pgadmin_sqlite.db will be created automatically 
# when pgadmin4 runs for the first time.
# ------------------------------------------------------------------
# 2. This is an SQLite database. SQLite databases are great for
# embedding data directly into an application, or in this case,
# your project.
# ------------------------------------------------------------------
# 3. You do not need to setup anything in this file.
# It is simply where pgadmin4 stores its own internal data
# to manage itself. The data is separate from the data that your
# application uses.
# ------------------------------------------------------------------
