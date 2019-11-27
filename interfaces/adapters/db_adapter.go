package adapters

import (
  "time"
  "mantis/domain"
)

type Handler interface {
  Query(string, ...interface{}) (Row, error)
}

type Row interface {
  Scan(...interface{}) error
  Err() error
  Next() bool
  Close() error
}

type BugAdapter struct {
  Handler
}

func (a *BugAdapter) FindByProjectAfterSpecifiedTime(category string, t time.Time) (bugs domain.Bugs, err error) {
  statement := `select bug.id, user.realname, bug.status, bug.category, bug.last_updated, bug.summary from mantis_bug_table as bug join mantis_user_table as user on bug.handler_id = user.id where category like ? and bug.last_updated > ? order by bug.last_updated desc`
  rows, err := a.Query(statement, category, t.Format("2006-01-02 15:04:05"))
  if err != nil {
    return
  }
  defer rows.Close()

  for rows.Next() {
    var id int
    var summary, category, handler, status string
    var lastupdated time.Time

    err = rows.Scan(&id, &handler, &status, &category, &lastupdated, &summary)
    if err != nil {
      return
    }
    bug := domain.Bug {
      ID: id,
      Summary: summary,
      Category: category,
      Handler: handler,
      Status: status,
      LastUpdated: lastupdated,
    }
    bugs = append(bugs, bug)
  }

  if err = rows.Err(); err != nil {
    return
  }

  return
}
