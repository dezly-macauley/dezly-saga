# NOTE: To use these commands:
# `make name-of-command`
# E.g. 
# `make status`
# `make start`, 

# Check if the PostgreSQL server is working
status:
	@sudo systemctl status postgresql

# Start the  PostgreSQL server
start:
	@sudo systemctl start postgresql

# Stop the  PostgreSQL server
stop:
	@sudo systemctl stop postgresql

# Interact with your database using `pgcli`
pgcli:
	@pgcli \
	--dbname practice_db \
        --host 127.0.0.1 \
        --port 5432 \
	--username dezly_macauley \
        --pgclirc $(pgcli_config) 
