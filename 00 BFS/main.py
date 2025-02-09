from collections import defaultdict, deque

[n, m] = [int(x) for x in input().split()]

graph = defaultdict(list)

for _ in range(m):
    [u, v] = [int(x) for x in input().split()]

    graph[u].append(v)
    graph[v].append(u)

r = int(input())

colors = ["white"] * (n + 1)


def bfs(graph, colors, v):
    planned = deque()

    planned.append(v)
    colors[v] = "gray"

    while len(planned) > 0:
        v = planned.popleft()
        print(v, end=" ")

        graph[v].sort()

        for child in graph[v]:
            if colors[child] == "white":
                colors[child] = "gray"
                planned.append(child)

    colors[v] = "black"


bfs(graph, colors, r)
