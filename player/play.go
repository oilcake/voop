package player

// type Unit[T any] T

// func PlayLibrary(p *Player, r Reader) {
// 	for {
// 		// choose folder
// 		element := r.What(r.Now())
// 		path, ok := element.(*string)
// 		if !ok {
// 			log.Fatal("type cast failed")
// 		}
// 		// preload set from folder
// 		set, err := library.NewSet(path, p.Transport)
// 		if err != nil {
// 			log.Fatal("cannot preload folder", err)
// 		}

// 		action := PlaySet(p, set)
// 		fmt.Println(action)
// 		// (don't forget to close everything)
// 		library.CloseSet(set)
// 		switch action {
// 		case "rnd":
// 			r.Random()
// 		case "next":
// 			r.Next()
// 		case "prev":
// 			r.Previous()
// 		case "stop":
// 			return
// 		}
// 	}
// }

// func PlaySet(p *Player, r Reader) (action string) {

// 	for {
// 		// play media
// 		element := r.What(r.Now())
// 		media, ok := element.(*clip.Media)
// 		if !ok {
// 			log.Fatal("type conversion failed")
// 		}
// 		action = PlayMedia(media, p) // until any keyboard action
// 		fmt.Println(action)
// 		switch action {
// 		case "rnd":
// 			r.Random()
// 		case "next":
// 			r.Next()
// 		case "prev":
// 			r.Previous()
// 		case "stop":
// 			return
// 		case "nextChapter":
// 			return "next"
// 		case "prevChapter":
// 			return "prev"
// 		case "randomChapter":
// 			return "rnd"
// 		}
// 	}

// }

// func ChooseRandomFile(path *string) (string, error) {

// 	files, err := ioutil.ReadDir(*path)
// 	if err != nil {
// 		return "", err
// 	}
// 	log.Println("files total", len(files))
// 	rand.Seed(time.Now().UnixNano())
// 	file := *path + "/" + files[rand.Intn(len(files)-1)].Name()
// 	fmt.Println()
// 	log.Printf("Playing file %v\n", file)
// 	return file, nil
// }
