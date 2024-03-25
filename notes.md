# Notes from a beginner golang developer

## Testing sql queries in golang

Testing sql queries is a thing. When I test my API handlers, at some point I need to be able to test against a data source. In that case, I see two main options:

- mocking databases -> most popular option: [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock?source=post_page-----5af19075e68e--------------------------------)
- testing against a real database

Mocking database is most of the time an overkill and pretty cumbersome, even for big organisations. Testing against a real database is a good option since spinning up a database nowadays is just spinning up a container. As containers are easy to destruct, it is a solution for db testing.

There are some tools out there to spin up docker containers easily in go code. [dockertest](https://github.com/ory/dockertest) is one of those.

Read [this post](https://www.reddit.com/r/golang/comments/u62emg/mocking_database_or_use_a_test_database/) for the whole argument around this topic.

## TestMain and the use case of it

TODO: read [this article](https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc) and summarise it here.
