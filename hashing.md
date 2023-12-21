# Hashing

## Definition

Hashing is a process of obtaining a hash for specified data.
A hash is a unique representation of the data it was calculated on.

## Hash Functions

A hash function is a function that takes data of arbitrary size and produces a fixed size hash.

## Some key features of Hashing

1. Original data cannot be restored from a hash. Thus, hashing is not encryption. (*It is often misunderstood that hashing is encryption, hope this clears it out*)

2. Certain data can have only one hash and the hash is unique.

3. Changing even one byte in the input data will result in a completely different hash.

## Use of Hashing in Blockchain

Hashing functions are widely used to check the consistency of data. Some software providers publish checksums in addition to a software package. 

1. In blockchain, hashing is used to guarantee the consistency of a block. 

2. The input data for a hashing algorithm contains the hash of the previous block, thus making it quite difficult to modify a block in the chain (*one has to recalculate its hash and hashes of all the blocks after it*).

# Hashcash

Bitcoin uses Hashcash (a Proof-of-Work algorithm that was initially developed to prevent email spam).

Hashcash is a brute force algorithm, hence computationally expensive.

Steps involved for implementing hashcash:

1. Take some publicly known data (in case of email, it’s receiver’s email address; in case of Bitcoin, it’s block headers).

2. Add a counter to it (starts from 0).

3. Get a hash of the ```data + counter``` combination.

4. Check that the hash meets certain "<u>requirements</u>" (_depending on the specific implementation_).
    - If it does, you're done.
    - Else, increase the counter and repeat steps 3 and 4.