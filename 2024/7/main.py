import copy

with open(r"data.txt", "r") as f:
    data = f.readlines()
    f.close()

x = 0
solvable = []
for j, calc in enumerate(data):
    solv = False
    result = int(calc.split(":")[0])
    term_values = [int(x) for x in calc.split(": ")[1].split(" ")]
    #print(result, term_values)

    operators = {"+": term_values[0] + term_values[1], "*": term_values[0] * term_values[1]}
    operatorsv2 = {"+": term_values[0] + term_values[1], "*": term_values[0] * term_values[1]}
    for i in range(2, len(term_values)):
        for op in operators:
            if len(op) == i - 1:
                operatorsv2[f"{op}+"] = operators[op] + term_values[i]
                operatorsv2[f"{op}*"] = operators[op] * term_values[i]
        operators = copy.deepcopy(operatorsv2)
    
    for op in operatorsv2:
        if operatorsv2[op] == result and len(op) == len(term_values) - 1:
            x += result
            #print(j+1, x)
            solvable.append(calc)
            break

print(len(data), len(solvable), x)