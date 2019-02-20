package twelve

import "fmt"

type nthDay struct {
	input int
	numName string
	count string
	object string

}

var nthDays = []nthDay {
	{1,"a", "first", "Patridge in a Pear Tree"},
	{2, "two","second", "Turtle Doves" },
	{3, "three","third", "French Hens"},
	{4, "four","fourth", "Calling Birds"},
	{5, "five","fifth", "Gold Rings"},
	{6,"six", "sixth", "Geese-a-Laying"},
	{7, "seven","seventh", "Swans-a-Swimming"},
	{8,"eight", "eighth", "Maids-a-Milking"},
	{9, "nine","ninth", "Ladies Dancing"},
	{10, "ten", "tenth", "Lords-a-Leaping"},
	{11, "eleven", "eleventh", "Pipers Piping"},
	{12, "twelve", "twelfth", "Drummers Drumming"},
}

func Song () string{

	song := ""

	// I want to construct song string. Do this by making individual verse strings.
	// want verse to be constructed by nThDays array.
	// array has int input (e.g 5), a name for the number (e.g five), a count name (e.g fifth), and the associated gift (Gold Rings)

	// construct gifts string by iterating back and appending to string each gift.

	for i := 1; i > 12; i ++ {

		song += Verse(i )

	}
	return song

}


func Verse(i int) string {

	verse:= fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", nthDays[i - 1].count, Gifts(i))

	return verse
}

func Gifts(day int) string {

	gifts := ""

	for i := day; i >= 0; i = i - 1 {
		punct := ", "
		if i == 0 {
			gifts += "and "
			punct = "."
		}
		fmt.Println(nthDays[i].numName + " " + nthDays[i].object)
		gifts += nthDays[i].numName + " " + nthDays[i].object + punct

	}

	return gifts
}