n = int(input())

vector = []

for _ in range(n):
    d = int(input())
    vector.append(d)


max_len = 0

i = 0
j = 0
while i < len(vector):
    if vector[i] == 1:

        j = i
        while j < len(vector) and vector[j] == 1:
            max_len = max(max_len, j - i + 1)
            j += 1

        i = j

    i += 1

print(max_len)
