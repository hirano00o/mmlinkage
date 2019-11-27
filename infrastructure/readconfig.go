package infrastructure

import (
  "github.com/go-yaml/yaml"
  "io/ioutil"
  "mantis/domain"
)

func ReadConfig(file string) (config domain.MantisConfig, err error) {
  buf, err := ioutil.ReadFile(file)
  if err != nil {
    return
  }

  config, err = readYaml(buf)
  if err != nil {
    return
  }
  return
}

func readYaml(fb []byte) (cnf domain.MantisConfig, err error) {
  cnfs := make([]domain.MantisConfig, 100, 100)
  err = yaml.Unmarshal(fb, &cnfs)
  if err != nil {
    return
  }
  cnf.Projects = cnfs[0].Projects
  cnf.Mantis = cnfs[1].Mantis
  cnf.Mattermost = cnfs[2].Mattermost
  cnf.Users = cnfs[3].Users
  return
}

