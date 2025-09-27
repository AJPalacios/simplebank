# SQLC Guide - Creando Nuevas Queries

Esta guía te explica cómo crear nuevas queries y funciones usando SQLC en este proyecto.

## Estructura de Queries

Cada query debe seguir este formato:

```sql
-- name: NombreFuncion :tipo_retorno
QUERY_SQL_AQUI;
```

## Tipos de Retorno

| Tipo | Descripción | Retorna en Go |
|------|-------------|---------------|
| `:one` | Una sola fila | `(Struct, error)` |
| `:many` | Múltiples filas | `([]Struct, error)` |
| `:exec` | Sin datos de retorno | `error` |
| `:execrows` | Número de filas afectadas | `(int64, error)` |

## Ejemplos Básicos

### Crear registro
```sql
-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;
```

### Obtener un registro
```sql
-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;
```

### Listar registros
```sql
-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;
```

### Actualizar
```sql
-- name: UpdateAccount :exec
UPDATE accounts 
SET balance = $2
WHERE id = $1;
```

### Eliminar
```sql
-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
```

## Parámetros

- Usa `$1`, `$2`, `$3`, etc. para parámetros
- Los tipos se infieren automáticamente del schema de la base de datos

## Comandos

Para generar el código Go después de crear/modificar queries:

```bash
# Usando makefile
make sqlc

# O directamente
sqlc generate
```

## Tips

1. **Nombres descriptivos**: Usa `GetAccountByID` en vez de `GetAccount`
2. **Usa LIMIT**: Siempre incluye LIMIT en queries que pueden retornar muchos resultados
3. **Prueba primero**: Prueba tus queries en PostgreSQL antes de agregarlas
4. **Un archivo por entidad**: Mantén queries relacionadas en el mismo archivo

## Archivos de Query

- `account.sql` - Operaciones de cuentas
- Crea nuevos archivos `.sql` para otras entidades según necesites

## Estructura del Proyecto

```
db/
├── query/           # Archivos .sql con queries (aquí agregas nuevas)
│   └── account.sql
└── sqlc/           # Código Go generado (NO EDITAR)
    ├── db.go
    ├── models.go
    └── account.sql.go
```