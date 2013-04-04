package main 

import (
	"fmt"
	// "bufio"
	// "os"
	//"strconv"
	//"math"
    //"log"
)

var gameElementsArray 	[9]string
var gameBoardArray 		[11][11]string

/*
	This function sets up the board, populating the gameBoardArray variable with the 
		appropriate elements for displaying the game-board to the user
*/
func setUpBoard() {
	var skipCounter			int 		// Tests the inefficiency of my for-loops
	
	// This for-loop sets up the empty space in the board, while avoiding the spaces
	//	where the game elements will be located
	// Idealy, the locations [1,5,9][1,5,9] will remain nil in the array
	for i:=0; i<11; i++ {
		for e:=0; e<11; e++ {
			if (i==1 && e==1) ||
			   (i==1 && e==5) ||
			   (i==1 && e==9) ||
			   (i==5 && e==1) ||
			   (i==5 && e==5) ||
			   (i==5 && e==9) ||
			   (i==9 && e==1) ||
			   (i==9 && e==5) ||
			   (i==9 && e==9) {
				skipCounter+=1
			} else {
				gameBoardArray[i][e] = " "
			}
		}

	}

	// This for-loop ov_ides the " " set in the first for-loop for the spaces that would
	// 	be appropriately formated with "-" or "|"
	for e:=0; e<11; e++ {
		if e==3 || e==7 {
			skipCounter+=1
		} else {
			gameBoardArray[3][e] = "-"
			gameBoardArray[7][e] = "-"
			gameBoardArray[e][3] = "|"
			gameBoardArray[e][7] = "|"
		}
	}

	// Couldn't think of a way to loop this very well, so I just wrote it out completely
	// This adds the "+" where it would be appropriate for formatting.
	// I thought of a way to condense this in the code above, but I'm too lazy to move it
	gameBoardArray[3][3] = "+"
	gameBoardArray[3][7] = "+"
	gameBoardArray[7][3] = "+"
	gameBoardArray[7][7] = "+"

	for i:=0; i<9; i++ {
		gameElementsArray[i] = " "
	}

	fmt.Println(skipCounter)
}// End Function setUpBoard

/*
	This function actually draws the game-board and displays it to the user.
	I tried to use a case statement at first, but then I realized I was trying to do this
		the hard way rather than the easy way, and I went back to the if with multiple checks
*/
func drawBoard() {
	elementCounter := 0
	for i:=0; i<11; i++ {
		for e:=0; e<11; e++ {
			if 	(i==1 && e==1) ||
			   	(i==1 && e==5) ||
			   	(i==1 && e==9) ||
			   	(i==5 && e==1) ||
				(i==5 && e==5) ||
				(i==5 && e==9) ||
				(i==9 && e==1) ||
				(i==9 && e==5) ||
				(i==9 && e==9) {
				fmt.Print(gameElementsArray[elementCounter])
				elementCounter++
			} else {
				fmt.Print(gameBoardArray[i][e])
			} // end if/else

		} // end for
		fmt.Printf("\n")

	} // end for
	fmt.Printf("\n")

} // end function drawBoard

/*
	This function handles a single turn for the player
*/
func userPlayerTurn(){
	for {
		fmt.Println("Where would you like to place an X? (1-9)")
		// input:= bufio.NewReader(os.Stdin)
		// userInput, _ := input.ReadString('\n');
		var userInput int
		fmt.Scan(&userInput)
		if (userInput == 1 || userInput == 2 || userInput == 3 || userInput == 4 ||
		   userInput == 5 || userInput == 6 || userInput == 7 || userInput == 8 || 
		   userInput == 9) && gameElementsArray[userInput-1] == " " {
			// fmt.Println("Please try again...")
			// fmt.Printf("\n")
			PlaceX(userInput)
			break
		} else {
			fmt.Println("Please try again...")
			fmt.Printf("\n")
			// PlaceX(userInput)
			// break
		}
	} // enf of for-loop
}

/*
	Player's turn, set's the X in the gameElementsArray
*/
func PlaceX(x int) {
	// temp, _ := strconv.ParseInt(x, 10, 0)
	gameElementsArray[x-1] = "X"
}

/*
	Computer player's turn
*/
func computerPlayerTurn() {
	for i:=0; i<9; i++ {
		if gameElementsArray[i] == " " {
			gameElementsArray[i] = "O"
			break
		}
	}
}

/*
	This function checks to see if we have a winner for the game, and then determines who the winner is if so
*/
func checkWinStatus() int {
	
	var winner string

	if (gameElementsArray[0] == gameElementsArray[1]) && (gameElementsArray[1] == gameElementsArray[2]) {
		winner = gameElementsArray[0]
	} else if (gameElementsArray[3] == gameElementsArray[4]) && (gameElementsArray[4] == gameElementsArray[5]) {
		winner = gameElementsArray[3]
	} else if (gameElementsArray[6] == gameElementsArray[7]) && (gameElementsArray[7] == gameElementsArray[8]) {
	   	winner = gameElementsArray[6]
	} else if (gameElementsArray[0] == gameElementsArray[3]) && (gameElementsArray[3] == gameElementsArray[6]) {
		winner = gameElementsArray[0]
	} else if (gameElementsArray[1] == gameElementsArray[4]) && (gameElementsArray[4] == gameElementsArray[7]) {
		winner = gameElementsArray[1]
	} else if (gameElementsArray[2] == gameElementsArray[5]) && (gameElementsArray[5] == gameElementsArray[8]) {
		winner = gameElementsArray[2]
	} else if (gameElementsArray[0] == gameElementsArray[4]) && (gameElementsArray[4] == gameElementsArray[8]) {
		winner = gameElementsArray[0]
	} else if (gameElementsArray[2] == gameElementsArray[4]) && (gameElementsArray[4] == gameElementsArray[6]) {
		winner = gameElementsArray[0]
	}

	if winner == "X" {
		fmt.Println("YOU WIN!!!11!1!")
		return 1
	} else if winner == "O" {
		fmt.Println("You lost... :( ... to a computer player with no strategy")
		return 2
	}
	return 0
}

/*
	I decided to make setting up a new game separate from the main() function so that playing
		multiple games was easier to implement rather than running the file multiple times
*/ 
func newGame() {
	setUpBoard()
	drawBoard()

	for {
		var winStatus int
		userPlayerTurn()
		drawBoard()
		winStatus = checkWinStatus()
		if winStatus == 1 || winStatus == 2 {
			break
		}
		computerPlayerTurn()
		drawBoard()
		winStatus = checkWinStatus()
		if winStatus == 1|| winStatus == 2 {
			break
		}
	}
	
}

/*
	Our main function which will set us up the game
*/
func main() {
	
	for {
		newGame()
		fmt.Printf("\n \n \n")
		fmt.Println("Would you like to play again? (Y/N)")
		// input:= bufio.NewReader(os.Stdin)
		// userInput, _ := input.ReadString('\n');
		var userInput string
		fmt.Scan(&userInput)
		if userInput == "Y" || userInput == "y"|| userInput == "yes"|| userInput == "Yes" {

		} else if userInput == "N"|| userInput == "n"|| userInput == "no"|| userInput == "No" {
			break
		} else {
			fmt.Println("Input not recognized.")
			fmt.Println("Exiting Program")
			break
		}
	}
	fmt.Println("Thanks for playing!")
}