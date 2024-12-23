from collections import defaultdict

[n, m] = [int(x) for x in input().split()]

vertexes = defaultdict(list)

for _ in range(m):
    vertex = [int(x) for x in input().split()]
    vertexes[vertex[0]].append(vertex[1])

for i in range(1, n + 1):
    nodes = vertexes[i]
    nodes.sort()

    print(len(nodes), *nodes)
