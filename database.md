In its essence, blockchain is a distributed database.

As of ***commit 6***, there’s no database in our implementation; instead, we create blocks every time we run the program and store them in memory. We cannot reuse a blockchain, we cannot share it with others, thus we need to store it on the disk.

We'll be using **BoltDB** for this application.

Quoting some part from the BoltDB's README on GitHub

> Bolt is a pure Go key/value store inspired by Howard Chu’s LMDB project. The goal of the project is to provide a simple, fast, and reliable database for projects that don’t require a full database server such as Postgres or MySQL.

> Since Bolt is meant to be used as such a low-level piece of functionality, simplicity is key. The API will be small and only focus on getting values and setting values. That's it.

Since we’ll store Go structs (Block, in particular) in it, we’ll need to serialize them (*i.e. implement a mechanism of converting a Go struct into a byte array and restoring it back from a byte array*). We'll use encoding/gob for this.