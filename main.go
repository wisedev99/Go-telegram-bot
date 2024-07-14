package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Initialize bot using API token
	bot, err := tgbotapi.NewBotAPI("5784978869:AAFnW1shgcv_zygvCRci3HzSTNsCnLvshrU")
	if err != nil {
		log.Fatal(err)
	}

	// Set bot debug mode
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up a channel to receive updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Process incoming messages
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Handle different commands
		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi there! Welcome to your bot.")
			bot.Send(msg)
		case "/help":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Here are the available commands:\n"+
				"/start - Start the bot\n"+
				"/help - Display this help message\n"+
				"You can also say 'hi' for a welcome message.")
			bot.Send(msg)
		case "hi":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Welcome to the bot.")
			bot.Send(msg)
	    case "/photo":
            // Call a function to handle the /photo command
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "i dont have a photo: ")
			bot.Send(msg)
            // sendRandomPhoto(bot, update.Message.Chat.ID)
        default:
			// Echo the received message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}
}


// func sendRandomPhoto(bot *tgbotapi.BotAPI, chatID int64) {
//     // Use current time to seed random number generator for variety
//     rand.Seed(time.Now().UnixNano())

//     // Make a request to Pexels API to fetch a random photo
//     resp, err := grequests.Get("https://api.pexels.com/v1/curated", &grequests.RequestOptions{
//         Headers: map[string]string{
//             "Authorization": "3umB3XVkJT2PXVyx8kBYmY3N1303fJNmCpDlS2rGPzQPwV0wxIUsP8oJ", // Replace with your Pexels API key
//         },
//         Params: map[string]string{
//             "per_page": "1", // Request one photo
//         },
//     })

//     if err != nil {
//         log.Println("Unable to make request to Pexels:", err)
//         return
//     }

//     defer resp.Close()

//     // Process the response
//     if resp.StatusCode != http.StatusOK {
//         log.Printf("Pexels API returned non-OK status code: %d", resp.StatusCode)
//         log.Println(resp.String())
//         return
//     }

//     // Parse the response JSON
//     var photos struct {
//         Photos []struct {
//             Src struct {
//                 Original string `json:"original"`
//             } `json:"src"`
//         } `json:"photos"`
//     }

//     err = resp.JSON(&photos)
//     if err != nil {
//         log.Println("Unable to parse Pexels response:", err)
//         return
//     }

//     // Ensure we have at least one photo
//     if len(photos.Photos) < 1 {
//         log.Println("No photos found in Pexels response")
//         return
//     }

//     // Get the original photo URL
//     photoURL := photos.Photos[0].Src.Original
//     log.Println("_____________________________________-------------------------------", photos.Photos[0].Src.Original)

//     // Ensure the URL is HTTPS (if necessary, depending on Pexels response)
//     if !strings.HasPrefix(photoURL, "https://") {
//         log.Println("Photo URL is not HTTPS")
//         return
//     }

//   photoBytes, err := ioutil.ReadFile(photoURL)
//     if err != nil {
//         log.Panic(err)
//     }

//     // Create a file reader
//     fileReader := tgbotapi.FileReader{
//         Name:   "photo.jpg", // Name of the file
//         Reader: bytes.NewReader(photoBytes),
//         Size:   int64(len(photoBytes)),
//     }

//     // Send the photo
//     msg := tgbotapi.NewPhotoUpload(chatID, fileReader)
//     _, err = bot.Send(msg)
//     if err != nil {
//         log.Panic(err)
//     }
// }
