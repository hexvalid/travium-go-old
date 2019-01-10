package main

import (
	"fmt"
	"github.com/hexvalid/travium-go/mari"
)

func main() {

	serverListesi, _ := mari.GetGameWorlds("https://www.travian.com/tr")

	fmt.Println("Toplam alınan server sayısı", len(serverListesi))

	for _, server := range serverListesi {
		fmt.Println("---------------------------------------")

		fmt.Println("Server UUIDsi", server.UUID)
		fmt.Println("Server URLsi", server.URL)
		fmt.Println("Server kayıt için key gerektiryormu?", server.RegistrationKeyRequired)
		fmt.Println("Server kısayolu: ", server.Shortcut)
		fmt.Println("Servera kayıt açık mı?: ", server.IsRegistrationOpen())
		fmt.Println("Server durumu: ", server.Status)
		fmt.Println("Başlama zmanı", server.StartTime())

		fmt.Println("---------------------------------------")

	}

}
