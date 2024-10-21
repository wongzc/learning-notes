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

    4. network access layer
        - OSI: data link, physical layer