import os

LOCAL = os.environ.get("REMOTE_JUDGE", "false") != "true"

if LOCAL:

    class Node:
        def __init__(self, value, left=None, right=None):
            self.value = value
            self.right = right
            self.left = left


def deep(root) -> int:

    if root is None:
        return 0

    return max(deep(root.left), deep(root.right)) + 1


def state(root) -> bool:
    if root is None:
        return True

    l = state(root.left)
    r = state(root.right)

    return l and r and abs(deep(root.right) - deep(root.left)) <= 1


def solution(root) -> bool:
    return state(root)


def test():

    node6 = Node(12)

    node4 = Node(4, node6, None)
    node5 = Node(8)

    node1 = Node(2)
    node2 = Node(7, node4, node5)
    node3 = Node(0, node1, node2)

    assert not solution(node3)


if __name__ == "__main__":
    test()
