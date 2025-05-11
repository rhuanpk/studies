import sys

# get names of attributes in sys module
dir(sys)

# get names of attributes for current module
dir()

# create a new variable 'a' and get names of attributes for current module again
a = 5
dir()

# delete/remove a name and get names of attributes for current module again
del a
dir()

