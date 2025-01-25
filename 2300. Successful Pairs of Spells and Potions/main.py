class Solution:
    def successfulPairs(self, spells, potions, success: int):

        potions.sort()
        spells_array = [0] * len(spells)
        potions_length = len(potions)
        max_potion = potions[len(potions) - 1]

        for i, s in enumerate(spells):
            min_potion = ceil(success / s)
            if min_potion > max_potion:
                continue

            j = bisect_left(potions, min_potion)
            spells_array[i] += potions_length - j

        return spells_array


if __name__ == "__main__":
    spells = [5, 1, 3]
    potions = [1, 2, 3, 4, 5]
    success = 7

    s = Solution()
    print(s.successfulPairs(spells, potions, success))
