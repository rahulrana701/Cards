package deck


func Sliceandloop() {

	// card := []string{"go course is good",newcard()}
	// card := deck{"go course is good", newcard()}
	// card = append(card, "go course is important")
	// THIS WILL PRINT ALL VALUES IN CARD
	// fmt.Println(card);

	// will be using print method to iterate over the card.
	// card.Print()

	// FOR LOOP THIS LOOP WILL RUN TILL THE RANGE OF card.
	// for index, cards := range card {
	// 	fmt.Println(index, cards)
	// }
	// BUT HERE YOU WILL SAY THAT WE USE := ONCE WHEN WE DECLARE THE VARIABLE HENCE THE LOOP WILL RUN MANY TIME
	// S AND WHY WE ARE DECLARING AGAIN AND AGAIN IT IS NOT CORRECT , THIS SYNTAX IS CORRECT BECAUSE EVERY TIME
	// WE ARE USING FOR LOOP LIKE THIS WE ARE THRWOING AWAY THE PRVIOUS INDEX AND CARDS THE HAD BEEN DECLARED
	// HENCE WE ARE REINITALIZING OR REDECLARING THE VALUE

	card := newdeck()

    card.Shuffle();
	card.Print();



	// card.Newdeckfromfile("my_cards");
	// card.Print();

	// card.Savetofile("my_cards");

	// card.Tostring()
	// fmt.Println(card.Tostring());

	// hand, remainingcard:=card.Deal(5);
	// hand.Print();
	// remainingcard.Print();

	// card.Print();

}

// func newcard() string {
// 	return "the go course is amazing"
// }

// IF WE WANT TO RUN TWO FILES MEANS CODE OF ONE FILE INTO ANTOTHER THEN WRITE IN TERMINAL go run file1.go
// file2.go