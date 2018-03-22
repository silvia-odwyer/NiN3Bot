package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"math/rand"
	"strconv"
	"time"
	"strings"
	geopattern "github.com/pravj/geopattern"
	discord "github.com/bwmarrin/discordgo"
)

func main(){
	session, err := discord.New("Bot TOKEN")
		
	if err != nil {
		fmt.Println(err)
		return
	}
	session.AddHandler(newMessage)
	fmt.Println("NiN3Bot has connected :)")
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
	println := fmt.Println
	outputMessage := ""
	messageContent := m.Content
	if m.Author.Bot {
		return
	}

	switch {
	case strings.HasPrefix(messageContent, "9help"):
		outputMessage = `I'm glad you want to know more about me! ^^ 
		Since I'm called NiN3Bot, all my commands start with the prefix '9'. 
		My commands include: 9font, 9art, 9repeat, 9sayhi, 9predict, 9time.
		Examples: 
		'9font Hello there' 
		'9repeat NiN3Bot likes to repeat things.'`
	case strings.HasPrefix(messageContent, "9img"):
		// Doesn't send the image yet, it generates the SVG but I need to work on embedding the SVG somehow :thinking:
		alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		
		var randomPhrase = ""
		for i := 0; i < 100; i++ {
			randomNumber := rand.Intn(26)
			println(randomNumber)
			randomLetter := alphabet[randomNumber]
			randomPhrase += randomLetter 
		}
		
		fmt.Println(randomPhrase)
		outputMessage = "Generating image based on string: " + randomPhrase
		args := map[string]string{"generator": "squares", "phrase" : randomPhrase}
		generatedPattern := geopattern.Generate(args)
		println(generatedPattern)

	case strings.HasPrefix(messageContent, "9font"):
		messageContent = strings.TrimPrefix(messageContent, "9font ")
		circularFontMap := map[rune]string{
			'A':"Ⓐ", 'B':"Ⓑ", 'C':"Ⓒ", 'D':"Ⓓ", 'E':"Ⓔ", 'F':"Ⓕ", 'G':"Ⓖ", 'H':"Ⓗ", 'I':"Ⓘ", 'J':"Ⓙ", 'K':"Ⓚ", 
			'L':"Ⓛ", 'M':"Ⓜ", 'N':"Ⓝ", 'O':"Ⓞ", 'P':"Ⓟ", 'Q':"Ⓠ", 'R':"Ⓡ", 'S':"Ⓢ", 'T':"Ⓣ", 'U':"Ⓤ", 'V':"Ⓥ", 
			'W':"Ⓦ", 'X':"Ⓧ", 'Y':"Ⓨ", 'Z':"Ⓩ", 'a':"ⓐ", 'b':"ⓑ", 'c':"ⓒ", 'd':"ⓓ", 
			'e':"ⓔ", 'f':"ⓕ", 'g':"ⓖ", 'h':"ⓗ", 'i':"ⓘ", 'j':"ⓙ", 'k':"ⓚ", 'l':"ⓛ", 'm':"ⓜ", 'n':"ⓝ", 'o':"ⓞ", 
			'p':"ⓟ", 'q':"ⓠ", 'r':"ⓡ", 's':"ⓢ", 't':"ⓣ", 'u':"ⓤ", 'v':"ⓥ", 'w':"ⓦ", 'x':"ⓧ", 'y':"ⓨ", 'z':"ⓩ", ' ':" "}

		var convertedMessage = ""
		for _, c := range messageContent{
			var convertedCharacter = circularFontMap[c]
			convertedMessage += convertedCharacter
		}
		outputMessage = convertedMessage

		// for key, value := range circularFontMap{
		// }
			
	case strings.HasPrefix(messageContent, "9art"):
		asciiArtArray := [16]string{"♫ *´” ¸.•´¸.•*´¨) ¸.•*¨) ♫ ♪ ¸.•´ (¸.• ♫ ♪", "[̲̅$̲̅(̲̅5̲̅)̲̅$̲̅]", "(•_•)", "◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻", "(•_•) ( •_•)>⌐■-■ (⌐■_■)", "(¬_¬)", "ˁ˚ᴥ˚ˀ", "¸¸♬·¯·♩¸¸♪·¯·♫¸¸", "[̲̅$̲̅(̲̅ιοο̲̅)̲̅$̲̅]", "✈__✈", "¯(ツ)_/¯" , "( ‘-’)人(ﾟ_ﾟ )", "¯㋡_/¯", "{•̃̾_•̃̾}", "〴⋋_⋌〵", "UPLOADING VIRUS.EXE [████████████████] 99%"}
		var randomNumber = rand.Intn(16)
		var randomAsciiArt = asciiArtArray[randomNumber]
		println(randomAsciiArt)
		outputMessage = randomAsciiArt
	case strings.HasPrefix(messageContent, "9predict"):
		var randomNumber = rand.Intn(5)
		var a [5]string 
		a[0] = "Tomorrow will be your lucky day."
		a[1] = "Looks like good luck is on the way!"
		a[2] = "Next week is looking bright for you ^^ "
		a[3] =  "Watch out for a clue tomorrow."
		a[4] = "You should keep an eye out for something cool tomorrow."
		
		var randomPrediction = a[randomNumber]
		var stringOfNumber = strconv.Itoa(randomNumber)
		var finalPrediction = "Your lucky number tomorrow is " + stringOfNumber
		finalPrediction = randomPrediction + "\n" + finalPrediction
		println(finalPrediction)
		outputMessage = finalPrediction
		
	case strings.HasPrefix(messageContent, "9sayhi"):
		outputMessage = "Hello o/"
		
	case strings.HasPrefix(messageContent, "9repeat"):
		messageContent = strings.TrimPrefix(messageContent, "9repeat ")
		outputMessage = messageContent
	
	case strings.HasPrefix(messageContent, "9time"):
		var currentTime = time.Now()
		// t, _ := time.Parse(time.UnixDate, outputMessage)
		// println(t)
	
		// Yeah, I've got two ways of formatting the time, still have to decide which I prefer.
		var formattedTime = currentTime.Format(time.RFC3339)
		println(formattedTime)
		//outputMessage = formattedTime
			
		formattedTime2 := currentTime.Format("3:04 PM, Monday Jan 2 2006 MST")
		outputMessage = formattedTime2
	} 
	s.ChannelMessageSend(m.ChannelID, outputMessage)

}
