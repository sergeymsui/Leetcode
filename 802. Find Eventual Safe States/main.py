class Solution:

    def isEndedDest(self, node, graph, history, trues) -> bool:

        if graph[node] == []:
            return True

        result = True
        for child in graph[node]:

            if child in trues:
                continue
            elif child in history:
                return False

            history.add(child)
            result = result and self.isEndedDest(child, graph, set(history), trues)

            if result:
                trues.add(child)

        return result

    def eventualSafeNodes(self, graph):
        result = list()
        trues = set()

        for i in range(0, len(graph)):
            if self.isEndedDest(i, graph, set(), trues):
                result.append(i)

        result.sort()
        return result


def main():

    cases = [
        {
            "graph": [[1, 2], [2, 3], [5], [0], [5], [], []],
            "result": [2, 4, 5, 6],
        },
        {"graph": [[1, 3, 4, 5], [], [], [], [], [2, 4]], "result": [0, 1, 2, 3, 4, 5]},
    ]

    for i, case in enumerate(cases):
        graph = case["graph"]
        result = case["result"]

        s = Solution()
        ans = s.eventualSafeNodes(graph)
        print(i, ans == result, ans)


if __name__ == "__main__":
    main()
