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

- Upsert único:
```sql
WITH upsert AS (
    UPDATE sua_tabela
    SET coluna1 = valor1,
        coluna2 = valor2
    WHERE coluna_verificacao = valor_verificacao
    RETURNING *
)
INSERT INTO sua_tabela (coluna1, coluna2, coluna_verificacao)
SELECT valor1, valor2, valor_verificacao
WHERE NOT EXISTS (SELECT 1 FROM upsert);
```

- Upser múltiplos:
```sql
WITH upsert AS (
    UPDATE sua_tabela AS t
    SET
        coluna1 = v.coluna1,
        coluna2 = v.coluna2
    FROM (
        VALUES
            ('valor1_1', 'valor2_1', 'valor_verificacao_1'),
            ('valor1_2', 'valor2_2', 'valor_verificacao_2'),
            ('valor1_3', 'valor2_3', 'valor_verificacao_3')
    ) AS v (coluna1, coluna2, coluna_verificacao)
    WHERE t.coluna_verificacao = v.coluna_verificacao
    RETURNING t.*
)
INSERT INTO sua_tabela (coluna1, coluna2, coluna_verificacao)
SELECT v.coluna1, v.coluna2, v.coluna_verificacao
FROM (
    VALUES
        ('valor1_1', 'valor2_1', 'valor_verificacao_1'),
        ('valor1_2', 'valor2_2', 'valor_verificacao_2'),
        ('valor1_3', 'valor2_3', 'valor_verificacao_3')
) AS v (coluna1, coluna2, coluna_verificacao)
WHERE NOT EXISTS (SELECT 1 FROM upsert WHERE upsert.coluna_verificacao = v.coluna_verificacao);
```
