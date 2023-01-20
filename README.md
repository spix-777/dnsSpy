dnsSpy - the inernet spy
------------------------

dnsSpy is a Go program that takes in command line arguments for a base URL, a wordlist file, and an output file. It scans the wordlist file and performs HTTP GET requests on each word appended to the base URL. If the GET request is successful, it prints the valid URL and the corresponding IP address. It also writes the valid URL and IP address to the specified output file.

Usage:

The program can be run by providing the -u flag for the base URL, the -w flag for the wordlist file and the -o flag for the output file.

'dnsSpy -u example.com -w wordlist.txt -o output.txt

If no flags are provided, the program will use default values of url for the base URL, wordlist for the wordlist file, and dnsSpoy.txt for the output file.

Dependencies:

The program uses the following Go packages:

- flag for parsing command line arguments

- net and net/http for performing GET requests and looking up IP addresses

- bufio for reading the wordlist file

- os for opening and writing to the output file

- strings for formatting the IP address strings

- log for logging errors


Note:

It is important to note that this program is for educational and research purposes only. The use of this program for any unauthorized or illegal activity is strictly prohibited.