# 
#   Вычислите расстояние редактирования двух данных непустых строк,
#   содержащих строчные буквы латинского алфавита.
# 

import math

def my_diff(a, b):
    if a == b:
        return 0
    return 1

def editing_distance(A, B):
    D = []
    for i in range(len(A) + 1): # initializing
        D.append([])
        for j in range(len(B) + 1):
            if j == 0:
                D[i].append(i)
            elif i == 0:
                D[i].append(j)
            else:
                D[i].append(math.inf)

    i = 1
    j = 1
    while i < len(A) + 1:
        j = 1
        while j < len(B) + 1:
            D[i][j] = int(min(D[i][j - 1] + 1, D[i - 1][j] + 1, D[i - 1][j - 1] + my_diff(A[i - 1], B[j - 1])))
            j += 1
        i += 1

    return D[len(A)][len(B)]

A = str(input())
B = str(input())

print(editing_distance(A, B))
