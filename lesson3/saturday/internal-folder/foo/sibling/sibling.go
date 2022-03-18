package sibling

import "PicusBootcamp/lesson3/saturday/internal-folder/foo/internal"

//because sibling and internal folders are the same roots,
//inside the sibling folder, you can reach the internal folders code

func UseInternalSum() int {
	return internal.Sum(2, 3)
}
