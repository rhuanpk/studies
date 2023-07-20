#!/usr/bin/env bash

arr=(150 15 1 30 2)

echo "~~~ Desordenado ~~~"
for((i=0;i<=1;++i)); do
	for((j=0;j<${#arr[@]}-1;++j)); do
		if [ ${i} -eq 0 ]; then
			echo "Indice[${j}]: ${arr[${j}]}"
			if [ ${j} -eq 3 ]; then
				echo "Indice[4]: ${arr[4]}"
			fi
		else
			if [ ${arr[j]} -ge ${arr[j+1]} ]; then
				aux=${arr[j]}
				arr[${j}]=${arr[${j+1}]}
				arr[${j+1}]=${aux}
			fi
		fi
	done
done

echo "~~~ Ordenado ~~~"
for((i=0;i<${#arr[@]};++i)); do
	echo "Indice[${i}]: ${arr[${i}]}"
done
