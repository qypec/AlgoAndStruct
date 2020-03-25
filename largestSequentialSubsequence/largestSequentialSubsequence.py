
# Дано целое число n и массив A[1…n] натуральных чисел. 
# Выведите максимальное k, для которого найдётся подпоследовательность
# 1 <= i1 < i2 < ... < ik <= n длины k,
# в которой каждый элемент делится на предыдущий.

def LSS(A, n):
    lengths = []
    maxLen = 1

    for i in range(n):
        lengths.append(1)
        for j in range(i):
            if (A[j] <= A[i]) and (A[i] % A[j] == 0) and (lengths[i] < lengths[j] + 1):
                lengths[i] = lengths[j] + 1
                if lengths[i] > maxLen:
                    maxLen = lengths[i]
    return maxLen


n = int(input())
A = list(map(int,input().split()))

print(LSS(A, n))
