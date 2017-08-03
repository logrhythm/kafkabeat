# Kafkabeat

Kafkabeat is an elastic [beat](https://www.elastic.co/products/beats) that reads
events from a [Kafka](https://kafka.apache.org) topic and forwards them to
Elasticsearch.

## Getting Started with Kafkabeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Building

```
# Make sure $GOPATH is set
go get github.com/justsocialapps/kafkabeat
cd $GOPATH/src/github.com/justsocialapps/kafkabeat
make
```

### Running

To run Kafkabeat with debugging output enabled, run:

```
./kafkabeat -c kafkabeat.yml -e -d "*"
```


### Test

To test Kafkabeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`
