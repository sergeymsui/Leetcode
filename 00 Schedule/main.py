n = int(input())

shedule = []

for _ in range(n):
    [s, f] = [float(x) for x in input().split()]

    start = int(s) if int(s) == float(s) else float(s)
    finish = int(f) if int(f) == float(f) else float(f)

    shedule.append([start, finish])

sorted_shedule = []

sorted_shedule = sorted(shedule, key=lambda x: [x[1], x[0]])
shedule = [sorted_shedule[0]]

for i in range(1, len(sorted_shedule)):
    lesson = sorted_shedule[i]
    if shedule[-1][1] <= lesson[0]:
        shedule.append(lesson)

print(len(shedule))
for lesson in shedule:
    print(*lesson)
