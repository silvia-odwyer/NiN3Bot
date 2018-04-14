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
	session, err := discord.New("Bot TOKEN ID")
		
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
		fontMap := map[rune]string{}
		circularFontMap := map[rune]string{
			'A':"Ⓐ", 'B':"Ⓑ", 'C':"Ⓒ", 'D':"Ⓓ", 'E':"Ⓔ", 'F':"Ⓕ", 'G':"Ⓖ", 'H':"Ⓗ", 'I':"Ⓘ", 'J':"Ⓙ", 'K':"Ⓚ", 
			'L':"Ⓛ", 'M':"Ⓜ", 'N':"Ⓝ", 'O':"Ⓞ", 'P':"Ⓟ", 'Q':"Ⓠ", 'R':"Ⓡ", 'S':"Ⓢ", 'T':"Ⓣ", 'U':"Ⓤ", 'V':"Ⓥ", 
			'W':"Ⓦ", 'X':"Ⓧ", 'Y':"Ⓨ", 'Z':"Ⓩ", 'a':"ⓐ", 'b':"ⓑ", 'c':"ⓒ", 'd':"ⓓ", 
			'e':"ⓔ", 'f':"ⓕ", 'g':"ⓖ", 'h':"ⓗ", 'i':"ⓘ", 'j':"ⓙ", 'k':"ⓚ", 'l':"ⓛ", 'm':"ⓜ", 'n':"ⓝ", 'o':"ⓞ", 
			'p':"ⓟ", 'q':"ⓠ", 'r':"ⓡ", 's':"ⓢ", 't':"ⓣ", 'u':"ⓤ", 'v':"ⓥ", 'w':"ⓦ", 'x':"ⓧ", 'y':"ⓨ", 'z':"ⓩ", ' ':" "}
		monospaceFontMap := map[rune]string{'A':"𝙰", 'B':"𝙱", 'C':"𝙲", 'D':"𝙳", 'E':"𝙴", 'F':"𝙵", 'G':"𝙶", 'H':"𝙷", 'I':"𝙸", 'J':"𝙹", 'K':"𝙺", 'L':"𝙻", 'M':"𝙼", 'N':"𝙽", 'O':"𝙾", 'P':"𝙿", 'Q':"𝚀", 'R':"𝚁", 'S':"𝚂", 'T':"𝚃", 'U':"𝚄", 'V':"𝚅", 'W':"𝚆", 'X':"𝚇", 'Y':"𝚈", 'Z':"𝚉", 'a':"𝚊", 'b':"𝚋", 'c':"𝚌", 'd':"𝚍", 'e':"𝚎", 'f':"𝚏", 'g':"𝚐", 'h':"𝚑", 'i':"𝚒", 'j':"𝚓", 'k':"𝚔", 'l':"𝚕", 'm':"𝚖", 'n':"𝚗", 'o':"𝚘", 'p':"𝚙", 'q':"𝚚", 'r':"𝚛", 's':"𝚜", 't':"𝚝", 'u':"𝚞", 'v':"𝚟", 'w':"𝚠", 'x':"𝚡", 'y':"𝚢", 'z':"𝚣"}
		traditionalFontMap := map[rune]string{'A':"𝕬", 'B':"𝕭", 'C':"𝕮", 'D':"𝕯", 'E':"𝕰", 'F':"𝕱", 'G':"𝕲", 'H':"𝕳", 'I':"𝕴", 'J':"𝕵", 'K':"𝕶", 'L':"𝕷", 'M':"𝕸", 'N':"𝕹", 'O':"𝕺", 'P':"𝕻", 'Q':"𝕼", 'R':"𝕽", 'S':"𝕾", 'T':"𝕿", 'U':"𝖀", 'V':"𝖁", 'W':"𝖂", 'X':"𝖃", 'Y':"𝖄", 'Z':"𝖅", 'a':"𝖆", 'b':"𝖇", 'c':"𝖈", 'd':"𝖉", 'e':"𝖊", 'f':"𝖋", 'g':"𝖌", 'h':"𝖍", 'i':"𝖎", 'j':"𝖏", 'k':"𝖐", 'l':"𝖑", 'm':"𝖒", 'n':"𝖓", 'o':"𝖔", 'p':"𝖕", 'q':"𝖖", 'r':"𝖗", 's':"𝖘", 't':"𝖙", 'u':"𝖚", 'v':"𝖛", 'w':"𝖜", 'x':"𝖝", 'y':"𝖞", 'z':"𝖟"}

		messageContent = strings.TrimPrefix(messageContent, "9font ")
		switch{
		case strings.HasPrefix(messageContent, "c "):
			messageContent = strings.TrimPrefix(messageContent, "c ")
			fontMap = circularFontMap
		
		case strings.HasPrefix(messageContent, "m "):
			messageContent = strings.TrimPrefix(messageContent, "m ")

			fontMap = monospaceFontMap
		case strings.HasPrefix(messageContent, "t"):
			messageContent = strings.TrimPrefix(messageContent, "t ")
			fontMap = traditionalFontMap
		default:
			fontMap = circularFontMap
		}
		
		var convertedMessage = ""
		for _, c := range messageContent{
			var convertedCharacter = fontMap[c]
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
		predictionList := [...]string{"Tomorrow will be your lucky day.", "Looks like good luck is on the way!", "Next week is looking bright for you ^^", "Watch out for a clue tomorrow.", "You should keep an eye out for something cool tomorrow."}
		
		var randomPrediction = predictionList[randomNumber]
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
