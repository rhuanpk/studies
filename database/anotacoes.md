# SQL

## Sintaxe

- Select:
```sql
SELECT <coluna(s)> FROM <tabela(s)> [WHERE,GROUP BY, ORDER BY [AND, OR, NOT]];
```

- Operadores relacionais padrões:
	- Menor: `<`
	- Maior: `>`
	- Menor ou igual: `<=`
	- Maior ou igual: `>=`
	- Igual: `=`
	- Diferente: `<>`

## PostgreSQL

- Sync Serial Seq:
```sql
SELECT SETVAL(
    pg_get_serial_sequence('table_name', 'column_name'),
    (SELECT MAX(column_name) FROM table_name>)
)
```

- Upsert único:
```sql
WITH upsert AS (
    UPDATE table_name
    SET column1 = value1,
        column2 = value2
    WHERE column_verify = value_verify
    RETURNING *
)
INSERT INTO table_name (column1, column2, column_verify)
SELECT value1, value2, value_verify
WHERE NOT EXISTS (SELECT 1 FROM upsert);
```

- Upser múltiplos:
```sql
WITH upsert AS (
    UPDATE table_name AS t
    SET
        column1 = v.column1,
        column2 = v.column2
    FROM (
        VALUES
            ('value1_1', 'value2_1', 'value_verify1'),
            ('value1_2', 'value2_2', 'value_verify2'),
            ('value1_3', 'value2_3', 'value_verify3')
    ) AS v (column1, column2, column_verify)
    WHERE t.column_verify = v.column_verify
    RETURNING t.*
)
INSERT INTO table_name (column1, column2, column_verify)
SELECT v.column1, v.column2, v.column_verify
FROM (
    VALUES
        ('value1_1', 'value2_1', 'value_verify1'),
        ('value1_2', 'value2_2', 'value_verify2'),
        ('value1_3', 'value2_3', 'value_verify3')
) AS v (column1, column2, column_verify)
WHERE NOT EXISTS (SELECT 1 FROM upsert WHERE upsert.column_verify = v.column_verify);
```
