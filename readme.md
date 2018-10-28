#Paragliding

Created by André Gyrud Gunhildberget

Student Number : 493561

Heroku Link : https://paragliding-andregg.herokuapp.com/

Discord link to check clock trigger (https://discord.gg/tCNeAKf)

Discord webhook: [Discord webhook link](https://discordapp.com/api/webhooks/506067534336753664/vrUY-fQ6A-dkRiIgR8SxV0-78HsC1ttVa2fQi0iD2ezyLuZo1lHoiD6tSsXa-_U2NpL9)

#Clock Trigger Explanation
The clock trigger code is located in the "clockTrigger" folder for the convenience of having all the code in one repository. The code is however standalone and running on openstack without affiliation to the api codebase. (You can copy both files and run it independently.)

#Choices and motivation
I decided to use Unix time for the timestamps. One motivation behind the choise is that unix time is easier for the user to remember and interpret. A downside is that it is quite easy for multiple users to get the same timestamp. It is also quite easy to scrape an api if it is indexed by unixtime. A improvement would be to hash the timestamp so it would become unique or do as following: unixtimestamp + 6 UNIQUE random numbers. (Based on some random variable)

#Additional Information
I decided to use mgo as the MongoDB GO Driver because i found there to be more and better documentation and code examples. 

When testing don't use VSCodes default package tester. This times out due to tests taking more than 30 sec. The database requests are not fast enough. I suggest using "go test -cover" from the terminal or changing the database to some local database.
