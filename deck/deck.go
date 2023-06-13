package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings

// this thing means that a deck is a type which borrows all the behaviour of a string
type deck []string

// we will creat a nuew function of newdeck which will create or return a list of playing cards.Essentaily
// a array of strings.

// this means this function will return a value of type deck.

func newdeck() deck {

	//    we created a card variable of type deck and it will have all  card combinations , but it will be
	// very difficult  to write all 52 cards henc e we will play a trick. we will make two slices and iterate
	// over them and will add their combination in a new card
	card := deck{}

	//    we are not making it a type of deck , because deck will be used for actual cards or playing cards ,
	// this wil behave as as string.
	cardsuits := []string{"spades", "diamond", "heart", "clubs"}
	cardvalues := []string{"ace", "two", "three", "four"}

	// now we will iterate and put all the cominatons in card

	for _, suit := range cardsuits {
		for _, value := range cardvalues {
			card = append(card, value+" of "+suit)
		}
	}
	// returning  the deck
	return card
}

// we made a function this (d deck) is a reciever which we will talk about later and now we will use this func
// tion in sliceandloop.go to iterate over cards , we will look at it's syntax just wait
func (d deck) Print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

// THIS FUNCTION WILL is used to split the cards.
func (d deck) Deal(handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// we have to convert our data that is all combination of cards into byte so that we can use writefile function
// of os package because if we don't convert our data we will not be able to use the function as that function
// requires a arguement of data []byte this means this slice requires values in bytes. There are certain values
// in bytes which corresponds to a particular character for example byte 105 corresponds to i.

// Here we are going to convert our deck to slice of string and then we will make that slice of string into
// a string by mashing them all as one string then we will gonna convert that string to byte.

// WE ARE GOING TO MAKE A HELPER FUNCTION WHICH WILL CONVERT OUR DECK INTO A STRING
// We will joing our slice of string with the help of the function join which is in the pacakge strings.
// It takes two arguements in which one will take slice of string and other will take a seperate string
// which says how you want to get your string joined

func (d deck) Tostring() string {
	// made a slice of string after we call the  function via card which have all combinations this []string(d)
	// will have all the values as slice of string we didn't had to with any problem in this as deck is itslef
	// a type of slice of string, this will work.
	// []string(d)

	return strings.Join([]string(d), ",")
}

// We are making this function to save the list of all cards to a file on the local machine.The third arguement
// of the function writefile is used in case the function does not already exsists.0666 is the permission which
// means anyone can read and write this file

func (d deck) Savetofile(filename string) error {
	return os.WriteFile(filename, []byte(d.Tostring()), 0666)
}

// After we wrote the file then we read the content from the file that is saved on our local machine
// after we read the file we gonnna get back a slice of string in byte.
// THIS FUNCTION IS MADE TO READ THE FILE.

func (d deck) Newdeckfromfile(filename string) deck {
	// done this because readfile function returns two values a byte slice and a error if occurs
	// if something went wrong while reading the file this error will be populated.And if everything
	// went correct value of error will be nill.
	bs, error := os.ReadFile(filename)
	//  HENCE
	if error != nil {
		// we will use exit function to exit from the program after we print the error.If we pass zero in that
		// function it means there is no error in the program and if not zero it means there is some error, hence
		// exit
		fmt.Println("Error", error)
		os.Exit(1)
	}
	//  after we get the byte slice if no error occurs then we have to convert that byte slice into a actual
	// deck
	// converted byte slice into string. Now after we get the string we will seperate the string with , with the
	// help of the function split which will take two arguements one slice of string and other is sep string.
	// and will finally return the slice of string.

	// done this because it will return a slice of string.
	s := strings.Split(string(bs), ",")

	//  converted the slice of string into a actual deck and returned it.
	return deck(s)
}

// Now we are going to shuffle all the cards for that we will create a function shuffle.There is not any special
// function to shuffle the cards we have to do it manually.
// There is Rand package and inside that pacakge there is a function called intn which take one arguement as int n
// and generates a random number between o to n .

func (d deck) Shuffle() {

	// here we didn't use anything in for i, as we don't want anything with that I want only index i in this case.
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		// this is basically a fancy code to swap the numbers
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
