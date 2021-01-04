# birthday-server

A simple insecure server, meant as a game for a birthday present. Use `python manage.py populate` to populate database
with the keys from `keys.csv`.

Solution can be found using the following steps:

1. Visit the `/` url.
2. Inspect elements and un-hide the hidden div. This is where you find clues for encryption algorithm as well as the
   secret to decrypt.
3. Under network calls, notice requests failing with `404`. This is where django lists the `runSQLiteQuery?q=<query>`
   url.
4. Use SQL injection to find the tables and query keys from encryption table. Query params must be URL encoded.
    - To list all tables in sqlite use `select * FROM sqlite_master WHERE type="table" ORDER BY name;`
    - All keys can be listed by visiting url `/runSQLiteQuery?q=select%20%2A%20FROM%20bday_key%3B`.
5. Factorize keys. Multiple keys may be factorized. Something along the lines of
   ```python
   for k1 in keys:
      for k2 in keys:
         p = gcd(k1, k2)
         if p != 1:
            q = k1 // p
            # Factorized key k1.
   ```
6. Try decrypting the message. Sample code in `python_encrypt.py`.
