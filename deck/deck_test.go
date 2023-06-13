package deck

import (
	"os"
	"testing"
)

// we want to test newdeck() function of deck;
func TestNewdeck(t *testing.T)  {
	d:=newdeck();

	// the lenght of deck should be equal to 16.
	if len(d) !=16 {
		// this Errorf function is writtn like this therefore if we pass any arguements to it , it somehow 
		// must be injected into the string we  wrote in that function
		t.Errorf("expected length of deck of 16 but got %v", len(d));
	}
}

func TestSavetofileandNewdeckfromfile(t *testing.T)  {
	// this is done to remove if the file with that name already exisits in our directory. This is to play safe 
	// as if there exsisted a file like this there would arise a problem in loading.
	os.Remove("_decktesting");

	deck:=newdeck();
	deck.Savetofile("_decktesting");
	loadeddeck:=deck.Newdeckfromfile("_decktesting");

	if len(loadeddeck) !=16 {
		t.Errorf("expected length of deck of 16 but got %v", len(loadeddeck))
	}

}