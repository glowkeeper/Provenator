#!/bin/bash

# This doesn't do all the job, quite, but gets nearly there ;)
CWD=$(pwd)

cd ./contracts || exit 1

#echo -e "Enums.sol\n"
#cat ./contracts/Enums.sol | tr -d '\n' | grep -Po 'enum[^}]*}' | sed 's/    /\n    /g' | sed 's/}/\n}\n/'

# Turn structs into tuples
echo -e "Types.sol\n"
cat ./Types.sol | tr -d '\n' | grep -Po 'struct[^}]*}' | sed 's/struct[^{]*{ /tuple(/' | sed 's/;/,/g' | sed 's/,}$/)\n/' | tr -s " " | sed 's/e( /e(/'

#echo -e "---\n\nStorage.sol\n"
#cat ./Storage.sol | tr -d '\n' | grep -Po 'struct[^}]*}' | sed 's/    /\n    /g' | sed 's/}/\n}/'

echo -e "\n---\n"

for FILE in ./*.sol
do

  [[ -e "$FILE" ]] || exit 1
  echo -e "$FILE \n"

  # Turn events and functions into ethers.js human readable ABI form
  NAME=`echo $FILE | sed 's/\.sol//' | sed 's/^./\L&/'`
  echo "static ${NAME}ABI = ["
  egrep "event|function" $FILE | sed 's/^[ ]*\([e|f]\)/\1/' | sed s'/^\(.*\);$/\t"\1",/'

  echo -e "]\n\n---\n"

done

cd "$CWD" || exit 1
exit 0
