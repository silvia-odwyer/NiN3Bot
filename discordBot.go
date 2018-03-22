package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"math/rand"
	"strconv"
	"time"
	discord "github.com/bwmarrin/discordgo" 
)

func main(){
	session, err := discord.New("Bot TOKEN")
		
	if err != nil {
		fmt.Println(err)
		return
	}
	session.AddHandler(newMessage)
	defer session.Close()

	if err = session.Open(); err != nil {
		fmt.Println(err);
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-sc
	
}

func newMessage(s *discord.Session, m *discord.MessageCreate){
	var outputMessage = ""
	if m.Author.Bot  {
		return
	}

	if m.Content == "9predict"{
		var randomNumber = rand.Intn(5)
		var a [5]string 
		a[0] = "Tomorrow will be your lucky day."
		a[1] = "Looks like good luck is on the way!"
		a[2] = "Hmm, sorry, but I don't seem to be in a great mood for making predictions right now. :/"
		a[3] =  "Watch out for a clue tomorrow."
		a[4] = "You should keep an eye out for something cool tomorrow."
	
		var randomPrediction = a[randomNumber]
		var stringOfNumber = strconv.Itoa(randomNumber)
		var finalPrediction = "Your lucky number tomorrow is " + stringOfNumber
		finalPrediction = randomPrediction + "\n" + finalPrediction
		fmt.Println(finalPrediction)
		outputMessage = finalPrediction
	}
	if m.Content == "9sayhi"{
		outputMessage = "Hello o/"
	}
	if m.Content == "9repeat"{
		outputMessage = m.Content
	}

	if m.Content == "9time"{
		outputMessage = time.Now()
	}
	s.ChannelMessageSend(m.ChannelID, outputMessage)

}
