class Solution:
    def magnificentSets(self, n, edges):
        adj = [[] for _ in range(n)]
        for node1, node2 in edges:
            adj[node1 - 1].append(node2 - 1)
            adj[node2 - 1].append(node1 - 1)

        color = [0] * n
        for node in range(n):
            if color[node] != 0:
                continue
            color[node] = 1
            queue = deque([node])

            while queue:
                curr = queue.popleft()
                for neighbor in adj[curr]:
                    if color[neighbor] == color[curr]:
                        return -1
                    if color[neighbor] == 0:
                        color[neighbor] = -color[curr]
                        queue.append(neighbor)

        def find_depth(root):
            visited = [False] * n
            visited[root] = True
            queue = deque([root])
            distance = 0
            while queue:
                for _ in range(len(queue)):
                    node = queue.popleft()
                    for neighbor in adj[node]:
                        if not visited[neighbor]:
                            visited[neighbor] = True
                            queue.append(neighbor)
                distance += 1
            return distance

        def find_max_depth(node):
            visited[node] = True
            max_dist = distances[node]
            for neighbor in adj[node]:
                if not visited[neighbor]:
                    max_dist = max(max_dist, find_max_depth(neighbor))
            return max_dist

        distances = [find_depth(node) for node in range(n)]
        visited = [False] * n
        groups = 0
        for node in range(n):
            if not visited[node]:
                groups += find_max_depth(node)

        return groups
