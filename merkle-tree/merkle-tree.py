from hashlib import sha256
import sys
from typing import Any, List
from xmlrpc.client import Boolean


def sha256_hash(value):
    return sha256(value.encode('utf-8')).hexdigest()

# left and right are actually MerkleTree objects themselves


class MerkleTree:
    def __init__(self, value: str, left: Any, right: Any, isLeaf: Boolean = False) -> None:
        self.right = right
        self.left = left
        self.value = value
        self.isLeaf = isLeaf
        if self.isLeaf:
            self.hash = sha256_hash(value)
        else:
            combined_hash = None
            if left and right:
                combined_hash = sha256_hash(left.hash + right.hash)
            elif left:
                combined_hash = sha256_hash(left.hash)
            else:
                combined_hash = sha256_hash(right.hash)
            self.hash = combined_hash


def build_merklee_tree(transactions: List[str]) -> MerkleTree:
    # First for loop is to transform str values into initial merkle-tree nodes
    leaves: List[MerkleTree] = []
    for transaction in transactions:
        leaves.append(MerkleTree(transaction, None, None, isLeaf=True))

    current_array_of_nodes: List[MerkleTree] = leaves[:]
    while len(current_array_of_nodes) > 1:
        new_array_of_nodes: List[MerkleTree] = []
        for i in range(0, len(current_array_of_nodes), 2):
            first_node: MerkleTree = current_array_of_nodes[i]
            second_node: MerkleTree = current_array_of_nodes[i+1]
            new_merklee_node: MerkleTree = MerkleTree(
                '', first_node, second_node)
            new_array_of_nodes.append(new_merklee_node)
        current_array_of_nodes = new_array_of_nodes
    merklee_root = current_array_of_nodes[0]
    return merklee_root


def verify_merkle_tree(merkle_tree) -> Boolean:
    if merkle_tree is None:
        return True
    if merkle_tree.left is None and merkle_tree.right is None:
        return sha256_hash(merkle_tree.value) == merkle_tree.hash
    if merkle_tree.left is None:
        return sha256_hash(merkle_tree.right.hash) == merkle_tree.hash
    if merkle_tree.right is None:
        return sha256_hash(merkle_tree.left.hash) == merkle_tree.hash
    return merkle_tree.hash == sha256_hash(merkle_tree.left.hash + merkle_tree.right.hash) and verify_merkle_tree(merkle_tree.left) and verify_merkle_tree(merkle_tree.right)


transactions = ["trans_1", "trans_2", "trans_3", "trans_4"]
merkle_root = build_merklee_tree(transactions=transactions)
isValidMerkleeTree = verify_merkle_tree(merkle_root)
print("Verified merkle tree") if isValidMerkleeTree else print(
    "Failed to verify merkle tree")
