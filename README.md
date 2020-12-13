# requestaggregator
One of the most common data structures used around is a tree. This project coded with building the same which can collect data and aggregate on some dimensions and metrics.

# Description of the Keys and Values

There is a website which has traffic globally from different devices. In this case, your dimensions are:

 1. Country: Represented by a string like IN,US,KR etc
 2. Device: Represented by a string like Mobile, Desktop, Tablet

We also have a couple of metrics as well on these dimensions. Metrics are the actionable values that a user generates. Some examples of the same are:

 1. WebRequests: Represented by an integer denoting the number of times a request has been made to a website
 2. TimeSpent: Represented by an integer which denotes the time spent in seconds by users on the website

# Running this Project

 1. Install Golang and Docker on your system
 2. Clone this repo using `git clone https://github.com/shivamverma4/requestaggregator`
 3. Keep your Docker running, and create a docker image using command `docker build -t requestaggregator .`
 4. Now run this docker image with the docker IMAGE_ID, using command `docker run -d -p 8081:8081 xxxxxxxxxxxxx`
 5. Now the project is UP and running on port `8081`
 6. Get Request Aggregator API, `/v1/query`
 7. Insert Request Aggregator API, `v1/insert`
