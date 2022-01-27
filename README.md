# Example dockerization of golang microservice with mysql database

### docker commands

Build individual image

```sh
docker build -t imagename:multistage -f sample.Dockerfile .
```

Run individual image

```sh
docker run -dp 8181:8181 --name containername imagename
```

Tag and Push individual image

```sh
docker tag imagename username/imagename
docker push username/imagename
```

### docker-compose commands

up

```sh
docker-compose up
```

up & build

```sh
docker-compose up -d --build
```

down

```sh
docker-compose down
```

push

```sh
// docker login first.
docker login -u username -p password
docker-compose push
```
