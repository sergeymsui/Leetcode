class Solution:
    def countDays(self, days, meetings) -> int:
        n = len(meetings)
        meetings.sort(key=lambda x: x[0])
        delta = 0

        i = 0
        while i < n:
            j = i + 1
            while j < n and meetings[i][1] >= meetings[j][0]:
                meetings[i][1] = max(meetings[i][1], meetings[j][1])
                j += 1

            delta += meetings[i][1] - meetings[i][0] + 1
            i = j

        return days - delta


if __name__ == "__main__":
    s = Solution()
    print(s.countDays(10, [[5, 7], [1, 3], [9, 10]]))
