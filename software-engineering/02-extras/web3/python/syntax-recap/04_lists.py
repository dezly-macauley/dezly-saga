# This is a list that should contain integers 
my_list: list[int] = [ 7, 78, 12]
print(my_list)

print(type(my_list))
# NOTE: The variable type will be shown as `list`

# This is a list that can contain strings and letters
my_mixed_list: list[int | str] = [ "Dezly", 7, 50, "Hello"]
print(my_mixed_list)
print(type(my_mixed_list))

my_mixed_list.append("Europe")
print(my_mixed_list)
