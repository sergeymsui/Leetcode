from collections import defaultdict

[n, m] = [int(x) for x in input().split()]

graph = defaultdict(list)

for _ in range(m):
    [u, v] = [int(x) for x in input().split()]

    graph[u].append(v)
    graph[v].append(u)

r = int(input())

stack = []
colors = ["white" for _ in range(n + 1)]


def dfs(stack, graph, colors, n):

    stack.append(n)

    while len(stack) > 0:

        v = stack.pop()

        if colors[v] == "white":
            colors[v] = "gray"
            stack.append(v)

            print(v, end=" ")

            childs = graph[v]
            childs.sort(reverse=True)

            for c in childs:
                if colors[c] == "white":
                    stack.append(c)

        elif colors[v] == "gray":
            colors[v] = "black"


dfs(stack, graph, colors, r)
