# Welcome to Codemotion voting webapp !

This aplication is built with [Revel Web framework](https://revel.github.io), the [Golang](https://golang.org) webframework.

Install docker and run the app with the provided image.

## Install and run Mongodb in a Docker Container

```
docker pull mongo

docker run -i -t --name mongo_cmvoting -p 27017:27017 mongo
```

To check if this is working go to the database console


```
mongo --host 192.168.99.100 (or localhost if you are not on using like me a docker-machine)`
```


## Install and run the web app in a Docker Container

* Clone this project

* Build the docker image from the cm-voting directory cloned from Github

```
cd cm-voting

docker build -t cm-voting .
```

* Run the docker image linking the already running MongoDB container

```
docker run -i -t -p 9000:9000 --link mongo_cmvoting:mongo cm-voting env
```

If you are lazy :) you can execute buildAndRun.sh which will execute build image and run 

```
./buildAndRun.sh
```

Or with docker-compose :

```
docker-compose up
```

* Check if it works !

```
http://localhost:9000/
```

If you are running on mac like me, you might be using Docker-Machine. Instead of using localhost, you must use the IP of your host which can be retrieved via : 

```
docker-machine ls
perso    *        virtualbox   Running   tcp://192.168.99.100:2376
```

In my case http://192.168.99.100:9000/
