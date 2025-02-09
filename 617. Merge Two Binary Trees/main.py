# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def mergeTrees(self, root1, root2):

        if root1 == None and root2 == None:
            return None

        if root1 != None and root2 == None:
            root2 = TreeNode()
        elif root1 == None and root2 != None:
            root1 = TreeNode()    
        
        root1.val += root2.val

        root1.left = self.mergeTrees(root1.left, root2.left)
        root1.right = self.mergeTrees(root1.right, root2.right)

        return root1


def main():

    root2 = TreeNode(9)

    n1 = TreeNode(2)
    n2 = TreeNode(3)
    root1 = TreeNode(1, n1, n2)

    s = Solution()
    s.mergeTrees(root1, root2)


if __name__ == "__main__":
    main()
