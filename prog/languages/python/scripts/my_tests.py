print("""
    hello world
jovens
""")

age = 22
name = 'rhuan'

print("{0} was {1} years old when he wrote this book".format(name, age))
print("Why is {0} playing with that python?".format(name))
print(name, "tem", age, "anos de idade!")
name = name + " patriky"
print(name, "tem", age, "anos de idade!")
# print(name + "tem" + age + "anos de idade!")
print(f"Me chamo {name} e tenho {age} anos de idade!")
myfloat = 10/3
print("valor: {:.2f}\n".format(1/3))

while True:
    answer = input("enter with \"any\": ")
    if answer == "any":
        print("nice =)")
        break
    print("Length of answer:", len(answer))

def default_value(any_string, any_float = 12.5):
    print(
        "\nstring passada:", any_string,
        "\nfloat passado:", any_float
    )
default_value("ola mundo")
default_value("hello world", 15.2)

print("----------")

# def myhelp(some_string: str):
#    print(f"{some_string}".__doc__)
# 
# myhelp("int")

def soma(first: int, second: int) -> int:
    return first + second

print(soma(3, 5))

print("----------")

# import sys
# print("\nThe command line arguments are:")
# for i in sys.argv:
#     print(i)
# print(f"""
#     \rThe PYTHONPATH is {sys.path}
# """)

print("----------")

import os; print(f"\n{os.getcwd()}")

print("----------")

from sys import argv; print(argv)

print("----------")

def hello():
	"""Says a simple "hello world!"."""
	print("hello world!")
print(hello.__doc__)
