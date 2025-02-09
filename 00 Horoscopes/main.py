n = int(input())
n_arr = [int(x) for x in input().split()]

m = int(input())
m_arr = [int(x) for x in input().split()]

n_m_array = []

for _ in range(m + 1):
    n_m_array.append([0 for _ in range(n + 1)])

for i in range(1, m + 1):
    for j in range(1, n + 1):
        c = i - 1
        r = j - 1

        if m_arr[c] == n_arr[r]:
            n_m_array[i][j] = n_m_array[c][r] + 1
        else:
            n_m_array[i][j] = max(n_m_array[i][r], n_m_array[c][j])

answer_m = []
answer_n = []

i = m
j = n
while i > 0 and j > 0:
    c = i - 1
    r = j - 1

    if m_arr[c] == n_arr[r]:
        answer_m.append(i)
        answer_n.append(j)

        i = c
        j = r
    else:
        if n_m_array[i][j] == n_m_array[c][j]:
            i = c
        else:
            j = r


if len(answer_n) == 0:
    print(0)
else:
    answer_n.reverse()
    answer_m.reverse()

    print(len(answer_n))
    print(*answer_n)
    print(*answer_m)
