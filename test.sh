#!/bin/bash

INPUTS=(
    "hello"
    "HELLO"
    "HeLlo HuMaN"
    "1Hello 2There"
    "Hello\n\nThere"
    "{Hello & There #}"
    "hello There 1 to 2!"
    "MaD3IrA&LiSboN"
    "1a\"#FdwHywR&/()="
    "{|}~"
    "[\]^_ 'a"
    "RGB"
    ":;<=>?@"
    "\!\" #$%&'()*+,-./"
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    "abcdefghijklmnopqrstuvwxyz"
    "four LETTER"
    "fivel 55"
    "L $ * @"
    "to 2 ^@ tree"
)

GO_PROGRAM="go run ."

printf "\n \033[36m%s\033[0m" "////////////////////////TESTS///////////////////////////"
printf "\n \033[35m◼\033[0m : Input \033[32m◼\033[0m : Output \n"

for input in "${INPUTS[@]}"; do
  printf "\033[35m \n Test for go run . %s | cat -e : \n" "\"$input"\" 
  output="$($GO_PROGRAM "$input" | cat -e)" 
  printf "\033[32m%s\n" "$output"
  printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"
done
