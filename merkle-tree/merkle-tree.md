# Merkle-tree
Merkle tree is a data structure intensively used in blockchain technologies.
It is just a binary tree which follows next rules:
- each node has hash value stored in it, which varies dependning upon whether it's a leaf node or not a leaf node.
Leaf nodes in merkle tree are the only nodes that store actual values and hashes of these values.
Other nodes ( all non-leaf ones ) store hash of their (1 or 2) childrens' (combined) hashes.

# Usage
Merkle tree is very good mechanism to control data change and react to it, since we can compare original and to verify structures by simply comparing root hashes.
Since the root hash technically combines all the recursive children's hashes under the hood it is enough.
Since it is simply a chain/tree of hashes if even one value changes in the leaves
this change propagates to the root node of merkle tree and it won't become veirifaiable.

# Script
Python script implements merkle-tree using simple primitives. the code at first builds merkle tree with arbitrary transactions data and then verifies that hashes are correct.
