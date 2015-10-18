package main

import (
  "os"
  "fmt"

  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/aws"
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
          fmt.Println(err.Error())
          return
        }

        fmt.Println(resp)
      },
    },
    {
      Name:      "get",
      Aliases:     []string{"g"},
      Usage:     "get stack",
      Action: func(c *cli.Context) {
        stack_name := c.Args()[0]
        svc := cloudformation.New(nil)
        params := &cloudformation.GetTemplateInput{StackName: aws.String(stack_name)}
        resp, err := svc.GetTemplate(params)
        if err != nil {
          fmt.Println(err.Error())
          return
        }

        fmt.Println(*resp.TemplateBody)
      },
    },
  }

  app.Run(os.Args)
}
