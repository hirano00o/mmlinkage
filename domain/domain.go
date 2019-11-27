package domain

import (
  "time"
)

type Bugs []Bug

type Bug struct {
  ID int
  Summary string
  Category string
  Handler string
  Status string
  LastUpdated time.Time
}

type MantisConfig struct {
    Projects []Project `yaml:"projects"`
    Mantis Mantis `yaml:"mantis"`
    Mattermost Mattermost `yaml:"mattermost"`
    Users Users `yaml:"users"`
}

type Mantis struct {
    FirstURL string `yaml:"first_url"`
    SecondURL string `yaml:"second_url"`
}

type Mattermost struct {
    URL string `yaml:"url"`
    Username string `yaml:"username"`
}

type Users struct {
    MantisUser []string `yaml:"mapping_user_mantis"`
    MattermostUser []string `yaml:"mapping_user_mattermost"`
}

type Project struct {
    Project string `yaml:"project"`
    Statuses Status `yaml:"status"`
}

type Status struct {
      S10 string `yaml:"10"`
      S20 string `yaml:"20"`
      S30 string `yaml:"30"`
      S40 string `yaml:"40"`
      S50 string `yaml:"50"`
      S60 string `yaml:"60"`
      S70 string `yaml:"70"`
      S80 string `yaml:"80"`
      S90 string `yaml:"90"`
}

type MattermostMessage struct {
  Username string `json:"username"`
  Text string `json:"text"`
}
