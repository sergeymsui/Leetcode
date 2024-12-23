import sys
from collections import defaultdict, deque

sys.setrecursionlimit(2000)

[n, m] = [int(x) for x in input().split()]

graph = defaultdict(list)

for _ in range(m):
    [k, l] = [int(x) for x in input().split()]
    graph[k].append(l)

colors = ["white" for _ in range(n + 1)]

stack = deque()


def dfs(colors, graph, v, stack: deque):
    colors[v] = "gray"

    for c in graph[v]:
        if colors[c] == "white":
            dfs(colors, graph, c, stack)

    colors[v] = "black"
    stack.append(v)


while n > 0:
    if colors[n] == "white":
        dfs(colors, graph, n, stack)
    n -= 1

for el in stack:
    print(el, end=" ")
