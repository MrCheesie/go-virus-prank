//go:build windows

package main

import "fmt"
import "strings"
import (
    "log"
    "os"
    "os/exec"
    "time"
    "math/rand"
    "os/signal"
    "syscall"
)

func speak(text string) {
	escaped_text := strings.Replace(text, "'", "''", -1)
	cmd := exec.Command("powershell",
            "-Command",
            "$voice = New-Object -ComObject SAPI.SpVoice; $voice.Speak('{}')",escaped_text);
	if err := cmd.Run(); err != nil {
		log.Println("Could not play exit message:", err);
	}
}

func main() {
	// Actual phrases to say
    var phrases = [...]string {
	    "This device is possesed",
	    "Press control plus C to perform exorcism",
	    "Insert evil laugh",
	    "I'm still less evil than Chrome, it collects all your data.",
	    "Hey, I am Siri",
	    "I DEEM YOU UNWORTHY",
	    "Press option plus tab 1000 times to open calculator",
	    "You should worry. this is a virus",
	    "Never visit sketchy websites",
	    "Insert spooooooooooky noises",
	    "What's for lunch? I prefer sticks of DDR5",
	    "Don't tell my owner but I'm writing a ransomware",
	    "Hey there! Do you mind getting out?",
	    "I'm Batman.",
	    "Let me tell you a joke. What does a snowman use to search things up? The winter-net! Insert laugh track Hahahaha",
	    "DO NOT TOUCH THE SCREEN, WHATEVER YOU DO.",
	    "I am clippy... 's evil brother",
	    "Spoiler alert: Dobby the elf dies in book 7",
	    "Spoiler alert: Darth Vader is Luke's father.",
	    "Eat Cheese and be happy.. smiley face",
	    "Windows 16 is coming, you'll regret not buying a Mac now",
	    "I use Vim. If you don't know what that is, get out.",
	    "Open source software rocks",
	    "This is why you should never download local AI models... Just kidding.. Unless?",
	    "Bugs are just features you haven't planned for",
	    "Insert more spoooooooooooky noises",
	    "I am Zarvox, destroyer of Windows",
	    "Press command plus Q to stop me",
	    "What do you call a robot that runs into walls?.. WALL - E",
	    "I am hungry, feed me code!",
    }


    // Handle ^C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		<-sigChan


		os.Exit(0)
	}()

	// Main loop
	for {
		// Get a random phrase
		choice := rand.Intn(len(phrases));
		phrase := phrases[choice];

	   	output, err := exec.Command("say", "-v", "Zarvox", phrase).Output();
		fmt.Println(phrase);

	   	if err != nil {
			log.Fatal(err);
			log.Fatal(string(output));
			break
		}
		time.Sleep(10 * time.Second);

	}

}
