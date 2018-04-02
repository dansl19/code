#!/bin/bash

read -p "Please enter a word " word

if [ ${#word} -ge 5 ]; then echo "error" ; exit
else
for i in $(seq 0 ${#word})
    do array[$i]=${word:$i:1}
done

second_last=${array[$(expr ${#array[@]} - 3)]}
last=${array[$(expr ${#array[@]} - 2)]}
echo "$last $second_last"
fi