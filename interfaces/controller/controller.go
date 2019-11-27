package controller

import (
  "mantis/usecases"
  "mantis/interfaces/adapters"
  "mantis/domain"
)

type BugController struct {
  Interactor usecases.BugInteractor
}

func NewController(db adapters.Handler, config domain.MantisConfig) *BugController {
  return &BugController {
    Interactor: usecases.BugInteractor {
      Bug: &adapters.BugAdapter {
        Handler: db,
      },
      Mt: &adapters.MattermostAdapter{},
    },
  }
}

func (b *BugController) Read(pj domain.Project, users domain.Users, tick int) (bugs domain.Bugs, err error) {
  bugs, err = b.Interactor.Check(pj, users, tick)
  return
}

func (b *BugController) Post(mm domain.Mattermost, bugs domain.Bugs, mantis domain.Mantis, users domain.Users) (err error) {
  err = b.Interactor.CreateMessage(mm, bugs, mantis, users)
  return
}
