from collections import defaultdict


class Solution:
    def deepFind(self, vertex, m_dict, colors, target):
        f = False

        if vertex in m_dict and colors[vertex] == 0:
            colors[vertex] = 1
            if target in m_dict[vertex]:
                return True

            for v in m_dict[vertex]:
                f = f or self.deepFind(v, m_dict, colors, target)

        return f

    def checkIfPrerequisite(self, numCourses, prerequisites, queries):

        m_dict = defaultdict(list)
        for [u, v] in prerequisites:
            m_dict[u].append(v)

        result = list()
        for [u, v] in queries:
            f = False
            colors = [0] * numCourses

            if u in m_dict:
                f = f or self.deepFind(u, m_dict, colors, v)
            result.append(f)

        return result


def main():
    numCourses = 3
    prerequisites = [[1, 2], [1, 0], [2, 0]]
    queries = [[1, 0], [1, 2]]

    s = Solution()
    print(s.checkIfPrerequisite(numCourses, prerequisites, queries))


if __name__ == "__main__":
    main()
