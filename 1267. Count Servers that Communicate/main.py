class Solution:

    def countServers(self, grid) -> int:
        self.history = set()

        row = [0] * len(grid)
        col = [0] * len(grid[0])

        for i in range(0, len(grid)):
            for j in range(0, len(grid[0])):
                row[i] += grid[i][j]
                col[j] += grid[i][j]

        count = 0
        for i in range(0, len(grid)):
            for j in range(0, len(grid[0])):
                if grid[i][j] == 1 and (col[j] > 1 or row[i] > 1):
                    count += 1

        return count


def main():
    grid = [
        [0, 0, 1, 0, 1],
        [0, 1, 0, 1, 0],
        [0, 1, 1, 1, 0],
        [1, 0, 0, 1, 1],
        [0, 0, 1, 1, 0],
    ]

    s = Solution()
    print(s.countServers(grid))


if __name__ == "__main__":
    main()
