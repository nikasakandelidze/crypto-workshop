# Lamport signature algorithm
Lamport signature is one way of constructing digital signature. It is using random generator and hash functions very intesnively

## Pre-requisites ( Paragraph 1 )

Given 2 entiies:

- 256 bit hash function
- random generator

## Creating a private key ( Paragraph 2 )

-Generate 256 random number pairs ( 512 in total, Each number being 256 bits in size so totally we generate '256*2*256=1024 bits' ). This is the
private key.  
[(X X) (X X) (X X)...... X] <--- private key (256 random pairs ) 512 random number units

## Creating a public key:

- Hash, using hash function specified in Paragraph 1, each of 512 numbers ( 256 pairs ) from generated during the private key generation.  
  [X X X X X ...... X] => map(e=>hash(e)) => [ (hash(X), hash(X)), ...... hash(X)]

## Signing the message M

- Hash the message M using hash function specified in Paragraph 1
- For each bit in the hash based on the bit ( 1 or 0 ) pick one number from the pairs of the private key ( since len(hash(M)) = 256 and there are 256
  pairs in the private key everything lines up.) This process produces a sequence of 256 numbers. These numbers are the signature and are published with the message.

## Veryfing a signature

Let's say we now want to verify Message and it's signature we got: M and S.

- Hash a M using hash function specified in Parahraph 1
- For each bit in the hashed message, choose the according number from the 256 pairs in the public key that is publicly known to anyone.
  Basically this process of choosing hashes from the public key is the same as sender trying to choose correct number from the pair during signature creation.
- Now we should hash each number from S ( signature that was sent with the message itself ). If this hash result exactly matches the chosen numbers from public key on the previous step than signature is ok.


## Usage of python script
Python script contains example usage and implementation of Lamport signature algorithm by generating public and private keys, signing some message with private key and validating signature and message with public key.
