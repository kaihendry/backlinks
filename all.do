find . -name "*.mdwn" | while read src
do
	echo ${src%.*}.html
done | xargs redo-ifchange
