package gadget

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

type Recorder interface {
	Player
	Record()
}

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

// import "goCourse/src/gadget"

// func playList(device gadget.Player, songs []string) {
// 	for _, song := range songs {
// 		device.Play(song)
// 	}
// 	device.Stop()
// }

// func main() {
// 	var player gadget.Player = gadget.TapePlayer{}
// 	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
// 	playList(player, mixtape)
// }
