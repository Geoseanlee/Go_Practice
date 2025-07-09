package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?
func getMessageText(analytics *Analytics, msg Message) {
	analytics.MessagesTotal++
	if msg.Success {
		analytics.MessagesSucceeded++
	} else {
		analytics.MessagesFailed++
	}

}