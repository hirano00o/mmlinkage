package infrastructure

import (
  "flag"
  "os"
  "os/signal"
  "strconv"
  "syscall"
  "time"
  "log"
  "mantis/interfaces/controller"
  "mantis/domain"
)

var (
  sig chan os.Signal
)

func Run() {
  var f = flag.String("f", "mantis.yaml", "Your Config File")
  flag.Parse()

  config, err := ReadConfig(*f)
  if err != nil {
    log.Fatalln("ERROR: can't read config.\n", err)
  }

  c := controller.NewController(NewDBHandler(), config)

  sig = make(chan os.Signal, 1)
  signal.Notify(sig,
    syscall.SIGHUP,
    syscall.SIGINT,
    syscall.SIGTERM,
    syscall.SIGQUIT,
  )
  defer signal.Stop(sig)

  for i, _ := range config.Projects {
    go ticker(c, i, config)
  }

  for {
    select {
    case s := <-sig:
      switch s {
      case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
        log.Println("INFO: Quit")
        return
      }
    }
  }
}

func ticker(c *controller.BugController, th int, config domain.MantisConfig) {
  tick := 1
  t := time.NewTicker(time.Duration(tick) * time.Minute)
  defer t.Stop()

  log.Println("INFO: running monitoring project is "+config.Projects[th].Project+" ID is "+strconv.Itoa(th))
  for {
    select {
    case <-t.C:
      bugs, err := c.Read(config.Projects[th], config.Users, tick)
      if err != nil {
        log.Println("WARN: ID = "+strconv.Itoa(th)+" : "+err.Error())
      }
      err = c.Post(config.Mattermost, bugs, config.Mantis, config.Users)
      if err != nil {
        log.Println("WARN: ID = "+strconv.Itoa(th)+" : "+err.Error())
      }
    case s := <-sig:
      switch s {
      case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
        return
      }
    }
  }
}
