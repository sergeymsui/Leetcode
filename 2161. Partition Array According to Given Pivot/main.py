class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        left_arr = []
        pivot_count = 0
        right_arr = []

        for d in nums:

            if d == pivot:
                pivot_count += 1
            elif d < pivot:
                left_arr.append(d)
            else:
                right_arr.append(d)

        return left_arr + [pivot] * pivot_count + right_arr


if __name__ == "__main__":
    s = Solution()
    print(s.pivotArray([9, 12, 5, 10, 14, 3, 10], 10) == [9, 5, 3, 10, 10, 12, 14])
