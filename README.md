# groupby
Golang implementation of unix command similar to SQL group by statement

## Description
Reads standard input and groups each line and counts the number of occurrences.

## Usage
Pipe text to it, eg: cat file.txt | go run groupby.go

## Alternative
This is similar to what can be done by first sorting and then using the uniq command, eg: cat file.txt | sort | uniq -c

## Performance
Why create this command with standard unix commands available to solve the same problem? Mainly since sorting and then using uniq
 is a lot less efficient. For large data sets sorting (n log n operation) is fairly expensive and can be avoided. Example run times
 on a 200M line file of IP addresses:

```
time zcat ip.txt.gz | sort -n | uniq -c | wc -l
744380

real	5m12,954s
user	5m20,367s
sys	0m8,921s



time zcat ip.txt.gz | go run groupby.go | wc -l
744380

real	0m31,146s
user	0m45,988s
sys	0m2,354s
```
In this slightly extreme example groupby.go is 10x faster.