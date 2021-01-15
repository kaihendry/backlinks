redo-ifchange $2.mdwn $2.bl
cmark < $2.mdwn

echo "<h1>Backlinks</h1>"

ls $2.bldir | while read bl
do
	echo $bl backlinks to $2 >&2
	page=${bl%.*}
	echo "[${page}](${page}.html)" | cmark
done
