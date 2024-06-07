# Ascii Art FS

It's a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.
An optional arguments is the style of the output ASCII art. You can choose between standard, shadow and thinkerjoy.

## **[main.go](main.go)**

First, arguments are being handled, and saved correctly. Each line of the input string is turned into ASCII art character by character.

## **[src/getword.go](src/getword.go)**

The ```GetWord``` function takes a string as argument and applies the ```GetLetter``` function to each character. It returns a slice of string representing each line to display.

## **[src/getletter.go](src/getletter.go)**

The ```GetLetter``` function takes an int as argument repesenting the ASCII code of the character that needs to be turned into ASCII art. Its position in the banner file is calculated. A string is returned containing the ASCII art for the provided character.
