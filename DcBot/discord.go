package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token = "MTE4MzI1NTU5MzczODMxMzc5MA.G6XodZ.FftEWJm5AeFudcM3WUkLZSqEpFhRc2bK_Qo9_g"
var botChID = "343076685169557504"
var Dg *discordgo.Session

func initSession() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("DC oturumu acilamadi, Error :", err)
		l.Fatalf("DC oturumu acilamadi, Error :", err)
	}

	Dg = dg

}

func ConnectToDC() {
	initSession()

	Dg.AddHandler(messageCreate)
	Dg.AddHandler(runScrapper)
	Dg.Identify.Intents = discordgo.IntentGuildMessages

	fmt.Println("Bot Calisiyor. Cikmak icin ctrl+c basiniz")
	l.Println("RSSParserBot calisiyor")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	Dg.Close()
	l.Println("[INFO] Bot Durduruldu")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "!rssbot" {
		s.ChannelMessageSend(m.ChannelID, "Buradayimm!!!")
		l.Printf("[INFO] %s kullanicisi beni cagirdi", m.Author)
	}
}
func runScrapper(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!rssrust" {
		s.ChannelMessageSend(m.ChannelID, rustnews())
	}

}
