BEGIN {
	OFS="\t"
}
substr($0, 1, 1) == "S" {
	match($0, /^Snapshot ([-a-z]+) revision ([0-9]+) created at ([-0-9]{10} [:0-9]{5})/, arr)
	snap=arr[1]
	rev=arr[2]
	revdate=arr[3]
	next
}
/^[ 0-9]/ {
	match($0, /^ *([0-9]+) ([-0-9]{10} [:0-9]{8}) (.{64}) (.*)$/, arr)
	hash=arr[3]
	if (latest[hash]) {
		overwrites++
	} else {
		items++
	}
	if (!latest[hash] || revdate > latest[hash]) {
		size=arr[1]
		date=arr[2]
		path=arr[4]
		# This one's the latest
		latest[hash]=revdate
		snaps[hash]=snap
		revs[hash]=rev
		dates[hash]=date
		sizes[hash]=size
		paths[hash]=path
		pathtime[hash]=path "__" date "__" hash
	}
}
END {
	print "snap", "rev", "revdate", "date", "size", "path", "hash"
	# Make an array with hashes sorted by pathtime
	for (hash in latest) {
		hashes[pathtime[hash]] = hash
	}
	n = asorti(hashes, hashpaths)
	for (i = 1; i <= n; i++) {
		hash = hashes[hashpaths[i]]
		print snaps[hash],revs[hash],latest[hash],dates[hash],sizes[hash],paths[hash],hash
	}
}

