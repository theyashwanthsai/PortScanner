import socket

target_host = input("Enter the host to scan: ")
target_ip = socket.gethostbyname(target_host)

# scan ports between 1 to 1024
for port in range(1, 1025):
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.settimeout(1)
    result = sock.connect_ex((target_ip, port))
    if result == 0:
        print("Port {} is open".format(port))
    sock.close()
