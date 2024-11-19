# PostgreSQL

- Sync Serial Seq:
```sql
SELECT SETVAL(
    pg_get_serial_sequence('table_name', 'column_name'),
    (SELECT MAX(column_name) FROM <table_name>)
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

- Saber se há registros nulos ou vazios:
```sql
DO $$
DECLARE
    tbl_name TEXT := 'your_table'; -- Define the table name as a variable
    cl_name TEXT; -- Column name
    dt_type TEXT; -- Data type
    result BOOLEAN;
BEGIN
    -- Iterate over all columns in the table
    FOR cl_name, dt_type IN
        SELECT column_name, data_type
        FROM information_schema.columns
        WHERE table_name = tbl_name
    LOOP
        -- Define the condition to check NULL or default value based on column type
        IF dt_type IN ('integer', 'smallint', 'bigint', 'numeric', 'real', 'double precision') THEN
            EXECUTE format('SELECT EXISTS (SELECT 1 FROM %I WHERE %I IS NULL OR %I = 0)', tbl_name, cl_name, cl_name) INTO result;
        ELSIF dt_type IN ('text', 'varchar', 'char') THEN
            EXECUTE format('SELECT EXISTS (SELECT 1 FROM %I WHERE %I IS NULL OR %I = '''')', tbl_name, cl_name, cl_name) INTO result;
        ELSIF dt_type = 'boolean' THEN
            EXECUTE format('SELECT EXISTS (SELECT 1 FROM %I WHERE %I IS NULL OR %I = FALSE)', tbl_name, cl_name, cl_name) INTO result;
        ELSIF dt_type IN ('json', 'jsonb') THEN
            EXECUTE format('SELECT EXISTS (SELECT 1 FROM %I WHERE %I IS NULL OR %I = ''{}''::jsonb)', tbl_name, cl_name, cl_name) INTO result;
        ELSE
            EXECUTE format('SELECT EXISTS (SELECT 1 FROM %I WHERE %I IS NULL)', tbl_name, cl_name) INTO result;
        END IF;

        -- Log the result for each column
        RAISE NOTICE 'Column: % / Type: % / Has NULL or Default: %', cl_name, dt_type, result;
    END LOOP;
END $$;
```
