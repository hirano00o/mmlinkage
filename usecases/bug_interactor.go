package usecases

import (
  "time"
  "errors"
  "strconv"
  "encoding/json"
  "mantis/domain"
)

type BugAdapter interface {
  FindByProjectAfterSpecifiedTime(string, time.Time) (domain.Bugs, error)
}

type MattermostAdapter interface {
  Post(string, []byte) error
}

type BugInteractor struct {
  Bug BugAdapter
  Mt MattermostAdapter
}

func (b *BugInteractor) Check(pj domain.Project, users domain.Users, tick int) (bugs domain.Bugs, err error) {
  checkTiming := time.Now().Add(time.Duration(-1 * tick) * time.Minute)
  bugs, err = b.Bug.FindByProjectAfterSpecifiedTime("%"+pj.Project+"%", checkTiming)
  if err != nil {
    return
  }
  bugs = statusFormat(bugs, pj.Statuses)
  return
}

func statusFormat(b domain.Bugs, status domain.Status) (bugs domain.Bugs) {
  for _, v := range b {
    switch v.Status {
    case "10":
      v.Status = status.S10
    case "20":
      v.Status = status.S20
    case "30":
      v.Status = status.S30
    case "40":
      v.Status = status.S40
    case "50":
      v.Status = status.S50
    case "60":
      v.Status = status.S60
    case "70":
      v.Status = status.S70
    case "80":
      v.Status = status.S80
    case "90":
      v.Status = status.S90
    }
    bugs = append(bugs, v)
  }
  return
}

func (b *BugInteractor) CreateMessage(mm domain.Mattermost, bugs domain.Bugs, mantis domain.Mantis, users domain.Users) (err error) {
  if bugs == nil {
    return errors.New("bugs is nil")
  }
  if len(bugs) == 0 {
    return nil
  }
  bugMessage, err := marshalMessage(bugs, mantis, users)
  if err != nil {
    return
  }
  message, err := json.Marshal(domain.MattermostMessage{
    Username: mm.Username,
    Text: bugMessage,
  })
  if err != nil {
    return
  }
  return b.Mt.Post(mm.URL, message)
}

func marshalMessage(bugs domain.Bugs, mantis domain.Mantis, users domain.Users) (messages string, err error){
  for _, v := range bugs {
    user, err := exchangeUserMantisToMattermost(v.Handler, users.MantisUser, users.MattermostUser)
    if err != nil {
      return "", err
    }
    message := "[" + strconv.Itoa(v.ID) + "(URL1)](" + mantis.FirstURL + "/view.php?id=" + strconv.Itoa(v.ID) + ") "
    if mantis.SecondURL != "" {
      message += "[" + strconv.Itoa(v.ID) + "(URL2)](" + mantis.SecondURL + "/view.php?id=" + strconv.Itoa(v.ID) + ") "
    }
    message += v.Category + "\t" + user + "\t" + v.Status + " **" + v.Summary + "**\n"
    messages += message
  }
  return
}

func exchangeUserMantisToMattermost(user string, mantisUsers, mattermostUsers []string) (string, error) {
  for i, v := range mantisUsers {
    if user == v {
      if len(mattermostUsers) > i {
        return "@"+mattermostUsers[i], nil
      } else {
        return "", errors.New("users index out of range")
      }
    }
  }
  return user, nil
}
