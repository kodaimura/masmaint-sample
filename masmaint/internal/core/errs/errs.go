package errs

import (
	"strings"
	"encoding/json"
    "database/sql"
    "github.com/lib/pq"
    "github.com/go-sql-driver/mysql"
    "github.com/mattn/go-sqlite3"
    "masmaint/config"
)


func NewError(err error) error {
    if err == nil {
        return nil
    }

	if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
		return NewBadRequestError(jsonErr.Field, jsonErr.Type.String())
	}
	if _, ok := err.(*json.SyntaxError); ok {
		return NewBadRequestError("", "")
	}

    cf := config.GetConfig()
    dbDriver := cf.DBDriver

    switch dbDriver {
    case "postgres":
        return classifyPostgresError(err)
    case "mysql":
        return classifyMySQLError(err)
    case "sqlite3":
        return classifySQLite3Error(err)
    default:
        return NewUnexpectedError(err.Error())
    }
}

func classifyPostgresError(err error) error {
    if err == sql.ErrNoRows {
        return NewNotFoundError()
    }
    if pgErr, ok := err.(*pq.Error); ok {
        if pgErr.Code == "23505" {
            columnName := extractColumnNameFromPostgresError(pgErr)
            return NewUniqueConstraintError(columnName)
        }
    }
    return NewUnexpectedError(err.Error())
}

func classifyMySQLError(err error) error {
    if err == sql.ErrNoRows {
        return NewNotFoundError()
    }
    if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
        columnName := extractColumnNameFromMySQLError(mysqlErr)
        return NewUniqueConstraintError(columnName)
    }
    return NewUnexpectedError(err.Error())
}

func classifySQLite3Error(err error) error {
    if err == sql.ErrNoRows {
        return NewNotFoundError()
    }
    if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
        columnName := extractColumnNameFromSQLite3Error(sqliteErr)
        return NewUniqueConstraintError(columnName)
    }
    return NewUnexpectedError(err.Error())
}

func extractColumnNameFromPostgresError(pgErr *pq.Error) string {
    if pgErr.Detail != "" {
        parts := strings.Split(pgErr.Detail, "violates unique constraint")
        if len(parts) > 0 {
            return strings.TrimSpace(parts[0])
        }
    }
    return "unknown"
}

func extractColumnNameFromMySQLError(mysqlErr *mysql.MySQLError) string {
    if strings.Contains(mysqlErr.Message, "for key") {
        parts := strings.Split(mysqlErr.Message, "for key")
        if len(parts) > 1 {
            return strings.TrimSpace(parts[1])
        }
    }
    return "unknown"
}

func extractColumnNameFromSQLite3Error(sqliteErr sqlite3.Error) string {
    if strings.Contains(sqliteErr.Error(), "UNIQUE constraint failed") {
        parts := strings.Split(sqliteErr.Error(), ":")
        if len(parts) > 1 {
            return strings.TrimSpace(parts[1])
        }
    }
    return "unknown"
}