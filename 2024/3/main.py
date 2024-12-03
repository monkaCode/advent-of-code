import re

with open(r"data.txt", "r") as f:
    data = f.read()
    f.close()

IS_PART_ONE = False

instructions = re.findall("mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)", data)

result = 0
mul_enabled = True

for i in instructions:
    if i == "do()":
        mul_enabled = True
        continue
    elif i == "don't()":
        mul_enabled = False
        continue

    numbers = [int(n) for n in i.replace("mul(", "").replace(")", "").split(",")]
    if mul_enabled or IS_PART_ONE:
        result += numbers[0] * numbers[1]

print(result)