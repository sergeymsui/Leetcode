class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def recursivebuild(self, preorder, postorder):
        if len(preorder) == 1:
            return TreeNode(preorder[0])

        if len(preorder) == 0:
            return None

        i = 1
        root = TreeNode(preorder[0])
        while i < len(preorder):
            j = 0
            while j < len(postorder):
                if preorder[i] == postorder[j]:
                    m = i
                    while m < len(preorder) and preorder[m] in postorder[: j + 1]:
                        m += 1

                    root.left = self.recursivebuild(preorder[i:m], postorder[: j + 1])
                    root.right = self.recursivebuild(preorder[m:], postorder[j + 1 :])
                    return root

                j += 1
            i += 1

    def constructFromPrePost(self, preorder, postorder):
        return self.recursivebuild(preorder, postorder)


if __name__ == "__main__":
    s = Solution()

    t = s.constructFromPrePost([1, 2, 4, 8, 9, 5, 3, 6, 7], [8, 9, 4, 5, 2, 6, 7, 3, 1])

    print(t)

    t = s.constructFromPrePost([2, 1], [1, 2])
