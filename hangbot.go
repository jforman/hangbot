package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

var (
	credentialsFile = flag.String("credentialsFile", "", "path to credentials file")
	project         = flag.String("project", "", "Google Cloud Project name")
	psTopic         = flag.String("topic", "", "Pub/Sub Topic for Hangouts Chat")
	psSubscription  = flag.String("subscription", "", "Pub/Sub Subscription for Hangouts Chat")
	incomingMessage *chat.Message
	responseMessage *chat.Message
)

func main() {
	flag.Parse()
	log.Println("Hangbot Starting.")
	log.Printf("Configuration: Credentials File: %s, Project: %s, Topic: %s, Subscription: %s.\n", *credentialsFile, *project, *psTopic, *psSubscription)

	// This seems like a hack, but some of the oauth libraries expect an environment variable
	// if you use the JSON file, as opposed to being able to specify the path
	// as part of client creation.
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", *credentialsFile)
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, *project, option.WithCredentialsFile(*credentialsFile))

	if err != nil {
		log.Fatalln("error creating newclient: %v.\n", err)
	}

	sub := client.Subscription(*psSubscription)

	httpClient, err := google.DefaultClient(oauth2.NoContext, "https://www.googleapis.com/auth/chat.bot")
	if err != nil {
		log.Fatalf("Error creating httpClient: %v.\n", err)
	}

	chatService, err := chat.New(httpClient)
	if err != nil {
		log.Fatalf("Error creating chatService: %v.\n", err)
	}

	sms := chat.NewSpacesMessagesService(chatService)

	cctx, _ := context.WithCancel(ctx)

	ok, err := sub.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking if subscription exists. Err: %v.", err)
	}
	if !ok {
		log.Fatalln("Checked if subscription exists. It doesn't.")
	}

	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("Received Message %s.\n", string(msg.Data))
		msg.Ack()

		err := json.Unmarshal(msg.Data, &incomingMessage)
		if err != nil {
			log.Fatalf("Unable to decode Chat Message JSON: %v.\n", err)
		}

		responseMessage = new(chat.Message)
		responseMessage.Text = "Generic response."

		response, err := sms.Create(incomingMessage.Space.Name, responseMessage).Do()
		if err != nil {
			log.Printf("There was an error sending a response back to Hangouts Chat: %v.\n", err)
		}
		log.Printf("Hangouts Response: %+v.\n", response)
	})

	log.Println("Hangbot Exiting.")
}
