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

- Manual delete cascade:
```sql
-- execute all together (select all then alt+x on dbeaver)
DROP PROCEDURE IF EXISTS delete_cascade;
DELIMITER $$
CREATE PROCEDURE delete_cascade()
BEGIN
	DECLARE table_name TEXT;
	DECLARE query_base TEXT;
	DECLARE done INT DEFAULT FALSE;

	DECLARE cur CURSOR FOR
		SELECT td.name FROM tables_delete td;

	DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

	SET query_base =
		"FROM parent p "
		"LEFT JOIN child_1 c1 ON c1.parent_id = p.id "
		"LEFT JOIN child_2 c2 ON c2.parent_id = p.id "
		"LEFT JOIN child_3 c3 ON c3.parent_id = p.id "
		"LEFT JOIN child_child_1 cc1 ON cc1.child_id = c3.id "
		"LEFT JOIN other_1 o1 ON o1.parent_id = p.id "
		"WHERE 1 = 1 "
		"	AND p.id = 42 "
	;

	CREATE TEMPORARY TABLE tables_delete (name TEXT);
	INSERT INTO tables_delete VALUES ('o1'), ('cc1'), ('c3'), ('c2'), ('c1'), ('p');

	OPEN cur;
	WHILE NOT done DO
		FETCH cur INTO table_name;
		IF NOT done THEN
			SET @stmt = CONCAT('DELETE ', table_name, ' ', query_base);
			PREPARE stmt FROM @stmt;
			EXECUTE stmt;
			DEALLOCATE PREPARE stmt;
		END IF;
	END WHILE;
	CLOSE cur;

	DROP TEMPORARY TABLE tables_delete;
END$$
DELIMITER ;

-- CALL delete_cascade();

-- execute one-by-one (select the query then ctrl+space on dbeaver)
DROP PROCEDURE IF EXISTS delete_cascade;

CREATE PROCEDURE delete_cascade()
BEGIN
	DECLARE table_name TEXT;
	DECLARE query_base TEXT;
	DECLARE done INT DEFAULT FALSE;

	DECLARE cur CURSOR FOR
		SELECT td.name FROM tables_delete td;

	DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

	SET query_base =
		"FROM parent p "
		"LEFT JOIN child_1 c1 ON c1.parent_id = p.id "
		"LEFT JOIN child_2 c2 ON c2.parent_id = p.id "
		"LEFT JOIN child_3 c3 ON c3.parent_id = p.id "
		"LEFT JOIN child_child_1 cc1 ON cc1.child_id = c3.id "
		"LEFT JOIN other_1 o1 ON o1.parent_id = p.id "
		"WHERE 1 = 1 "
		"	AND p.id = 42 "
	;

	CREATE TEMPORARY TABLE tables_delete (name TEXT);
	INSERT INTO tables_delete VALUES ('o1'), ('cc1'), ('c3'), ('c2'), ('c1'), ('p');

	OPEN cur;
	WHILE NOT done DO
		FETCH cur INTO table_name;
		IF NOT done THEN
			SET @stmt = CONCAT('DELETE ', table_name, ' ', query_base);
			PREPARE stmt FROM @stmt;
			EXECUTE stmt;
			DEALLOCATE PREPARE stmt;
		END IF;
	END WHILE;
	CLOSE cur;

	DROP TEMPORARY TABLE tables_delete;
END

-- CALL delete_cascade();
```

- Mostrar tabelas com nome:
```sql
SELECT
    TABLE_NAME, COLUMN_NAME
FROM
    INFORMATION_SCHEMA.COLUMNS
WHERE
    COLUMN_NAME LIKE '%<column>%'
    AND TABLE_SCHEMA = '<database>'
```
