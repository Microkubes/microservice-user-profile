User Profile Microservice
=========================

## Prerequisite
Create a project directory. Set GOPATH enviroment variable to that project. Add $GOPATH/bin to the $PATH
```
export GOPATH=/path/to/project-workspace
export PATH=$GOPATH/bin:$PATH
```
Install goa and goagen:
```
cd $GOPATH
go get -u github.com/goadesign/goa/...
```

## Compile and run the service:
Clone the repo:
```
cd $GOPATH/src
git clone https://github.com/JormungandrK/microservice-user-profile.git /path/to/project-workspace/src/github.com/JormungandrK/microservice-user-profile
```
Be sure to use the full domain name and resource path here (compatible with ```go get```).


Then compile and run:
```
cd /path/to/project-workspace/src/github.com/JormungandrK/microservice-user-profile
go build -o user-profile
./user-profile
```

## Change the design
If you change the design then you should regenerate the files. Run:
```
cd /path/to/project-workspace/src/github.com/JormungandrK/microservice-user-profile
go generate
```
**NOTE:** If the above command does not update the generated code per the changes in the design,
then run ```goagen bootstrap```:

```bash
goagen bootstrap -d github.com/JormungandrK/microservice-user-profile/design -o .
```


Also, recompile the service and start it again:
```
go build -o user-profile
./user-profile
```

## Other changes, not related to the design
For all other changes that are not related to the design just recompile the service and start it again:
```
cd $GOPATH/src/github.com/JormungandrK/microservice-user-profile
go build -o user-profile
./user-profile
```

## Tests
For testing we use controller_test.go files which call the generated test helpers which package that data into HTTP requests and calls the actual controller functions. The test helpers retrieve the written responses, deserialize them, validate the generated data structures (against the validations written in the design) and make them available to the tests. Run:
```
go test -v
```

## Set up MongoDB
Create users database with default username and password.
See: [Set up MongoDB](https://github.com/JormungandrK/jormungandr-infrastructure#mongodb--v346-)
```
export MS_DBNAME=users
./mongo/run.sh
```
Then install mgo package:
```
cd $GOPATH
go get gopkg.in/mgo.v2
```

# Docker Builds

First, create a directory for the shh keys:
```bash
mkdir keys
```

Find a key that you'll use to acceess JormungandrK organization on github. Then copy the
private key to the directory you created above. The build would use this key to
access ```JormungandrK/microservice-tools``` repository.

```bash
cp ~/.ssh/id_rsa keys/
```

**WARNING!** Make sure you don't commit or push this key to the repository!

To build the docker image of the microservice, run the following command:
```bash
docker build -t user-profile-microservice .
```

# Running the microservice

To run the user-microservice you'll need to set up some ENV variables:

 * **SERVICE_CONFIG_FILE** - Location of the configuration JSON file (default: config.json)
 * **API_GATEWAY_URL** - Kong API url (default: http://localhost:8001)
 * **MONGO_URL** - Host IP(example: 192.168.1.10:27017)
 * **MS_USERNAME** - Mongo username
 * **MS_PASSWORD** - Mongo password
 * **MS_DBNAME** - Mongo database name

Run the docker image:
```bash
docker run user-profile-microservice
```

## Check if the service is self-registering on Kong Gateway

First make sure you have started Kong. See [Jormungandr Infrastructure](https://github.com/JormungandrK/jormungandr-infrastructure)
on how to set up Kong locally.

If you have Kong admin endpoint running on http://localhost:8001 , you're good to go.
Build and run the service:
```bash
go build -o user-profile
./user
```

To access the user service, then instead of calling the service on :8080 port,
make the call to Kong:

```bash
curl -v --header "Host: user-profile.services.jormugandr.org" http://localhost:8000/user-profile/1
```

You should see a log on the terminal running the service that it received and handled the request.

## Running with the docker image

Assuming that you have Kong and it is availabel od your host (ports: 8001 - admin, and 8000 - proxy) and
you have build the service docker image (user-microservice), then you need to pass
the Kong URL as an ENV variable to the docker run. This is needed because by default
the service will try http://localhost:8001 inside the container and won't be able to connect to kong.

Find your host IP using ```ifconfig``` or ```ip addr```.
Assuming your host IP is 192.168.1.10, then run:

```bash
docker run -ti -e API_GATEWAY_URL=http://192.168.1.10:8001 -e MONGO_URL=192.168.1.10:27017 user-profile-microservice
```

If there are no errors, on a different terminal try calling Kong on port :8000

```bash
curl -v --header "Host: user-profile.services.jormugandr.org" http://localhost:8000/user/1
```

You should see output (log) in the container running the service.



# Service configuration

The service loads the gateway configuration from a JSON file (See **SERVICE_CONFIG_FILE**
ENV variable).
Here's an example of a JSON configuration file:

```json
{
  "name": "user-profile-microservice",
  "port": 8080,
  "virtual_host": "user-profile.services.jormugandr.org",
  "hosts": ["localhost", "user-profile.services.jormugandr.org"],
  "weight": 10,
  "slots": 100
}
```

Configuration properties:
 * **name** - ```"user-profile-microservice"``` - the name of the service, do not change this.
 * **port** - ```8080``` - port on which the microservice is running.
 * **virtual_host** - ```"user-profile.services.jormugandr.org"``` domain name of the service group/cluster. Don't change if not sure.
 * **hosts** - list of valid hosts. Used for proxying and load balancing of the incoming request. You need to have at least the **virtual_host** in the list.
 * **weight** - instance weight - user for load balancing.
 * **slots** - maximal number of service instances under ```"user-profile.services.jormugandr.org"```.
