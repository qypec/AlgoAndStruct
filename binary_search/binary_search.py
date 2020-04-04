#
#   В первой строке даны целое число n и массив A[1..n] из n различных натуральных чисел,
#   в порядке возрастания, во второй — целое число k и k натуральных чисел b1,...,bk.
#   Для каждого i от 1 до k необходимо вывести индекс j, для которого A[j] = bi, 
#   или -1−1, если такого j нет.
#

def binary_search(to_find, arr, l, r):
    pivot = l + int((r - l) / 2)
    if arr[pivot] == to_find:
        return pivot
    if r - l == 1:
        return -1
    if arr[pivot] > to_find:
        r = pivot
    else:
        l = pivot
    return binary_search(to_find, arr, l, r)


A = list(map(int,input().split()))
B = list(map(int,input().split()))

to_find = 0
result = []
i = 1
while i < B[0] + 1:
    to_find = B[i]
    result.append(binary_search(to_find, A, 1, len(A)))
    i += 1

print(*result)
