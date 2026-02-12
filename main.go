package main

import (
    "fmt"
    "os"
    "github.com/bwmarrin/discordgo"
    "strings"
)


var AlphabetMRC = map[rune]string{
    'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
    'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
    'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
    'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
    'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
    'Z': "--..",
    '0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
    '5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",
    '!': "-.-.--", '?': "..--..", '\'': ".----.", ' ': "",
}


var AlphabetFR = map[string]rune{
    ".-": 'A', "-...": 'B', "-.-.": 'C', "-..": 'D', ".": 'E',
    "..-.": 'F', "--.": 'G', "....": 'H', "..": 'I', ".---": 'J',
    "-.-": 'K', ".-..": 'L', "--": 'M', "-.": 'N', "---": 'O',
    ".--.": 'P', "--.-": 'Q', ".-.": 'R', "...": 'S', "-": 'T',
    "..-": 'U', "...-": 'V', ".--": 'W', "-..-": 'X', "-.--": 'Y',
    "--..": 'Z',
    "-----": '0', ".----": '1', "..---": '2', "...--": '3', "....-": '4',
    ".....": '5', "-....": '6', "--...": '7', "---..": '8', "----.": '9',
    "-.-.--": '!', "..--..": '?', ".----.": '\'',
}


func DecodeMessage(msg string, alpha map[string]rune) string {
    start := 0
    traduction := ""

    for {
        pos := strings.Index(msg[start:], "/")
        if pos == -1 {
            break
        }
        pos += start

        letter, ok := alpha[msg[start:pos]]
        if !ok {
            traduction = "symbole invalide"
            break
        }

        traduction += string(letter)
        start = pos + 1

        if start < len(msg) && msg[start] == '/' {
            traduction += " "
            start++
        }
    }
    return strings.ToUpper(traduction)
}

func CodeMessage(msg string, alpha map[rune]string) string {
	var (
		traduction string
		temp string
		err bool
	)
	for _, e := range strings.ToUpper(msg) {
		temp, err = alpha[e]
		if(!err) {
			traduction = "symbole invalide"
			break
		}
		traduction += temp + "/"
	}
	return traduction
}

func ConfirmMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        return
    }
    msg := m.Content
    var response string

    if strings.HasPrefix(msg, "!decode") {
        response = DecodeMessage(msg[8:], AlphabetFR)
    } else if strings.HasPrefix(msg, "!code") {
	response = CodeMessage(msg[6:], AlphabetMRC);
    }

    s.ChannelMessageSend(m.ChannelID, response)
}

func main() {
	token := os.Getenv("TOKEN_MORISS")
    dg, err := discordgo.New(token)
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }

    dg.AddHandler(ConfirmMessage)
    err = dg.Open()
    if err != nil {
        fmt.Println("Erreur ouverture:", err)
        return
    }

    fmt.Println("Bot en ligne. Appuyez sur CTRL+C pour quitter.")
    select {}
}
