package errs

import "os"

func Quit(err error) {
	Handle(err, func() {
		println(err.Error())
		os.Exit(99)
	})
}

func Handle(err error, h func()) {
	if err != nil {
		h()
	}
}
