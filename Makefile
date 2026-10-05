.PHONY : all

all : snapshot-fuse

darwin : snapshot-fuse

snapshot-fuse : *.go dpfs/*.go
	env GOOS=darwin asdf exec go build .

~/.cache/duplicacy-files-sample.csv: ~/.cache/duplicacy-files-sample.tsv
	<$< vd --batch -f tsv - -o $@

~/.cache/duplicacy-files-sample.tsv: ~/.cache/duplicacy-files.txt duplicacy-files.awk
	<$< pv | LANG=C gawk -F'\n' -f duplicacy-files.awk >$@

~/.cache/duplicacy-files.txt:
	duplicacy list -files -a -storage canasta > $@

clean:
	rm ~/.cache/duplicacy-files.txt
