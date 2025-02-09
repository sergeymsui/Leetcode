class NumberContainers:

    def __init__(self):
        self.indexToNumber = {}
        self.numberToIndex = defaultdict(SortedSet)

    def change(self, index: int, number: int) -> None:

        if index in self.indexToNumber:
            lastNumber = self.indexToNumber[index]
            self.numberToIndex[lastNumber].remove(index)

        self.indexToNumber[index] = number
        self.numberToIndex[number].add(index)
        

    def find(self, number: int) -> int:

        mmset = self.numberToIndex[number]
        return mmset[0] if len(mmset) > 0 else -1
    