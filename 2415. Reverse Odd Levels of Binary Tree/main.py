from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:

    def reverseOddLevels(self, root: Optional[TreeNode]) -> Optional[TreeNode]:

        if root.left == None:
            return root

        def changeValue(r1: TreeNode, r2: TreeNode, n: int):

            if r1 == None or r2 == None:
                return

            if n % 2 == 0:
                r1.val, r2.val = r2.val, r1.val

            changeValue(r1.left, r2.right, n + 1)
            changeValue(r1.right, r2.left, n + 1)

        changeValue(root.left, root.right, 0)

        return root


def show(root: Optional[TreeNode]):

    if root == None:
        return

    print(root.val)

    show(root.left)
    show(root.right)


if __name__ == "__main__":

    root = TreeNode(2)

    t1 = TreeNode(3)
    t2 = TreeNode(5)

    root.left = t1
    root.right = t2

    t3 = TreeNode(8)
    t4 = TreeNode(13)
    t5 = TreeNode(21)
    t6 = TreeNode(34)

    t1.left = t3
    t1.right = t4

    t2.left = t5
    t2.right = t6

    show(root)

    s = Solution()
    root = s.reverseOddLevels(root)

    print("--")

    show(root)
