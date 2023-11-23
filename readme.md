# Engine

A javascript Engine base on [Goja](https://github.com/dop251/goja)

## Extensions

### Engine

Simple wrapper for goja.Runtime with extended abilities.

### Code

Simple wrapper for goja.Program

### Naming

1. use struct tag `js:"name"` to mapping fields to js property
2. use struct tag `js:"-"` to ignore export field
3. default strategy for both functions methods and fields
   + `UUUlllUll` => `uuUllUll`
   + `UlllUll` => `ullUll`
   + `XlllUll` => `llUll`
   + `XUllUll` => `UllUll`

### Resolver with `require`

Simple support for CommonJs, ES script also compiled as CJS script, inspire by k6

### Module System

1. Module: a simple global instance
2. InitializeModule:  a module instanced when register to an Engine
3. TopModule: a module register some top level function or objects
4. TypeDefined: a module contains `d.ts` data

### compiler module

built-in compiler for both typescript and javascript

[esbuild](https://github.com/evanw/esbuild)

### console

slog console or byte buffer console

## Modules

pluggable modules

### Os

Operating system api with environment control and simple file api

### sqlx

sqlx with `pgx` `mysql` and `sqlite` driver

[sqlx](https://github.com/jmoiron/sqlx),
[sqlite](https://github.com/glebarez/go-sqlite)
[mysql](https://github.com/go-sql-driver/mysql) *port*

### Excelize

excel reading or generate

[excelize](https://github.com/xuri/excelize/)

### crypto

dev

### jose

draft

### fetch

draft

### http

draft