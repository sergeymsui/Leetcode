import heapq
from collections import defaultdict


class Solution:
    def minimumDistance(self, n: int, edges, s, marked) -> int:
        smarked = set(marked)

        weight = [float("inf")] * (n + 1)
        weight[s] = 0

        graph = defaultdict(list)
        for u, v, w in edges:
            graph[u].append((v, w))

        min_heap = [(0, s)]

        while min_heap:
            distance, node = heapq.heappop(min_heap)

            if node in smarked:
                return weight[node]

            for next_node, w in graph[node]:
                new_dist = distance + w

                if new_dist < weight[next_node]:
                    weight[next_node] = new_dist
                    heapq.heappush(min_heap, (new_dist, next_node))

        return -1


if __name__ == "__main__":
    s = Solution()
    print(s.minimumDistance(4, [[0, 1, 1], [1, 2, 3], [2, 3, 2], [0, 3, 4]], 0, [2, 3]))
    print(
        s.minimumDistance(
            5, [[0, 1, 2], [0, 2, 4], [1, 3, 1], [2, 3, 3], [3, 4, 2]], 1, [0, 4]
        )
    )
    print(s.minimumDistance(2, [[0, 1, 9]], 0, [1]))

    print(
        s.minimumDistance(
            3,
            [
                [1, 2, 5],
                [0, 2, 6],
                [0, 1, 1],
                [0, 2, 4],
                [0, 2, 8],
                [2, 1, 8],
                [0, 2, 1],
            ],
            2,
            [0],
        )
    )
