with open(r"data.txt", "r") as f:
    rows = f.readlines()
    f.close()

pos = (0, 0)
facing = None

for y, row in enumerate(rows):
    for x, col in enumerate(row):
        if col == "^":
            pos = (x, y)
            facing = "up"
        elif col == ">":
            pos = (x, y)
            facing = "right"
        elif col == "v":
            pos = (x, y)
            facing = "down"
        elif col == "<":
            pos = (x, y)
            facing = "left"
    rows[y] = [c for c in row.replace("\n", "")]

print(pos)
print(rows[pos[1]])
while (pos[0] >= 0 and pos[0] <= len(rows[0]) - 1 and pos[1] >= 0 and pos[1] <= len(rows[0]) - 1):
    try:
        if facing == "up":
            if rows[pos[1]-1][pos[0]] != "#":
                rows[pos[1]][pos[0]] = "X"
                rows[pos[1]-1][pos[0]] = "^"
                pos = (pos[0], pos[1]-1)
            else:
                facing = "right"
                rows[pos[1]][pos[0]] = ">"
        elif facing == "right":
            if rows[pos[1]][pos[0]+1] != "#":
                rows[pos[1]][pos[0]] = "X"
                rows[pos[1]][pos[0]+1] = ">"
                pos = (pos[0]+1, pos[1])
            else:
                facing = "down"
                rows[pos[1]][pos[0]] = "v"
        elif facing == "down":
            if rows[pos[1]+1][pos[0]] != "#":
                rows[pos[1]][pos[0]] = "X"
                rows[pos[1]+1][pos[0]] = "v"
                pos = (pos[0], pos[1]+1)
            else:
                facing = "left"
                rows[pos[1]][pos[0]] = "<"
        elif facing == "left":
            if rows[pos[1]][pos[0]-1] != "#":
                rows[pos[1]][pos[0]] = "X"
                rows[pos[1]][pos[0]-1] = "<"
                pos = (pos[0]-1, pos[1])
            else:
                facing = "up"
                rows[pos[1]][pos[0]] = "^"
    except:
        rows[pos[1]][pos[0]] = "X"
        if facing == "up":
            pos = (pos[0], pos[1]-1)
        elif facing == "right":
            pos = (pos[0]+1, pos[1])
        elif facing == "down":
            pos = (pos[0], pos[1]+1)
        elif facing == "left":
            pos = (pos[0]-1, pos[1])

x = 0
for row in rows:
    for col in row:
        x += 1 if col == "X" else 0

print(x)