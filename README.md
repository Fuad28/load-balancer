A Load balancer implementation in GO from [John Cricket coding challenges](https://codingchallenges.fyi/challenges/challenge-redis)

# Load balancer implementation in GO

A simple implementation of a load balancer using the round-robin algorithm in GO.
Part of my attempt at John Crickett's ![Coding Chanllenges](https://codingchallenges.fyi/challenges/challenge-redis)

## Features

-   Round-robin algorithm.
-   Development and Production mode.
-   Dummy server created when load balancer is started in dev mode.
-   A config.json file to pass in server configuration.
-   Loads DB for disk on startup.
-   Asynchronous implementation using goroutines.
-   Easy to use and extend.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Fuad28/load-balancer.git
    cd load-balancer
    ```

2. Install dependencies:

    ```bash
    go get
    ```

## Usage

1. Edit the config.json file in the root folder:

    - Sample dev config
    
    ```go
    {
    	"env": "dev",
    	"port": 8000,
    	"numberOfServers": 3,
    	"randomServerOff": true
    }

    ```
    
    - Sample prod config

    ```go
    {
    	"env": "prod",
    	"port": 8000,
    	"servers": [
		    {
    			"address": "https://www.google.com/",
    			"healthCheckPath": "/"
		    },

    		{
    			"address": "https://www.netflix.com/ng/",
    			"healthCheckPath": "/"
    		}
	    ]
    }

    ```
    

2. Start the server:

    ```bash
    go run .
    ```

3. The server begins to run on localhost port 8000. You can then test by sending a request

    ```bash
     curl http://localhost:8000/   
    ```

    You observe that each time, a different server responds.



## Basic example
- Server
  
<img width="528" alt="Screenshot 2024-03-11 at 3 52 53 PM" src="https://github.com/Fuad28/load-balancer/assets/63596779/aed64022-934c-49d7-ac6a-c91315b9f86a">

- Requests
  
  <img width="608" alt="Screenshot 2024-03-11 at 3 53 57 PM" src="https://github.com/Fuad28/load-balancer/assets/63596779/ff2a1242-1d9c-4903-a748-8bd2d40a8892">

