# 1. This will import the `datetime` class, 
# from Python's built-in `datetime` module.
from datetime import datetime

# Import the variable types `date` and `time`
from datetime import date, time

# The datetime class is a template that is used to create dat.
# Variables that are created from the datetime class will 
# be declared as the variable datetime.
# Any variable with this variable type will have the the methods of the 
# datetime class

#______________________________________________________________________________
# SECTION: Create an instance of the datetime class

my_datetime: datetime = datetime.now()
print(f"The value of my_datetime is {my_datetime}")

# NOTE: The default syntax is: YYYY-MM-DD HH:MM:SS.mmmmmm

# The value of my_datetime is 2025-02-02 20:55:11.160299

#______________________________________________________________________________
# SECTION: How to display only the date 

my_current_date: date = my_datetime.date()
print(f"The current date is {my_current_date}")
# The current date is 2025-02-02

#______________________________________________________________________________
# SECTION: How to display only the time

my_current_time: time = my_datetime.time()
print(f"The current time is {my_current_time}")
# The current time is 21:17:08.248676

#______________________________________________________________________________
