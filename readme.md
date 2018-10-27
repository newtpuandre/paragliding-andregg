#IGC Track Viewer

Created by André Gyrud Gunhildberget

Student Number : 493561

Heroku Link : http://igcviewer-andregg.herokuapp.com/

Discord link to check clock trigger (https://discord.gg/tCNeAKf)

#Clock Trigger Explanation
The clock trigger code is located in the "clockTrigger" folder for the convenience of having all the code in one repository. The code is however standalone and running on openstack without affiliation to the api codebase. (You can copy both files and run it independently.)

#Choices and motivation
I decided to use Unix time for the timestamps. One motivation behind the choise is that unix time is easier for the user to remember and interpret it. A downside is that it is quite easy for multiple users to get the same timestamp. A improvement would be to hash the timestamp so it would become unique.

#Additional Information
When testing don't use VSCodes default package tester. This times out due to tests taking more than 30 sec. The database requests are not fast enough. I suggest using "go test -cover" from the terminal or changing the database to localhost.

There is still improvements to be done. There are comments throughout the code where i state what i would like to improve about that specific function / part

