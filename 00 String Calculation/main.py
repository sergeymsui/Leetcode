s = "0"


def findNums(s, i):
    n = i - 1
    m = i + 1
    while s[n - 1].isdigit():
        n -= 1
    while m < len(s) and s[m].isdigit():
        m += 1

    return n, m


i = 0
while i < len(s):
    if s[i] == "*":
        n, m = findNums(s, i)
        result = str(int(s[n:i]) * int(s[i + 1 : m]))
        s = s[:n] + result + s[m:]
        i = n
    i += 1

i = 0
while i < len(s):
    if s[i] == "+":
        n, m = findNums(s, i)
        result = str(int(s[n + 1 : i]) + int(s[i + 1 : m]))
        s = s[: n + 1] + result + s[m:]
        i = n
    i += 1

print(s)
