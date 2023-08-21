# Cache
The cache directory includes 

- [Vulnerability Database][trivy-db][^1]
- [Java Index Database][trivy-java-db][^2]
- [Misconfiguration Policies][misconf-policies][^3]
- Cache of previous scans.
 
The cache option is common to all scanners.

## Clear Caches
The `--clear-cache` option removes caches.

**The scan is not performed.**

```
$ trivy image --clear-cache
```

<details>
<summary>Result</summary>

```
2019-11-15T15:13:26.209+0200    INFO    Reopening vulnerability DB
2019-11-15T15:13:26.209+0200    INFO    Removing image caches...
```

</details>

## Cache Directory
Specify where the cache is stored with `--cache-dir`.

```
$ trivy --cache-dir /tmp/trivy/ image python:3.4-alpine3.9
```

## Cache Backend
!!! warning "EXPERIMENTAL"
    This feature might change without preserving backwards compatibility.

Trivy supports local filesystem and Redis as the cache backend. This option is useful especially for client/server mode.

Two options:

- `fs`
    - the cache path can be specified by `--cache-dir`
- `redis://`
    - `redis://[HOST]:[PORT]`
    - TTL can be configured via `--cache-ttl`

```
$ trivy server --cache-backend redis://localhost:6379
```

If you want to use TLS with Redis, you can enable it by specifying the `--redis-tls` flag.

```shell
$ trivy server --cache-backend redis://localhost:6379 --redis-tls
```

Trivy also supports for connecting to Redis with your certificates.
You need to specify `--redis-ca` , `--redis-cert` , and `--redis-key` options.

```
$ trivy server --cache-backend redis://localhost:6379 \
  --redis-ca /path/to/ca-cert.pem \
  --redis-cert /path/to/cert.pem \
  --redis-key /path/to/key.pem
```

[trivy-db]: ./db.md#vulnerability-database
[trivy-java-db]: ./db.md#java-index-database
[misconf-policies]: ../scanner/misconfiguration/policy/builtin.md

[^1]: Downloaded when scanning for vulnerabilities
[^2]: Downloaded when scanning `jar/war/par/ear` files
[^3]: Downloaded when scanning for misconfigurations