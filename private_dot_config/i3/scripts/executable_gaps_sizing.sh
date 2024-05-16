#!/bin/bash

Bigger ()
{
	i3-msg gaps outer all plus $1 && i3-msg gaps inner all plus $1
}

Smaller () 
{
	i3-msg gaps outer all minus $1 && i3-msg gaps inner all minus $1
}

Zero () 
{
	i3-msg gaps outer all set 0 && i3-msg gaps inner all set 0
}

# Check the input flag
if [[ "$1" == "--bigger" ]]; then
    Bigger $2
elif [[ "$1" == "--smaller" ]]; then
    Smaller $2
elif [[ "$1" == "--zero" ]]; then
    Zero
else
    echo "Invalid flag. Please use --bigger or --smaller."
fi
