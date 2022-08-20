import socket
host = '127.0.0.1'
port = 7777

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host, port))