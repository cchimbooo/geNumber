# geNumber
Numeric Type. 

I created this to consume some Bad desinged APIS.

This week (2021-07-20) I had to consume 2 APIS witch messed up types.

One of those sended a type as str / int

And the other one as a int/bolean (Staging) and string  (production). 

I decided to create a generic type to reade those. 

The type implements the DB interfaces (scanner / Value).  Those methods will result on an error if you try to insert a float in a table that only acept integers.


# WARNING!!! #

Do not use this type to other things, this should be use as an "adapter" after recieving the data you should use a cast method 
(Int(), Base64(), String(), Bool()) to cast to the right type and use that one in your project.  



