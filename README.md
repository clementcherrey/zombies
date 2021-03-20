# Winter is coming Game

## Overview

This is a server to play the "Talk to Zombies" game using telnet. The server support multiple games running at the same
time, and can be deployed on different OS using docker. However, player can not play together at the moment.

## Start the server

Build the docker image using the dockerfile

```dockerfile
docker build -t zombies:local .
```

Then use this new docker image to launch your server

```dockerfile
docker run -p 8080:8080 zombies:local
```

## Play

use telnet to play the game

```shell
telnet localhost 8080
```

## Remark

This is a basic implementation. The server still require more tests, better error handling, etc...
