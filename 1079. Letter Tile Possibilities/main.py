class Solution:

    def buildTree(self, tray, tiles, used):

        if len(tray) > 0:
            self.mset.add(tray)

        for i, tile_example in enumerate(tiles):
            if used[i]:
                ucopy_array = used[:]
                ucopy_array[i] = False
                self.buildTree(tray + tile_example, tiles, ucopy_array)

    def numTilePossibilities(self, tiles: str) -> int:
        used = [True] * len(tiles)
        self.mset = set()

        self.buildTree("", tiles, used)
        return len(self.mset)



if __name__ == "__main__":
    s = Solution()
    print(s.numTilePossibilities("AAB"))
    print(s.numTilePossibilities("AAABBC"))

