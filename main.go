package main 

import (
	"fmt"
	// "bufio"
	// "os"
	//"strconv"      I tried a few different things here that didn't end up working
	//"math"			I just never took them out
    //"log"
)

var gameElementsArray 	[9]string 			// holds the "X"s and "O"s
var gameBoardArray 		[11][11]string 		// holds the elements that make the board

/*
	This function sets up the board, populating the gameBoardArray variable with the 
		appropriate elements for displaying the game-board to the user
*/
func setUpBoard() {
	var skipCounter			int 		// Tests the inefficiency of my for-loops
	
	// This for-loop sets up the empty space in the board, while avoiding the spaces
	//	where the game elements will be located
	// Idealy, the locations [1,5,9][1,5,9] will remain nil in the array, for spacing
	//  purposes
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
			} // end if/else
		} // end for

	} // end for

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
		} // end if/else
	} // end for

	// Couldn't think of a way to loop this very well, so I just wrote it out completely
	// This adds the "+" where it would be appropriate for formatting.
	// I thought of a way to condense this in the code above, but I'm too lazy to move it
	gameBoardArray[3][3] = "+"
	gameBoardArray[3][7] = "+"
	gameBoardArray[7][3] = "+"
	gameBoardArray[7][7] = "+"

	for i:=0; i<9; i++ {
		gameElementsArray[i] = " " // Populates the array with spaces, reseting the game
	} // end for

	// fmt.Println(skipCounter)     This was used to check the inefficiency of my loops

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
				(i==5 && e==5) || 			// these are the positions of "X"s and "O"s
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
	This function handles a single turn for the player, taking in input from the user
*/
func userPlayerTurn(){
	for {
		fmt.Println("Where would you like to place an X? (1-9)")
		// input:= bufio.NewReader(os.Stdin)			this method didn't work very well,
		// userInput, _ := input.ReadString('\n');     	  so I had to find another
		var userInput int
		fmt.Scan(&userInput)
		if (userInput == 1 || userInput == 2 || userInput == 3 || userInput == 4 ||
		   userInput == 5 || userInput == 6 || userInput == 7 || userInput == 8 || 
		   userInput == 9) && gameElementsArray[userInput-1] == " " {
			PlaceX(userInput)
			break
		} else {
			fmt.Println("Please try again...")
			fmt.Printf("\n")
		}
	} // enf of for-loop
} // end function userPlayerTurn

/*
	Player's turn, set's the X in the gameElementsArray
		I wanted to keep this separate for simplicity's sake, as well as avoiding
		corrupted input from the user.
*/
func PlaceX(x int) {
	gameElementsArray[x-1] = "X"
} // end function PlaceX

/*
	Computer player's turn
	The strategy that the computer uses is pretty darn awful.
		Basically, he checks to see if there's an open spot,
		and then takes the first one
*/
func computerPlayerTurn() {
	for i:=0; i<9; i++ {
		if gameElementsArray[i] == " " {
			gameElementsArray[i] = "O"
			break
		} // end if
	}  // end for
}  // end function computerPlayerTurn

/*
	This function checks to see if we have a winner for the game,
		and then determines who it is.
	There are 8 different win states that I set up in an if/else chain
	I set this function up to return 0, 1, or 2.
		Returning 0 means that the game is not yet over
		Returning 1 means that you won
		Returning 2 means that the computer won, and you suck
*/
func checkWinStatus() int {
	
	var winner string

	if (gameElementsArray[0] == gameElementsArray[1]) && 
	   (gameElementsArray[1] == gameElementsArray[2]) {
		winner = gameElementsArray[0]
	} else if (gameElementsArray[3] == gameElementsArray[4]) && 
			  (gameElementsArray[4] == gameElementsArray[5]) {
		winner = gameElementsArray[3]
	} else if (gameElementsArray[6] == gameElementsArray[7]) && 
			  (gameElementsArray[7] == gameElementsArray[8]) {
	   	winner = gameElementsArray[6]
	} else if (gameElementsArray[0] == gameElementsArray[3]) && 
			  (gameElementsArray[3] == gameElementsArray[6]) {
		winner = gameElementsArray[0]
	} else if (gameElementsArray[1] == gameElementsArray[4]) && 
			  (gameElementsArray[4] == gameElementsArray[7]) {
		winner = gameElementsArray[1]
	} else if (gameElementsArray[2] == gameElementsArray[5]) && 
			  (gameElementsArray[5] == gameElementsArray[8]) {
		winner = gameElementsArray[2]
	} else if (gameElementsArray[0] == gameElementsArray[4]) && 
			  (gameElementsArray[4] == gameElementsArray[8]) {
		winner = gameElementsArray[0]
	} else if (gameElementsArray[2] == gameElementsArray[4]) && 
			  (gameElementsArray[4] == gameElementsArray[6]) {
		winner = gameElementsArray[2]
	} // end if/else

	if winner == "X" {
		fmt.Println("YOU WIN!!!11!1!")
		return 1
	} else if winner == "O" {
		fmt.Println("You lost... :( ... to a computer player with no strategy")
		return 2
	} // end if/else
	return 0
} // end  function checkWinStatus

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
	} // end game for-loop
	
} // end function newGame

/*
	Our main function which will set us up the game
*/
func main() {
	
	for {
		newGame()
		fmt.Printf("\n \n \n")
		fmt.Println("Would you like to play again? (Y/N)")
		var userInput string
		fmt.Scan(&userInput)
		if userInput == "Y" || userInput == "y"|| userInput == "yes"|| userInput == "Yes" {

		} else if userInput == "N"|| userInput == "n"|| userInput == "no"|| userInput == "No" {
			break
		} else {
			fmt.Println("Input not recognized.")
			fmt.Println("Exiting Program")
			break
		} // end if/else
	} // end for
	fmt.Println("Thanks for playing!")
} // end function main