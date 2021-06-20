
`go get -u github.com/cchimbooo/geNumber
`
# geNumber
Numeric Type. 

I created this to consume some Bad designed APIS.

This week (2021-07-20) I had to consume 2 APIS witch messed up types.

One of those sent a type as str / int.

The other one as an int/boolean (Staging) and string  (production). 

I decided to create a generic type to reade those. 

The type implements the DB interfaces (scanner / Value).  Those methods will result on an error if you try to insert a float in a table that only accept integers.


## How to use: 

Import the package, and use ge.Numeric as a type in your models. 
This allows it to read: strings (numbers), integers, floats and booleans (cast to 0-1)


# WARNING!!! #

Do not use this type to other things, this should be use as an "adapter" after receiving the data you should use a cast method 
(Int(), Base64(), String(), Bool()) to cast to the right type and use that one in your project.  



