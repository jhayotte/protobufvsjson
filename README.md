#potobuf vs json
This project benchmark both technologies.

#explanation

There is two servers: One proto-server and one json-server reading messages received. Then there is one client sending messages in a predefine format either json either protobuf.

##scenario
A lapstime of 60 seconds is given to the client to send message to the server. Afterward we get the number of message sent and we clearly see a difference between both protocol:
![alt tag](https://github.com/jhayotte/protobufvsjson/blob/master/perfcompare-nbmessage.JPG)

This is clearly due to the fact that protobuf gives a better binary serialisation and the unmarshall step is faster due to the schema and contract between the client and server that avoid us to do reflection to determine the type of message received. 

#Profiling CPU

Run our profiling tool when the client is talking to the server with this command: `go tool pprof http://localhost:6060/debug/pprof/profile`

Install Graphviz to visualize pprof result http://www.graphviz.org/Download_windows.php

