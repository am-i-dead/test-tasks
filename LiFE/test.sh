#!/bin/bash

array=(https://files.testfile.org/AUDIO/C/M4A/sample1.m4a
https://files.testfile.org/AUDIO/C/M4A/sample2.m4a
https://files.testfile.org/AUDIO/C/M4A/sample3.m4a
https://files.testfile.org/AUDIO/C/M4A/sample4.m4a
https://files.testfile.org/anime.mp3
https://files.testfile.org/AUDIO/C/AAC/sample1.aac
https://files.testfile.org/AUDIO/C/AAC/sample3.aac
https://files.testfile.org/AUDIO/C/CAF/sample1.caf
https://files.testfile.org/AUDIO/C/CAF/sample4.caf
https://files.testfile.org/AUDIO/C/CVSD/sample1.cvsd
https://files.testfile.org/AUDIO/C/CVSD/sample3.cvsd)

file_counter=0
for item in ${array[*]}
do
    let "file_counter += 1"
done

if [ $1 = "-h" ] || [ $1 = "--help" ]
then
  echo "The script downloades all files from links listed in array."
  echo "Example of usage: ./tesh.sh downloaded_files"
  exit
else
  path=$1
  if [ -d ${path} ]
  then
    rm -rf ${path}
  fi
  mkdir ${path}

  count=0
  for item in ${array[*]}
    do
      let "count += 1"
      name=${item##*.}
      if [ ! -d "${path}/${name^^}" ]
      then
        mkdir "${path}/${name^^}"
      fi
      wget -q -P "${path}/${name^^}" $item 
      echo "${count} of ${file_counter} files downloaded"
    done
fi