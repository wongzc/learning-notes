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
            - TCP: 
                1. flow control
                2. timeout retransmission
                3. congestion control
            - UDP:
                1. faster
        - network support for application layer
        - data sent by application layer if bigger than MSS ( largest length og TCP segment), need to seperate into different segment
            - if any segment lost/ corrputed, resend only the segment
            - need to specify port on device to send, as many application on device is receving as well.

    3. network layer
        - IPV4, IPV6
        - the actual data transfer between device
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