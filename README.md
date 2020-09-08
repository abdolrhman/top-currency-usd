# Top Coins Prices In USD

Getting the top Currency prices in USD with support to limitation and more features

## Table of Contents
- [Getting Started](#getting-started)
- [Install](#install)
- [Objectives](#objectives)
- [Implementation Solution](#implementation-solution)
- [How Edge-cases being Handled](#edgeCases)
- [Intended Approach](#intended-approach)
- [Images Work flow](#images)
## Getting Started

This is an app to get Top assets prices in usd, build in top of 3 independent services
- limitation
- up to date information
- speedy 

### Installing

- To be able to make the project work you need at least, postgres and go compiler
- Then you need to install the dependencies of this application, unfortuante i didnt have the time to find a better way to install the dependcies,
I have worked with many programming language, Go Seems to need a bit more work to install dep,
- So i have worked with minimum depends on outside packages
- Any way the project consist of 3 services, so you need to get into each one and use this command
```
go get .
```  

- Also support docker-compose
- I didn't use .env for simplicity

## Objectives
1. Provide HTTP EndPoint, to fetch top assets and its currency
2. provide limitation support
3. merging the data of two-API's for getting the required information
4. up-to date information 
5. independent three services each for specific tasks
6. descriptive readme file
7. tests 
8. docker
8. output should be json or CSV

## Implementation Solution

- Simply my plan was to break the task into seprate thee service
1. to grape the top assets, run in cron job each minute
2. to grepe asset currency in USD, run in cron job each minute
3. to push data in a shared database, periodically
4. to provide API for the merged data in persistence matter

- there are many solution for such structure, I have gone simple approach
as to save the data, into Shared database, which is "POSTGRES", <br>
```with a cron job that runs each 1 minute to the save two services data```<br>
so service 1, will push into table : coin_prices <br>
```- id int, symbol text, price real``` <br>
service 2, will push into table : top_assets
```- id int, symbol text, rank int``` <br>
> the rank and the API provider Page were tricky, you can see the code for more information, but data with limited to max 100 per 1 page, and for rank we need to make it be built in a matter to make it easy to update documents with the new rank in future
And then, the API service, consume the data in away in which


## How Edge-cases being Handled

#### How to save retrieve data
- using the go-pg driver, and build it in matter of structuts instead of querying the db directly

#### Up-to date 
- up to date, is being solved by running the two servcies unlimited for each minute 
and then <br>
> Insert or update the data

#### Rank And Page Issue
- so the  cryptocompare API provider gives for limited data, 100 max recored per page <br>
> this been solved by running two loops when quering and insertingor update the records for top_assets
- so issue here that once i have made two loops to loop over all of the data one to grape 100 each so tatal 200,
i got into a tricky problem because i was saving the rank as with an iterator, so this made an issue which been solved   

#### The most one that took time, Retriveing the data for the API
- so how whould i constract the required data, and make the merge,
> this has been done by 
>1. Get all currency prices in array 
>2. Get the limited top_assets Data and order by rank *important > tricky
>3. for each element inside top_asset, map it to the array of currency prices to get the price, once the record find get out of loop *> Tricky
## Intended Approach

I believe A better approach would be to 
- use queues as something like "Producer and consumer", 
- where the producers are two services each push on seprate Queue,
- And the consumer is the transformer or merging service
- this can be done via kafka or rabitmq
- that would end up with stream processor that would listen to both 
and transform them into a 3rd queue. 
- Now your end client only has to listen to the final stream to get updates.
--------
Also the current implementation its not that bad, <br>
but maybe it can be better, if there is a cache storage maybe 'redis', that cache the transformed or merged data, <br>
so instead for each http call i get the data and merge it <br>


## Images Work flow
> check the postman result with each asset contain rank and price

![Alt Text](https://media.giphy.com/media/S5iQjzLXDnynkKyHcU/giphy.gif)

> check 200 records exist for the two tables
> check for each cycle of cron job, updating records if needed, with order and rank in respect
 
![Alt Text](https://media.giphy.com/media/Tk1By3jMt0LtJXdIpL/giphy.gif)
