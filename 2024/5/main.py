with open(r"data.txt", "r") as f:
    data = f.readlines()
    f.close()

rules = {}
updates = []

rules_section = True
for line in data:
    line = line.replace("\n", "")
    if line == "":
        rules_section = False
        continue
    if rules_section and int(line.split("|")[0]) not in rules:
        rules[int(line.split("|")[0])] = [int(line.split("|")[1])]
    elif rules_section:
        rules[int(line.split("|")[0])].append(int(line.split("|")[1]))
    else:
        updates.append([int(n) for n in line.split(",")])

correct_updates_1 = []
correct_updates_2 = []

def is_updated_correct(update):
    correct_update = True
    for i, u in enumerate(update):
        if u not in rules and i < len(update) - 1:
            temp = update[i-1]
            update[i-1] = u
            update[i] = temp
            return is_updated_correct(update)
        if i == 0 and u in rules:
            continue
        if u in rules[update[i-1]]:
            continue
        else:
            temp = update[i-1]
            update[i-1] = u
            update[i] = temp
            return is_updated_correct(update)
    return correct_update

for update in updates:
    if is_updated_correct(update):
        correct_updates_1.append(update)

x = 0
for cu in correct_updates_1:
    middle_value = cu[len(cu)//2]
    x += middle_value

print(x)