package infrastructure

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "mantis/interfaces/adapters"
)

type Handler struct {
  Conn *sql.DB
}

func NewDBHandler() adapters.Handler {
  db, err := sql.Open("mysql", "root:root@/bugtracker?parseTime=true")
  if err != nil {
    panic(err.Error())
  }
  handler := new(Handler)
  handler.Conn = db

  return handler
}

func (h *Handler) Query (statement string, args ...interface{}) (adapters.Row, error) {
  /*
  stmt, err := h.Conn.Prepare(statement)
  defer stmt.Close()
  if err != nil {
    return new(SqlRow), err
  }

  rows, err := stmt.Query(args)
  */
  rows, err := h.Conn.Query(statement, args)
  if err != nil {
    return new(SqlRow), err
  }
  row := new(SqlRow)
  row.Rows = rows

  return row, nil
}

type SqlRow struct {
  Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
  return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
  return r.Rows.Next()
}

func (r SqlRow) Close() error {
  return r.Rows.Close()
}

func (r SqlRow) Err() error {
  return r.Rows.Err()
}
