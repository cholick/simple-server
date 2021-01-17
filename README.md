# simple-deployment

This is a server that serves the contents of the `CONTENT` environment variable.

Motivation: a small container to quickly toss into various Kubenretes resources that both
* Has no dependencies
* Can be deployed multiple times & trivially made distinguishable across those different primitives
