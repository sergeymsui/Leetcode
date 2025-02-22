class Solution:
    def smallestNumber(self, pattern: str) -> str:
        n = len(pattern)
        nums = [i for i in range(1, n + 2)]
        result = [0] * (n + 1)

        i = 0
        while i < n:
            middle_point = i
            right_point = i

            if pattern[i] == "I":
                while middle_point < n and pattern[middle_point] == "I":
                    middle_point += 1
                left_point = middle_point

                while right_point < left_point and result[right_point] == 0:
                    result[right_point] = nums[right_point]
                    right_point += 1

                if middle_point == n:
                    result[n] = nums[left_point]
            else:

                while middle_point < n and pattern[middle_point] == "D":
                    middle_point += 1
                left_point = middle_point

                while middle_point >= i and result[right_point] == 0:
                    result[right_point] = nums[middle_point]
                    middle_point -= 1
                    right_point += 1

                i = left_point

            i += 1

        return "".join(map(str, result))


if __name__ == "__main__":
    s = Solution()
    print(s.smallestNumber("IIIDIDDD"))
    print(s.smallestNumber("DDD"))
    print(s.smallestNumber("DDDIII"))
