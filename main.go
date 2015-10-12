package main

import (
  "os"
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/service/cloudformation"
)

func main() {
  app := cli.NewApp()
  app.Name = "forecast"
  app.Usage = "cloudformation dry-run"

  app.Commands = []cli.Command{
    {
      Name:      "list",
      Aliases:     []string{"l"},
      Usage:     "list stacks",
      Action: func(c *cli.Context) {
        svc := cloudformation.New(nil)
        resp, err := svc.ListStacks(nil)
        if err != nil {
          println(err.Error())
          return
        }

        fmt.Println(resp)
      },
    },
  }

  app.Run(os.Args)
}
