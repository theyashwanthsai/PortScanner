import sys
import socket
from datetime import datetime

#target
if(len(sys.argv) == 2):
	target = socket.gethostbyname(sys.argv[1])
	#Translating host name to IP v4
else:
	print("Invalid amount of arguments.")
	print("Syntax: python3 scanner.py <ip>")
 
#making a banner to make things look good :)
print("_"*50)
print("Scanning target "+target)
print("Time started: " + str(datetime.now()))
print("_"*50)
 
try:
 	for port in range(50, 85):
       
      
 	    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
 		socket.setdefaulttimeout(1)
 		result = s.connect_ex((target, port)) #returns an error
 		if result == 0:
 			print("Port {} is open".format(port))
 		s.close()
  
except KeyboardInterrupt:
 	print("\n Quitting Port Scanner...")
 	sys.exit()
 	 
except socket.gaierror:
	print("Hostname could not be resolved...")
	sys.exit()

except socket.error:
	print("Couldn't connect to server...")
	sys.exit()