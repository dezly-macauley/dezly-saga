def main():
    # This is how you print a simple message in Python
    print("The following command was used:")
    print("python 01_basic_structure.py")
    print("The Python interpreter executed this file as a standalone program")

def my_module_message():
    print("Hey")
    
# NOTE: `if __name__ == __main__` 

# This is a special `if` statement in Python that checks how a file is being
# run when you use the command: `python name_of_your_file.py`

# There are two ways that a file can be run:
# 1. As a standalone program
# 2. As a module (a bunch of re-usable code),
# that is imported into another file

# `__name__` is a built-in variable in Python.
# When you run a file as as a standalone program, 
# `__name__` will automatically set its value to `__main__`

# So this line simply means,
# If this file is run using the command `python deploy_favorites.py`, 
# then that menas this is being run as a standalone program
# If this is being run as a standalone program then call the function called
# `main()` in this file
if __name__ == "__main__":
    main()
else:
    my_module_message()
