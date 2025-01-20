# This is how you import `titanoboa`
import boa

#______________________________________________________________________________
# This is a built-in Python module
import sys

# NOTE: How to make sure that a script only runs 
# with specific version of Python

# if sys.version_info > (3, 12):
#     raise RuntimeError("Your Python version should not be higher than 3.12")

#______________________________________________________________________________

print(boa.eval("empty(uint256)"))
# The output will be 0

print("Hello!")
