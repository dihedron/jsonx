# JSONX

A very, very simple tool to remove comments from files in JSON-like format.

## Rationale

JSON format does not support comments; unfortunately JSON is often used for keeping information in human-readable and machine-friendly format, e.g. in configuration files or other forms of data interchange. It would be very useful if there were a way to include comments in JSON files and have them purged when the file is fed into a parser conforming to the standard.

This is why I wrote `jsonx`.

## Example JSONX file

JSON with comments, or JSONX, supports C++-style and shell-style comments:

```JSON
{
	// this is a C++-style comment
	"name" : "John",
	"surname": "Doe", // this is another C++-style comment
	"address": "Liberty Plaza, 1",
	# don't forget shell-style comments
	"phone": "555-123-456-7890" # at end of line
}
```
`jsonx` can strip comments from these files and return the standard-conforming embedded JSON document:
```JSON
{
	"name" : "John",
	"surname": "Doe", 
	"address": "Liberty Plaza, 1",
	"phone": "555-123-456-7890" 
}
```

## Usage

In order to use `jsonx`, you can provide an input file, and input and an output file, or use it as a filter in a pipe.

1. using `jsonx` with input and output file:
   ```bash
   $> jsonx input.jsonx output.json
   ```
2. using `jsonx` with only an input file; the stamdarard-conformat JSON will be written to standard output:
   ```bash
   $> jsonx input.jsonx
   {
	    "name" : "John",
	    "surname": "Doe", 
	    "address": "Liberty Plaza, 1",
	    "phone": "555-123-456-7890" 
	}
	```
3. using `jsonx` in a pipe:
   ```bash
   $> cat input.jsonx | jsonx
   {
	    "name" : "John",
	    "surname": "Doe", 
	    "address": "Liberty Plaza, 1",
	    "phone": "555-123-456-7890" 
	}
	```
