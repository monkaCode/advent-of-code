with open(r"/Users/felix.hutt/Documents/projects/aoc/2024/4/data.txt") as f:
    data = f.readlines()
    f.close()

def has_horizontal_right(line: str, index: int):
    if line.find("XMAS", index) == index:
        return True

def has_horizontal_left(line: str, index: int):
    if index < 3:
        return False
    if line.find("SAMX", index-3) == index-3:
        return True

def has_vertical_up(data, i, j):
    if i < 3:
        return False
    if data[i][j] == "X" and data[i-1][j] == "M" and data[i-2][j] == "A" and data[i-3][j] == "S":
        return True

def has_vertical_down(data, i, j):
    if i+3 >= len(data):
        return False
    if data[i][j] == "X" and data[i+1][j] == "M" and data[i+2][j] == "A" and data[i+3][j] == "S":
        return True

def has_diagonal_up_right(data, i, j):
    if i < 3:
        return False
    if data[i][j] == "X" and data[i-1][j+1] == "M" and data[i-2][j+2] == "A" and data[i-3][j+3] == "S":
        return True

def has_diagonal_up_left(data, i, j):
    if i < 3 or j < 3:
        return False
    if data[i][j] == "X" and data[i-1][j-1] == "M" and data[i-2][j-2] == "A" and data[i-3][j-3] == "S":
        return True
    
def has_diagonal_down_right(data, i, j):
    if i+3 >= len(data):
        return False
    if data[i][j] == "X" and data[i+1][j+1] == "M" and data[i+2][j+2] == "A" and data[i+3][j+3] == "S":
        return True

def has_diagonal_down_left(data, i, j):
    if i+3 >= len(data) or j < 3:
        return False
    if data[i][j] == "X" and data[i+1][j-1] == "M" and data[i+2][j-2] == "A" and data[i+3][j-3] == "S":
        return True

xmas = 0

for i, line in enumerate(data):
    for j, char in enumerate(line):
        if char == "X":
            xmas += 1 if has_horizontal_right(line, j) else 0
            xmas += 1 if has_horizontal_left(line, j) else 0
            xmas += 1 if has_vertical_up(data, i, j) else 0
            xmas += 1 if has_vertical_down(data, i, j) else 0
            xmas += 1 if has_diagonal_down_left(data, i, j) else 0
            xmas += 1 if has_diagonal_down_right(data, i, j) else 0
            xmas += 1 if has_diagonal_up_left(data, i, j) else 0
            xmas += 1 if has_diagonal_up_right(data, i, j) else 0

print(xmas)