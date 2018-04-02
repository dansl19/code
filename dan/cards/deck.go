package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string /*Here we are defining that deck will inherete all functions of type string ,
  deck will have everyhting the same but in additional I will create my custom functions*/

// newDeck function create new cards
func newDeck() deck { /* Always when some function should return some value it should be defined which type of value func should return ,
	since we are working with deck type (like string) we need to declare that this func will return value from type "deck"
	we don't put here reciever as we don't make some action with vars but we are going to create here new value and with excution of this func to retrun it*/
	cards := deck{} /* this creates new vars from type "deck" like the same you would create cards := string{}  https://www.udemy.com/go-the-complete-developers-guide/learn/v4/t/lecture/7797282?start=0 */

	cardSuite := []string{"Spades", "Diamond", "Hearts", "Clubs"} //here we create slice [] of type string
	cardValue := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuite { // here we itirate throuth slice of "cardSuite"
		for _, value := range cardValue { // here we itirate throuth slice of "cardValue"
			cards = append(cards, value+" of "+suit) // here we takes value from slice cardSuite and add value from slice cardValue
		}
	}
	return cards //here we retrun the final value
}

//Print function will print cards
func (d /*cards*/ deck) print() { //here we tell that print function will be avaible for all variables of type deck ,
	/* "d" this is reciever , Receiver - set up methods on Variable that we creates, essentially when this function is being excuted
	"cards" var will be placed in "d" */
	for i, card := range d /*cards*/ { //mi prisvaevaem kajdone znachenie card i potom raspechaetivaem ego
		fmt.Println(i, card)
	}
}

// in this func we are taking and spliting kolody na number of handize which will be transferd from main func
func deal(d deck, handSize int) (deck, deck) { //there are two vars which func expect to get and two retrun
	return d[:handSize], d[handSize:] // 1 all cards till num of "handSize" ,2 all cards after num of "handSize"
}

// this func takes "d" as reciever and stores everything from "d" to one string with Join function and delimiter ","
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// This function takes string from "toString" function and save it to file on a disk with permission 0666
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//This func reads the string from a file and puts it to "s" var with "," seperated values
func newDeckFromFile(filename string) deck { // here we pass to the func filename and expect return value from type deck since we want to work with this value with our custom func
	bs, err := ioutil.ReadFile(filename) //here we have two vars for two parameters that we will get from reading the file "byte slice" and error
	if err != nil {                      // here we check if the error is not null , if error is not empty we need to print the error and exit the code.

		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // Takes value from bs vars (string) and split it by "," and save it in "s"
	return deck(s)                      // return values in "s" variable
}
