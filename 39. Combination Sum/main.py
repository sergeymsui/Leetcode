## !! https://leetcode.com/problems/combination-sum/description/


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        candidates = candidates.sort()

        weel = candidates[-1]

        if target - weel:
            return [weel]

        n = len(candidates) - 2

        while n >= 0:

            n -= 1

        return result


if __name__ == "__main__":
    s = Solution()
    print(s.combinationSum([2, 3, 6, 7], 7))
