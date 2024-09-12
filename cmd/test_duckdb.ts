import {SQLX} from "go/sqlx"
const db=new SQLX("duckdb","?access_mode=READ_WRITE&allow_unsigned_extensions=true")
console.log(db)
console.log(db.query("SELECT current_setting('access_mode')"))
console.log(db.exec("CREATE TABLE users(name VARCHAR, age INTEGER, height FLOAT, awesome BOOLEAN, bday DATE)"))
console.log(db.exec("INSERT INTO users VALUES('marc', 99, 1.91, true, '1970-01-01')"))
console.log(db.exec("INSERT INTO users VALUES('macgyver', 70, 1.85, true, '1951-01-23')"))
console.log(db.query('	SELECT name, age, height, awesome, bday	FROM users	WHERE (name = :name1 OR name = :name2) AND age > :age AND awesome = :awesome',	{
    name1:  "macgyver", name2:"marc", age:30, awesome:true,
}))
console.log(db.query("SELECT ('{\"duck\": 42}'::JSON->'$.duck')::INTEGER Val"))
console.log(db.query("SELECT excel_text(1234567.897, 'h:mm AM/PM') AS timestamp"))

console.log(db.exec("CREATE TABLE documents (document_identifier VARCHAR, text_content VARCHAR, author VARCHAR, doc_version INTEGER)"))
console.log(db.exec("INSERT INTO documents VALUES ('doc1', 'The mallard is a dabbling duck that breeds throughout the temperate.', 'Hannes Mühleisen', 3), ('doc2', 'The cat is a domestic species of small carnivorous mammal.', 'Laurens Kuiper', 2);"))
console.log(db.exec("PRAGMA create_fts_index('documents', 'document_identifier', 'text_content', 'author')"))
console.log(db.query("SELECT document_identifier, text_content, score FROM ( SELECT *, fts_main_documents.match_bm25(  document_identifier, 'Muhleisen', fields := 'author') AS score FROM documents) sq WHERE score IS NOT NULL AND doc_version > 2 ORDER BY score DESC"))
console.log(db.query("SELECT document_identifier, text_content, score FROM ( SELECT *, fts_main_documents.match_bm25( document_identifier,'small cats') AS score  FROM documents ) sq WHERE score IS NOT NULL ORDER BY score DESC"))
console.table(db.query("FROM duckdb_extensions()"))
console.log(db.query("SELECT '127.0.0.1'::INET AS addr"))
db.close()
