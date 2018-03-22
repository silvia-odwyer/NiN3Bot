package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"math/rand"
	discord "github.com/bwmarrin/discordgo" 
)

func main(){
	session, err := discord.New("Bot NDI2MTAwNjQzODMzMzE1MzI4.DZRXcg.0N9oJ7YIlOJIoZCHAmq2YHVVpuk")
		
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

	if m.Author.Bot  {
		return
	}

	var randomNumber = rand.Intn(4)
	var a [4]string 
	a[0] = "Tomorrow will be your lucky day."
	a[1] = "3 will be your lucky number."
	a[2] = "Hmm, I don't seem to be in a great mood for making predictions right now."
	a[3] =  "Watch out for a clue tomorrow."

	var randomPrediction = a[randomNumber]

	s.ChannelMessageSend(m.ChannelID, randomPrediction)

}
