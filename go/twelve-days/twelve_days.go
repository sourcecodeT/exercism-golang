package twelve

import (
	"bytes"
	"fmt"
)

var numerals = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var gifts = []string{"a Partridge in a Pear Tree.", "two Turtle Doves, ", "three French Hens, ", "four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ", "seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ", "ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, "}
var start = "On the %s day of Christmas my true love gave to me: "

func Song() string {
	var buffer bytes.Buffer
	for n := 1; n <= 12; n++ {
		buffer.WriteString(Verse(n))
		if n < 12 {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

func Verse(n int) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(start, numerals[n-1]))
	for i := n - 1; i >= 0; i-- {
		if n > 1 && i == 0 {
			buffer.WriteString("and ")
		}
		buffer.WriteString(gifts[i])
	}

	return buffer.String()
}
