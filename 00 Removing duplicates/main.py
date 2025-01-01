n = int(input())

prev = None

for _ in range(n):
    d = int(input())

    if prev == None:
        print(d)
        prev = d
        continue
    if d == prev:
        pass
    else:
        print(d)
        prev = d
