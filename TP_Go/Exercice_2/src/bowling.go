package bowling

import "fmt"

type Frame struct {
	firstThrow  int
	secondThrow int
}

// First version, only add values
func GetScore(game []Frame) (int, error) {

	// Before, just checked !=10, now we can have ==10 or ==11 (if strike or spare as 10th throw)
	if len(game) != 10 && len(game) != 11 {
		return -1, fmt.Errorf("Wrong table size")
	}
	if len(game) == 11 { // Check if the 10th throw is a strike or a spare
		if game[len(game)-2].firstThrow + game[len(game)-2].secondThrow == 10 { // if spare OR strike
			if game[len(game)-2].firstThrow != 10 && game[len(game)-1].secondThrow != 0{	// if spare with bad construction
				return -1, fmt.Errorf("Wrong Frame construction")
			}
		} else {
			return -1, fmt.Errorf("Wrong table size")
		}
	}
	// Now, we know that the frame is well formed, we can get the score normaly
	score := 0
	for i := 0; i < len(game); i++ {
		if game[i].firstThrow < 0 || game[i].secondThrow < 0 || (game[i].firstThrow+game[i].secondThrow > 10) { // If negative value or > 10 sum
					return -1, fmt.Errorf("Wrong Frame construction")
		}
		score = score + game[i].firstThrow + game[i].secondThrow 
		if game[i].firstThrow == 10 && i+1 < len(game) { // If strike
			score = score + game[i+1].firstThrow
			if game[i+1].firstThrow == 10 && i+2 < len(game) { // If another strike
				score = score + game[i+2].firstThrow
			} else {
				score = score + game[i+1].secondThrow
			}
		} else if game[i].firstThrow+game[i].secondThrow == 10 && i+1 < len(game) { // If spare
			score = score + game[i+1].firstThrow
		}
	}
	return score, nil
}
