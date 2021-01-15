find . -name "*.mdwn" | while read src
do
	echo ${src%.*}.bl
	echo ${src%.*}.html
done | xargs redo-ifchange
