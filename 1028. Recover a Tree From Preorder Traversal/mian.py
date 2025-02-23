class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def recoverFromPreorder(self, traversal: str):
        mstack = list()
        traversal = traversal.split("-")

        mnode_level = 1
        for dval in traversal:
            if dval == str():
                mnode_level += 1
            else:
                mstacklen = len(mstack)
                node = TreeNode(int(dval))
                mstack.append((mnode_level, node))
                mnode_level = 1

        _, root = mstack[0]
        dstack = [root]

        current_index = 1
        while current_index < len(mstack):
            mlevel, mnode = mstack[current_index]
            lastnode = dstack[-1]

            if len(dstack) == mlevel:
                if lastnode.left == None:
                    lastnode.left = mnode
                elif lastnode.right == None:
                    lastnode.right = mnode

            elif len(dstack) < mlevel:
                dstack.append(
                    lastnode.right if lastnode.right != None else lastnode.left
                )
                current_index -= 1
            elif len(dstack) > mlevel:
                dstack = dstack[:mlevel]
                current_index -= 1

            current_index += 1
        return root


if __name__ == "__main__":
    s = Solution()

    t = s.recoverFromPreorder("1-401--349---90--88")

    def printTree(node: TreeNode):
        if node == None:
            return

        print(node.val, end=" ")
        printTree(node.left)
        printTree(node.right)

    printTree(t)
