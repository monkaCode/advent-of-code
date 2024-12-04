with open(r"/Users/felix.hutt/Documents/projects/aoc/2024/4/data.txt") as f:
    data = f.readlines()
    f.close()

def has_xmas(i, j):
    if data[i][j] != "A" or i == 0 or j == 0 or i == len(data) - 1 or j == len(data[i]) - 2:
        return False
    if (data[i-1][j-1] == "M" and data[i+1][j+1] == "S" or data[i-1][j-1] == "S" and data[i+1][j+1] == "M") and (data[i-1][j+1] == "M" and data[i+1][j-1] == "S" or data[i-1][j+1] == "S" and data[i+1][j-1] == "M"):
        return True

xmas = 0

for i, line in enumerate(data):
    for j, char in enumerate(line):
        if char == "A":
            xmas += 1 if has_xmas(i, j) else 0

print(xmas)