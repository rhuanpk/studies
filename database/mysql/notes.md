# MySQL

- Encerramento de processos:
```sql
SHOW PROCESSLIST;
KILL <pid>;
```

- Definindo variáveis de configuração:
```sql
SHOW variables LIKE 'sql_mode';
SET @@GLOBAL.sql_mode='STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
```

- Definindo e usando Variáveis de consulta:
```sql
SET @email := '';
SELECT * FROM users WHERE email = @email;

SET @emails = JSON_ARRAY('user1@email.any','user2@email.any');
SELECT * FROM users WHERE JSON_CONTAINS(@emails, JSON_QUOTE(email));
```

- Insert ignore (pesado?):
```sql
INSERT INTO table_to_insert (field_1, field_2, field_3)
SELECT DISTINCT ? AS field_1, 42 AS field_2, ttf.column AS field_3
FROM table_to_fetch ttf
-- JOIN ... ON ...
LEFT JOIN table_to_insert tti ON tti.field_1 = ? AND tti.field_2 = 42 AND tti.field_3 = ttf.column
WHERE tti.id IS NULL
```

- Insert ignore (leve?):
```sql
INSERT INTO table_to_insert (field_1, field_2, field_3)
SELECT DISTINCT ? AS field_1, 42 AS field_2, t.column AS field_3
FROM (
	SELECT DISTINCT ttf.column
	FROM table_to_fetch ttf
	-- JOIN ... ON ...
) t
LEFT JOIN table_to_insert tti ON tti.field_1 = ? AND tti.field_2 = 42 AND tti.field_3 = t.column
WHERE tti.id IS NULL
```
