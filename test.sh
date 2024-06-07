#!/bin/bash

INPUTS=(
    "hello"
    "Hello There!"
    "Hello There!"
)

BANNERS=(
    "standard"
    "shadow"
    "thinkertoy"
)

GO_PROGRAM="go run ."

printf "\n \033[36m%s\033[0m" "////////////////////////TESTS///////////////////////////"
printf "\n \033[35m◼\033[0m : Input \033[32m◼\033[0m : Output \n"

for input in "${INPUTS[@]}"; do
  for banner in "${BANNERS[@]}"; do
    printf "\033[35m \n Test for go run . %s %s | cat -e : \n" "\"$input\"" "$banner"
    output="$($GO_PROGRAM "$input" $banner | cat -e)"
    printf "\033[32m%s\n" "$output"
    printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"
  done
done
