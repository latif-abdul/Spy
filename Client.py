#!/usr/bin/python           # This is client.py file

import socket               # Import socket module

s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)         # Create a socket object
host = "34.82.200.141" # Get local machine name
port = 12345                # Reserve a port for your service.

s.connect((host, port))
print (s.recv(1024))
s.close()                     # Close the socket when done