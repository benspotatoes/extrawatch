glide:
	glide update --strip-vendor
	glide-vc --only-code --no-tests

run:
	foreman start
