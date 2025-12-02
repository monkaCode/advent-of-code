import copy

with open(r"data.txt", "r") as f:
    input = f.read()

disk = []
ID = 0
for i, c in enumerate(input):
    if i % 2 == 0:
        for x in range(int(c)):
            disk.append(ID)
        ID += 1
    else:
        for x in range(int(c)):
            disk.append(".")

print(disk)
reverted_disk = copy.deepcopy(disk)
reverted_disk.reverse()

if disk == reverted_disk:
    print("error")
j = 0
for i in range(len(disk)):
    if reverted_disk[j] == ".":
        j += 1
    if disk[i] != ".":
        continue
    if reverted_disk[j] == ".":
        j += 1
        continue
    disk[i] = reverted_disk[j]
    disk[len(disk) - 1 - i] = "."

result = 0
for i, x in enumerate(disk):
    if x != ".":
        result += x * i
    else:
        print(f"{i} ", end="")
print(result)

with open(r"output.txt", "w") as f:
    f.write("".join(str(disk)))
    f.close()