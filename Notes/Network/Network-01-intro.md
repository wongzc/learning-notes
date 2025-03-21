# Network intro
### source: xiaolincoding
https://xiaolincoding.com/network

1. all network follow TCP/IP model, 4 layers
    1. application layer
        - OSI: application, presentation, session layer
        - HTTP, HTTPS
        - focus on provide function to user

    2. transport layer
        - TCP, UDP
            - TCP (Transmission Control Protocol): 
                1. flow control
                2. timeout retransmission
                3. congestion control
            - UDP:
                1. faster
        - network support for application layer
        - data sent by application layer if bigger than MSS ( largest length of TCP segment), need to seperate into different segment
            - if any segment lost/ corrupted, resend only the segment
            - need to specify port on device to send, as many application on device is receving as well.

    3. network layer/ internet layer
        - IPV4, IPV6
        - the actual data transfer between device, from 1 to 1
        - using IP ( internet protocol), translate message into IP data
            - if bigger than MTU (1500 byte) will split again
            - each split wtih TCP head, IP head, MAC head ( MAC not include in MTU)
        - to locate device when transfer data:
            - network id: identify subnet belong to which IP address
            - host id: identify different host under same subnet
                - use subnet mask to calculate network ID & host ID
                    - xxxx/24 means mask = 255.255.255.0 ( 11111111.11111111.11111111.00000000)
                    - use "AND" between mask & IP to get network ID
                        - 00001010.1100100.1111010.00000010 with 11111111.11111111.11111111.00000000
                        - get 0001010.1100100.1111010.00000000
            - when locating, will find the same network id, then look for host
        - each device connected through multiple gateways, routers, and switches
            - IP use routing algo to determine which path to go

    4. Link layer
        - OSI: data link, physical layer
        - add MAC head infront of IP head, encapsulates into data frame
        - enthernet:
            - interface & components to connect device in local area network (LAN)
            - ethernet port, wifi interface, ethernet cable
            - connect nearby device
            - identify destination by MAC header, contain:
                - destination MAC addess
                - source MAc address
            - use ARP (address resolution protocol) to discover MAc address of another device
        - link layer provide link level trsnaport service for network layer
        - bridge the gap between IP & physical network, ensure data packet transmitted over ethernet / wifi to correct device ( by MAC) 
    5. data in each layer 
        - link layer: frame
        - network: packet
        - transport: segment
        - application: message

    6. what happen after key in URL?
        1. HTTP
            - analyze URL to send info to web server
            - http://www.server.com/dir1/file1.html
                - http: protocol
                - www.server.com: web server
                - /dir1/file1.html: file name/path
            - if no file name/path, will visit default file
                - /index.html or /default.html
            - create HTTP request, GET/ POST etc
        2. DNS
            - request DNS server to get IP address
            - on the right is higher level
            - www.server.com. (at the end has a .)
                - root DNS server "."
                - top level DNS server ".com"
                - authoritative DNS server "server.com"
            - step:
                - request IP address from local DNS server (server address in client TCP/IP setting)
                - if in cache, return, else check with root DNS (.)
                - then root will redirect to top-level DNS (.com)
                - then redirect to authoratuve DNS, and get IP
                - browser then send http request to IP addres
            - will check browser cache, then os cache, then host file, before request from local DNS
        3. protocol stack
            - protocol stack in charge of HTTP transport after receive IP
            - different layer, move from up to down
                - application: browser, socket
                - os: where protocol stack is at
                - network card driver: control network card device
                - network card: actual transmission and reception of signal in network cable
            - browser use socket library to assign task to protocol stack
            - protocol stack has 2 part
                - TCP & UDP: to receive & send data (between application)
                - IP: to send & receive network packet, another 2 protocol in IP
                    - ICMP: information about error during network packet transfer
                    - ARP: use IP to query ethernet MAC address

        4. TCP
            - TCP add TCP header on data
            - HTTP is based on TCP
            - structure
            <br>
            <img src="https://networklessons.com/wp-content/uploads/2015/07/tcp-header.png" width="500">
                - source port (16 bits)
                    - identify the originating application on the sender's device.
                - destination port (16 bits)
                    - Ensures the data reaches the correct application on the receiving device
                - sequence number (32 bits)
                    - tracking, ordering packet
                - acknowledge number (32 bits)
                    - sent by receiver to ensure receive
                - data offset (4 bits)
                    - to locate where data begin
                - reserved (5 bits)
                - control flags (9 bits)
                    - URG
                    - ACK: reply
                    - PSH
                    - RST: to reconnect
                    - SYN: to establish connect
                    - FIN: to end connect
                - window size (16 bits)
                    - to control data flow, receiver & sender indicate its capacity
                - checksum (16 bits)
                - urgent pointer (16 bits)
                - options ( variable )
                - padding 
            - congestion control
                - congestion window ( cwnd)
                - not in TCP header
                - controlled by sender ( adjust dynamically based on network conditions)
            - 3 handshake
                - tcp connection before sending data
                - step:
                    1. both at `CLOSED` state, one of the server port in `LISTEN`
                    2. client send `SYN`, become `SYN-SENT` state
                    3. server receive, return `SYN` and `ACK`, become `SYN-RCVD`
                    4. client send `ACK`, become `ESTABLISHED`
                    4. server received `ACK` and become `ESTABLISHED`
                - to ensure both can receive and can send
            - to check TCP connection status:
                - linux: `netstat -napt`
                - <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/%E9%94%AE%E5%85%A5%E7%BD%91%E5%9D%80%E8%BF%87%E7%A8%8B/10.jpg" width="800">
            - TCP data segment
                - if HTTP data longer than `MSS`, TCP need to cut data into segment
                    - `MSS`: data only
                    - `MTU`: data + IP header+ TCP header, max MTU, 1500 byte
                    - each segment split out with its own TCP header
            - TCP packet Generation
                - 2 port in TCP:
                    1. browser listener port (random port number)
                    2. web server listener port, default 80(HTTP), 443 (HTTPS)
                - after establish connection, the data part in TCP packet will be http header + data
                - <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/%E9%94%AE%E5%85%A5%E7%BD%91%E5%9D%80%E8%BF%87%E7%A8%8B/13.jpg" width="600">

        5. IP
            - when TCP module establish/terminate connection, send/receive data, IP module encapsulate data into network packet
            - then send it out
            - IP packet header
            - <img src="https://networklessons.com/wp-content/uploads/2015/07/ip-packet-header-fields.png" width="500">
            - source & destination address:
                - source: client ip
                - destination: web server ip, that get from DNS domain resolution
            - protocol: if HTTP, it is throuhh TCP, protocol: 06 (hexadecimal)
            - if multiple network interface:
                - determine which destination to use by genmask the target address
                - then find the one that matched
                - if no match, go to 0.0.0.0
                - <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/%E9%94%AE%E5%85%A5%E7%BD%91%E5%9D%80%E8%BF%87%E7%A8%8B/15.jpg" width="500">
        - with IP, TCP, HTTP header
            - <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/%E9%94%AE%E5%85%A5%E7%BD%91%E5%9D%80%E8%BF%87%E7%A8%8B/17.jpg" width="500">

        6. MAC
            - need to add MAC header to internet packet
            - MAC is for ethernet
            - MAC header:
                1. receiver MAC address
                2. sender MAC address
                3. protocol
                    - 0800: IP protocol
                    - 0806: ARP protocol
            - how to get receiver MAC address?
                - find MAC adress of router
                - router foward packet to receiver, using ARP
            - ARP cache
                - system (desktop, router) will store MAC to avoid repeat checking
                - desktop can store router/printer/smart tv... MAC

        7. Network card
            - newtwork card and driver convert the binary data of internet packet into electric signal
            - network card driver copy the packet and put in internal buffer
                - add header and start frame delimiter at begining
                - append frame check sequence (FCS) at the end for error detection
            - network card has its own MAC address, and will check if destinatin MAC is itself ( which correct), it will discard if not mean to send to itself

        8. Switch
            - switch module convert electric signal to digital signal
            - use FCS to check if any error, if no, put into buffer
            - switch port wont chek destination MAC, switch port dont have its own MAC, it will place all received into buffer
            - check if MAC address is in MAC table
                - MAC table:
                    1. MAC address of device
                    2. switch port
                - if not found on siwtch table
                    1. maybe never sent any packet from that address to switch
                    2. maybe long time never used and deleted
                - switch will send packet to all port ( except source port)
                    - then only the correct received will respond, the other will ignore
                - receiver MAC can be Broadcast address
                    - this case packet will be send to all port except source
                    - in MAC: `FF:FF:FF:FF:FF:FF`
                    - in ip: `255.255.255.255`
        
        9. router
            - internet packet passed from switch to router
            - check table to decide where to send, similar to switch
            - but:
                1. router is IP based, 3 layer
                    - MAC & IP on all port
                2. switch is ethernet based, 2 layer
                    - no MAC on port
            - router port with IP & MAc, can receive from and send to ethernet
            - electric signal go to  network port, router change electric signal to mathematical signal, use FCS to check if ok
            - then check if MAC address is to itself, discard if not, place in buffer if yes
            - after receive, remove the MAC header

            
