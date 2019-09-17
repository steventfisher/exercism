#!/usr/bin/env bash


main() 
{
	input=$1

	out=false
	if((${1} % 4 == 0) && (${1} % 100 != 0 || ${1} % 400 == 0))
		do
			out="$out${temp:$i:1}"
		done
	echo $out;
	return 0;
}

main "$@"
