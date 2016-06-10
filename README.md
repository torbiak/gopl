# GOPL solutions

Solutions to every exercise in The Go Programming Language.

Many of the exercises are based on an example from the book or an earlier exercise, and I've copied files as needed so there's a separate package for each exercise. If I'd been thinking I would have committed as I went so that it's easier to use git to see what I started from. As it is you'll need to determine the predecessor example or exercise and use `diff` for exercises after 5.6.

## Interesting exercises

* [complex root visualization using Newton's method](ex3.7/main.go)
* [CSS-style selector for XML](ex7.17/main.go)
* [RFC959-compliant FTP server](ex8.2/ftpd.go)
* [du clone](ex8.9/main.go)
* [web mirroring tool](ex8.10/mirror.go)
* [chat server](ex8.15/chat.go)
* [cancellable function memoization](ex9.3/memo.go)
* [archive reader](ex10.2/arprint.go)
* [go package ancestor tool](ex10.4/ancestors.go)
* [generate random non-palindromes using a grammar](ex11.3/palindrome_test.go)

## Package list

* [ex1.1](ex1.1): ex1.1 prints command line arguments.
* [ex1.2](ex1.2): ex1.2 prints commandline indexes and arguments.
* [ex1.3](ex1.3): ex1.3 compares string concatenation techniques.
* [ex1.4](ex1.4): ex1.4 prints the count and text of lines that appear more than once in the input.
* [ex1.5](ex1.5): ex1.5 generates GIF animations of random Lissajous figures, with a green-on-black palette.
* [ex1.6](ex1.6): ex1.6 generates GIF animations of random Lissajous figures, with a gradient applied on the time dimension.
* [ex1.7](ex1.7): ex1.7 streams the content found at each specified URL to stdout.
* [ex1.8](ex1.8): ex1.8 streams the content found at each specified URL to stdout, and appends http:// to arguments as needed.
* [ex1.9](ex1.9): ex1.9 streams the content found at each specified URL to stdout and prints the HTTP status code.
* [ex1.10](ex1.10): ex1.10 fetches URLs in parallel and reports their times and sizes.
* [ex1.12](ex1.12): ex1.6 generates GIF animations of random Lissajous figures, with a gradient applied on the time dimension and the number of cycles to display can be supplied in the url query string when running as a web server.
* [ex2.1](ex2.1): ex2.1 performs Celsius, Fahrenheit, and Kelvin conversions.
* [ex2.2](ex2.2): ex2.2 prints measumrents given on the command line or stdin in various units.
* [ex2.3](ex2.3): ex2.3: compare popcount implementations: looping and single-expression bytewise table lookups.
* [ex2.4](ex2.4): ex2.4: compare popcount implementations, including looping table lookups and shift value.
* [ex2.5](ex2.5): ex2.4: compare popcount implementations, including clear rightmost.
* [ex3.1](ex3.1): ex3.1 prints an svg image, ignoring non-finite vertexes.
* [ex3.2](ex3.2): ex3.2 prints an SVG rendering of an eggbox or saddle.
* [ex3.3](ex3.3): ex3.3 prints an svg image, coloring its vertices based on their height.
* [ex3.4](ex3.4): ex3.4 serves SVG rendering of a 3-D surface function over http.
* [ex3.5](ex3.5): ex3.5 emits a full-color PNG image of the Mandelbrot fractal.
* [ex3.6](ex3.6): ex3.5 emits a supersampled image of the Mandelbrot fractal.
* [ex3.7](ex3.7): ex3.7 visualizes how many iterations it takes to find complex roots of a quartic equation using Newton's method, using different colors for different roots.
* [ex3.8](ex3.8): ex3.8 compares different numeric types when rendering fractals.
* [ex3.9](ex3.9): ex3.9 serves images of fractals over http.
* [ex3.10](ex3.10): ex3.10 inserts commas into integer strings given as command-line arguments, without using recursion.
* [ex3.11](ex3.11): ex3.11 inserts commas into floating point strings with an optional sign, given as command-line arguments.
* [ex3.12](ex3.12): ex3.12 determines if strings are anagrams of each other.
* [ex3.13](ex3.13): ex3.13 is a short definition of byte unit constants.
* [ex4.1](ex4.1): ex4.1 computes the number of bits different between two hashes.
* [ex4.2](ex4.2): ex4.2 prints the SHA hash of stdin.
* [ex4.3](ex4.3): ex4.3 reverses an array
* [ex4.4](ex4.4): ex4.4 rotates a slice of ints by one position to the left.
* [ex4.5](ex4.5): ex4.5 dedupes a slice of strings.
* [ex4.6](ex4.6): ex4.6 reverses a utf8 string.
* [ex4.7](ex4.7): ex4.7 reverses a utf8 string in-place.
* [ex4.8](ex4.8): ex4.8 computes counts of Unicode characters.
* [ex4.9](ex4.9): ex4.9 counts word frequency for stdin.
* [ex4.10](ex4.10): ex4.10 prints a table of GitHub issues matching the search terms, organized by the past day, month, and year.
* [ex4.11](ex4.11): Package github provides a Go API for the GitHub issue tracker.
* [ex4.12](ex4.12): ex4.12 gets, indexes, and searches xkcd comic metadata.
* [ex4.13](ex4.13): ex4.13 searches OMDB by title and downloads a movie poster.
* [ex4.14](ex4.14): Package github provides a Go API for the GitHub issue tracker.
* [ex5.1](ex5.1): ex5.1 prints the links in an HTML document read from standard input.
* [ex5.2](ex5.2): ex5.2 counts the frequency of different tags in an html document on stdin.
* [ex5.3](ex5.3): ex5.3 prints nonempty text tokens from an html document on stdin.
* [ex5.4](ex5.4): ex5.4 prints the links in an HTML document read from standard input, including those for images, scripts, and style sheets.
* [ex5.5](ex5.5): ex5.5 counts the number of words and images at a url.
* [ex5.6](ex5.6): ex5.6 computes an SVG rendering of a 3-D surface function, using a bare return statement.
* [ex5.7](ex5.7): ex5.7 pretty-prints html.
* [ex5.8](ex5.8): ex5.8 finds an html element node by id attribute.
* [ex5.9](ex5.9): ex5.9 expands shell-style variable references on stdin.
* [ex5.10](ex5.10): ex5.10 sorts courses topologically based on hard-coded dependencies.
* [ex5.11](ex5.11): ex5.11 reports on cycles in course dependencies.
* [ex5.12](ex5.12): ex5.12 prints the outline of an HTML document tree.
* [ex5.13](ex5.13): ex5.13 saves a local mirror of a website.
* [ex5.14](ex5.14): ex5.14 prints a random course's prerequisites, recursively.
* [ex5.15](ex5.15): ex5.15 explores variadic min and max functions.
* [ex5.16](ex5.16): ex5.16 provides a variadic string-joining function.
* [ex5.17](ex5.17): ex5.17 uses a variadic ElementsByTagName function to find html nodes.
* [ex5.18](ex5.18): ex5.18 saves the contents of a URL to a local file.
* [ex5.19](ex5.19): ex5.19 returns a non-zero value using panic and recover, contradicting the function signature.
* [ex6.1](ex6.1): ex6.1 is a bit vector integer set.
* [ex6.2](ex6.2): ex6.2 is a integer set with a variadic AddAll method.
* [ex6.3](ex6.3): ex6.3 is a bit vector integer set with binary set operations.
* [ex6.4](ex6.4): ex6.4 is an integer set with an Elems method.
* [ex6.5](ex6.5): ex6.5 is a variable-word-size bit vector intset.
* [ex7.1](ex7.1): ex7.1 provides line and word counters.
* [ex7.2](ex7.2): ex7.2 wraps a writer to count written words.
* [ex7.3](ex7.3): ex7.3 provides insertion sort using an unbalanced binary tree, and a String method to visualize the tree.
* [ex7.4](ex7.4): ex7.4 provides a simple string reader.
* [ex7.5](ex7.5): ex7.5 provides a LimitReader that reports EOF at a given offset.
* [ex7.6](ex7.6): ex7.6 prints flag arguments for different temperature scales, including Kelvin.
* [ex7.8](ex7.8): ex7.8 provides iterative columnar sorting for Persons.
* [ex7.9](ex7.9): ex7.9 serves an html table with a stable column sort.
* [ex7.10](ex7.10): ex7.10 uses sort.Interface to detect palindromes.
* [ex7.11](ex7.11): ex7.11 adds CRUD http endpoints to a simple PriceDB server.
* [ex7.12](ex7.12): ex7.12 convert's PriceDB list output to an html table.
* [ex7.13](ex7.13): ex7.13 adds pretty-printing to an arithmetic expression parser.
* [ex7.14](ex7.14): ex7.14 adds factorials to an arithmetic expression parser.
* [ex7.15](ex7.15): ex7.15 evaluates an expression using given variable bindings.
* [ex7.16](ex7.16): ex7.16 is a web-based calculator.
* [ex7.17](ex7.17): ex7.17 provides CSS-style selectors for XML.
* [ex7.18](ex7.18): ex7.18 parses XML into a tree of nodes, using the token-based API of encoding/xml.
* [ex8.1/clock](ex8.1/clock): clock is a TCP server that periodically writes the time.
* [ex8.1/clockwall](ex8.1/clockwall): clockwall listens to multiple clock servers concurrently.
* [ex8.2](ex8.2): ex8.2 is a minimal ftp server as per section 5.1 of RFC 959.
* [ex8.3](ex8.3): ex8.3 is a simple read/write client for TCP servers.
* [ex8.4](ex8.4): ex8.4 is a reverb server that uses sync.WaitGroup to choose when to close connections.
* [ex8.5](ex8.5): ex8.5 is a parallellized Mandelbrot fractal generator.
* [ex8.6](ex8.6): ex8.6 is a depth-limited web crawler.
* [ex8.7](ex8.7): ex8.7 mirrors a website to a given depth using multiple goroutines and rewrites local links.
* [ex8.8](ex8.8): ex8.8 is a reverb server that disconnects inactive clients.
* [ex8.9](ex8.9): ex8.9 is a concurrent du clone.
* [ex8.10](ex8.10): ex8.10 is a web-mirroring tool that can be gracefully interrupted using ctrl-c.
* [ex8.11](ex8.11): ex8.11 prints the first HTTP response received.
* [ex8.12](ex8.12): ex8.12 is a server that lets clients chat with each other.
* [ex8.13](ex8.13): ex8.13 is a chat server that disconnects inactive clients.
* [ex8.14](ex8.14): ex8.14 is a chat server that prompts clients for a name upon connection.
* [ex8.15](ex8.15): ex8.15 is a chat server that skips clients that are slow to process writes.
* [ex9.1](ex9.1): ex9.1 provides a concurrency-safe bank, with withdrawals.
* [ex9.2](ex9.2): ex9.2 lazily initializes the popcount LUT.
* [ex9.3](ex9.3): ex9.3 provides cancellable memoization of a function.
* [ex9.4](ex9.4): ex9.4 tests the performance of goroutine pipelines.
* [ex9.5](ex9.5): ex9.5 tests of performance of ping-ponging goroutines.
* [ex10.1](ex10.1): ex10.1 converts images to between jpg, png, and gif formats.
* [ex10.2](ex10.2): ex10.2 detects and reads zip and tar archives.
* [ex10.2/cmd/arprint](ex10.2/cmd/arprint): Print the names and contents of the files in a tar or zip archive.
* [ex10.4](ex10.4): ex10.4 lists go packages that transitively depend on the given packages.
* [ex11.1](ex11.1): ex11.1 computes counts of Unicode characters, and includes tests.
* [ex11.2](ex11.2): Package intset provides a set of integers based on a bit vector.
* [ex11.3](ex11.3): ex11.3 tests IsPalindrome on random non-palindromes.
* [ex11.4](ex11.4): ex11.4 tests IsPalindrome on strings including punctuation.
* [ex11.5](ex11.5): ex11.5 tests strings.Split using a table-driven test.
* [ex11.6](ex11.6): ex11.6 benchmarks popcount implementations.
* [ex11.7](ex11.7): Package intset provides a set of integers based on a bit vector.
* [ex12.1](ex12.1): ex12.1 uses reflection to print arbitrary values.
* [ex12.2](ex12.2): ex12.2 displays arbitrary values to a certain depth.
* [ex12.3](ex12.3): ex12.3 is a codec for s-expressions.
* [ex12.4](ex12.4): ex12.4 is a codec for s-expressions, with pretty-printing.
* [ex12.5](ex12.5): ex12.5 is a codec for json.
* [ex12.6](ex12.6): ex12.6 is an s-expression codec that doesn't encode zero values.
* [ex12.7](ex12.7): ex12.7 provides a streaming decoder for s-expressions.
* [ex12.8](ex12.8): ex12.8 can unmarshall s-expressions from an io.Reader.
* [ex12.9](ex12.9): ex12.9 is token-based API for decosing s-expressions.
* [ex12.10](ex12.10): ex12.10 is a s-expression codec that can decode booleans, floating point numbers, and registered interface values.
* [ex12.11](ex12.11): ex12.11 provides a reflection-based codec for URL query parameters.
* [ex12.12](ex12.12): ex12.12 provides a URL query parameter codec with validation triggered by struct tags.
* [ex12.13](ex12.13): ex12.13 is a s-expression codec that uses names found in struct tags.
* [ex13.1](ex13.1): ex13.1 provides a deep equivalence relation for arbitrary values.
* [ex13.2](ex13.2): ex13.2 determines if a value is cyclic.
* [ex13.3](ex13.3): ex13.3 bzip provides a concurrency-safe writer that uses bzip2 compression.
* [ex13.4](ex13.4): ex13.4 provides a bzip2 writer using the system's bzip2 binary.

## License

[CC BY-NC-SA 4.0](http://creativecommons.org/licenses/by-nc-sa/4.0/), the same as the example code for The Go Programming Language.
