// Google Hangouts Chat Bot providing updates from Nest thermostats.

package main

import (
  "flag"
  "fmt"
  "log"

  "cloud.google.com/go/pubsub"
  "golang.org/x/net/context"
  "google.golang.org/api/option"

)

var (
  credentialsFile = flag.String("credentialsFile", "", "path to credentials file")
  project = flag.String("project", "", "Google Cloud Project name")
  psTopic = flag.String("topic", "", "Pub/Sub Topic for Hangouts Chat")
  psSubscription = flag.String("subscription", "", "Pub/Sub Subscription for Hangouts Chat")
)

func main() {
  flag.Parse()
  fmt.Println("Hangbot Starting.")
  fmt.Printf("Configuration: Credentials File: %s, Project: %s, Topic: %s, Subscription: %s.\n", *credentialsFile, *project, *psTopic, *psSubscription)
  ctx := context.Background()

  client, err := pubsub.NewClient(ctx, *project, option.WithCredentialsFile(*credentialsFile))

  if err != nil {
    log.Fatalln("error creating newclient: %v.\n", err)
  }

  sub := client.Subscription(*psSubscription)

  cctx, _ := context.WithCancel(ctx)

  ok, err := sub.Exists(ctx)
  if err != nil {
    log.Fatalf("Error checking if subscription exists. Err: %v.", err)
  }
  if ok {
    log.Println("Subscription exists. Let's go.")
  } else {
    log.Fatalln("Checked if subscription exists. It doesn't.")
  }

  err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
    fmt.Printf("Message: Data: %s, More: %+v.\n", string(msg.Data), msg)
    msg.Ack()
  })

  // if err != context.Canceled {
  //   return err
  // }

  fmt.Println("Hangbot Exiting.")
}
