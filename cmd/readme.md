# engine

The cli binary for Engine

# builtin modules

| module      | comment                                                                |
|-------------|------------------------------------------------------------------------|
| go/esbuild  | [esbuild](https://github.com/evanw/esbuild), use with notify is better |
| go/big      | big number by golang.big                                               |
| go/buffer   | binary utils                                                           |
| go/codec    | common codec utils                                                     |
| go/compiler | esbuild compiler with Engine                                           |
| go/context  | context package                                                        |
| go/io       | io package                                                             |
| go/os       | os package                                                             |
| go/time     | time package                                                           |
| go/http     | http package                                                           |
| go/encoding | encoding package                                                       |

# external modules

| module      | comment                                                                     |
|-------------|-----------------------------------------------------------------------------|
| go/fetch    | fetch by go.http                                                            |
| go/fsNotify | [fsNotify](https://github.com/fsnotify/fsnotify)                            |
| go/sqlx     | [sqlx](https://github.com/jmoiron/sqlx) utils                               |
| go/pgx      | [pgx](https://github.com/jackc/pgx/v5) utils                                |
| go/duckdb   | [duckdb](https://github.com/marcboeker/go-duckdb) and driver (cgo required) |
| go/pug      | [pug/jade](https://github.com/Joker/jade)                                   |
| go/minify   | [minify](https://github.com/tdewolff/minify/v2)                             |
| go/excelize | [excelize](https://github.com/xuri/excelize/v2)                             |
| go/chrome   | [cdp](https://github.com/chromedp/chromedp)                                 |
| go/gse      | [gse](https://github.com/go-ego/gse), not really good one                   |

# build options

| flags       | modules                                       | comment                               |
|-------------|-----------------------------------------------|---------------------------------------|
| all         | all external modules except duckDB,chrome,gse ||
| sqlx        | sqlx without any driver                       ||
| sqlx,sqlite | sqlx with pure golang sqlite                  ||
| sqlx,duckdb | sqlx with duckdb                              | need plaform specific dynamic library |
| sqlx,pgx    | sqlx with pgx for postgreSQL                  |                                       |
| sqlx,mysql  | sqlx with mysql                               |                                       |
| chrome      | cdp                                           |                                       |
| excel       | excelize for access excel file                |                                       |
| fetch       | fetch implement by http.Client                |                                       |
| gse         | gse tokenizer                                 |                                       |
| minify      | css,html,svg,js minifier                      |                                       |
| notify      | fs notify                                     |                                       |

# build for full functions
```shell
CGO_LDFLAGS="-L.libs/win/"  go build -tags=duckdb_use_lib,chrome,ducdb,gse,all -o engine.exe -ldflags="-s -w" .
```