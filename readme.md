This is a command-line interface (CLI) program written in Go that implements an key-value store with persistence and additional features.

Features
Persistence to Disk: The program stores key-value pairs on disk, ensuring that the data is retained between program runs. Data is saved to and loaded from a file named data.txt.

Commands:

SET key value: Store a key-value pair in the store.
GET key: Retrieve the value associated with a given key.
DELETE key: Delete a key-value pair from the store. 
LIST: Display all stored keys in the store.
EXIT: Exit the program.
SAVE : Save the keys used till now to disk persistently
LOAD : Load the key values from the disk data.txt file

Usage
Installation: Ensure you have Go installed on your system.

Download: Clone or download this repository to your local machine.

Run the Program:

Open a terminal.
Navigate to the directory containing the downloaded files.
Run the program using the command: go run main.go.
Commands:

To store a key-value pair: SET key value
To retrieve a value: GET key
To delete a key-value pair: DELETE key
To list all stored keys: LIST
To exit the program: EXIT
Persistence:

The program automatically persists data to the data.txt file in the same directory.
Data is loaded from the file when the program starts.
Example
ruby
>> SET name John
Stored key 'name' with value 'John'
>> GET name
Value for key 'name': John
>> DELETE name
Key 'name' deleted
>> LIST
Stored keys:
>> SET age 30
Stored key 'age' with value '30'
>> LIST
Stored keys:
age
>> EXIT
Exiting the program.

Notes
The data.txt file is used for data persistence. Make sure to keep it safe and not modify its contents manually.
We can use it automatically persistent by using LoadData function on every SET and DELETE opeartion rather giving separate commands
