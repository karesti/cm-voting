# Welcome to Codemotion voting webapp !

This aplication is built with [Revel Web framework](https://revel.github.io), the [Golang](https://golang.org) webframework.

Install docker and run the app with the provided image.

* Clone the project

* Build the docker image from the cm-voting directory

```
docker build -t cm-voting .
```

* Run the docker image 

```
docker run -it -p 9000:9000 cm-voting
```

* Check if it works !

```
http://localhost:9000/
```

If you are running on mac like me, you might be using Docker-Machine. The ip on URL should 

```
docker-machine ls
perso    *        virtualbox   Running   tcp://192.168.99.100:2376
```

In my case http://192.168.99.100:9000/