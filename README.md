# birthday-server
A simple insecure server, meant as a game for a birthday present.

Solution can be found using the following steps:
1. Visit the `/` url.
2. Inspect elements and un-hide the hidden div. This is where you find clues for encryption algorithm as well as the secret to decrypt.
3. Under network calls, notice requests failing with `404`. This is where django lists the `runSQLiteQuery?q=<query>` url.
4. Use SQL injection to find the tables and query keys from encryption table. Query params must be URL encoded.
5. TODO: instructions on finding key collision.