# Lion solution
### Pt 1: finding the key
- lioncap is downloaded from the given link
- by running `file` on it, we can see it's a pcap
- open the pcap in wireshark or NetworkMiner (or preferred)
- There's lots of http data, namely a lot of get requests for files
- By looking at the files (in wireshark: file -> Export Objects -> HTTP), 12 rar files can be found. they are named f1.rar through f12.rar
- Save all the rar files
- To combine the files, use Winrar, select f1-f12 -> extract to -> choose destination
- the extracted file is a windows executable (compiled golang binary) 
- the binary must be reversed (throw it in IDA or GHIDRA or something)
- inside ghidra you can look at the dissassembler, see that there's an encrypt function which XORs a string with the string "REEvurs", which is the key
- there are 2 options from here

### Part 2: finding the encrypted flag - opt 1
 - when looking at the dissassembler, an ip address can be seen `192.168.206.161`
 - there's a socket being opened to the IP over port 33333
 - a user can spin up a VM and assign it that IP, then listen on that port
 - When the binary is run, the `.161` VM will recieve the encrypted flag over that port
 - data is `000c1125303108611d033f392d271a761a324126472f`

### Part 2: finding the encrypted flag - opt 2 <- easier way
 - In the pcap, there's lots of data going back and forth between the host and `192.168.206.161`
 - Filter for just those packets, filter out all the http requests (and their associated tcp streams)
 - There should be one TCP stream left, towards the bottom, that is headed to `.161:33333`
 - The `data` field of packet `30462` contains `000c1125303108611d033f392d271a761a324126472f`

### Part 3: Decrypting
 - From part one we could see that some data was being XOR'd with the key
 - We now have the key, and some potential encrypted data
 - In order to undo XOR, simply XOR again :^)
 - If users XOR the data with the key ("REEvurs"), they'll get the flag
 - example solve script can be found in solver.py