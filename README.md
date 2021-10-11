# About this project

This project was heavily inspired by the excellent [wg-meshconf](https://github.com/k4yt3x/wg-meshconf)[^1][^2][^3]. Each project will produce about the same results, so you should definitely check it out as an option for solving your problem.

If you've used OpenVPN or similar, you're probably familiar with hub-and-spoke network architecture. That's when you have one server with many clients connected to it, and all traffic flows through the server. This creates a single point of failure and a performance bottleneck.

**This project generates configuration for `wg-quick` in a mesh architecture**, which is where each client will connect with each other client. One node may fail, but the others have a direct connection and thus are able to continue otherwise uninterrupted. As you can imagine, the direct path is much faster than going through an intermediate server. As well, wireguard itself is just plain fast.

[^1]: I have a slightly different take on how to store the data. You may use one or many YAML files to store the configuration data for this project (vs CSV files for `wg-meshconf`).
[^2]: This program is also written in Golang (vs Python for `wg-meshconf`).
[^3]: No source code was used, but I did generally copy the behavior of `wg-meshconf`.

# Almost everything you need to know

Say you you want to make a wireguard-based VPN in a mesh configuration. Perhaps you have a fleet of servers with public IP addresses (or a different port opened up for each host if you're behind a NAT), or you're simply on an untrusted network of some kind (wifi, colocation, etc). `wg-meshify` makes it really easy to generate mesh architecture config files which are ready to be copied to the servers.

Make sure you have `wg-quick` available. I'll also assume you're using `systemd`. Adjust as ncessary.

1. create mesh:

    ```
    wg-meshify mesh create --name mesh
    ```

2. add some peers:

    > Please note:
    >
    > Address = internal/vpn IP
    >
    > Entrypoint = external/Internet hostname or address and may be dynamic

    ```sh
    wg-meshify peer set \
        --mesh mesh \
        --address 192.168.30.1/32 \
        --entrypoint host1.domain.tld \
        --port 12345 \
        --name grey # pretty average peer

    wg-meshify peer set \
        --mesh mesh \
        --address 192.168.30.2/32 \
        --entrypoint host2.domain.tld \
        --port 13254 \
        --name blue # another average peer

    wg-meshify peer set \
        --mesh mesh \
        --address 192.168.30.3/32 \
        --entrypoint host3.domain.tld \
        --port 54321 \
        --name red # yet another average peer

    wg-meshify peer set \
        -m mesh \
        -a 192.168.30.4/32 \
        -e host4.domain.tld \
        -p 53421 \
        -n yellow # short args example
    ```

3. generate all the peer config files:

    ```
    wg-meshify genconf
    ```

4. copy each one to the appropriate place:

    ```sh
    rsync -ahvP ./mesh-peer.conf root@peer:"/etc/wireguard/mesh.conf" && \
    ssh -t root@peer sudo systemctl enable --now wg-quick@mesh.service
    ```

# Sample output

`cat ./mesh-blue.conf`:

```ini
[Interface]
# blue
Address = 192.168.30.2/32
PrivateKey = kEHP2w+ZPWW01n9qH//l/SIBPYWi1MHdkFufcbQCyFs=
ListenPort = 13254

[Peer]
# grey
PublicKey = 6mjzGnU/FkOtMr0c+q3GHOCZFlBUW6gMJ0dCxPq6Oz4=
Endpoint = host1.domain.tld:12345
AllowedIPs = 192.168.30.1/32

[Peer]
# red
PublicKey = Skh35wt+k8Vw7gmz7uMlGCjimrNR6VljP2+WIEwUzj0=
Endpoint = host3.domain.tld:54321
AllowedIPs = 192.168.30.3/32

[Peer]
# yellow
PublicKey = LIJRphwG54p2Opwn7+QCa03nUfIMtuEIRxA+bo8/TnU=
Endpoint = host4.domain.tld:53421
AllowedIPs = 192.168.30.4/32
```

`cat ./mesh-red.conf`:

```ini
[Interface]
# name red
Address = 192.168.30.3/32
PrivateKey = OBFcKrCPEvyIGebRpsmjvT6RMmAlmXBFyv8vr1CXuWg=
ListenPort = 54321

[Peer]
# blue
PublicKey = RJmb/pkGL53hJxiaCsoqN3QBAgBcgi1ZYXzbWe4iN1o=
Endpoint = host2.domain.tld:13254
AllowedIPs = 192.168.30.2/32

[Peer]
# grey
PublicKey = 6mjzGnU/FkOtMr0c+q3GHOCZFlBUW6gMJ0dCxPq6Oz4=
Endpoint = host1.domain.tld:12345
AllowedIPs = 192.168.30.1/32

[Peer]
# yellow
PublicKey = LIJRphwG54p2Opwn7+QCa03nUfIMtuEIRxA+bo8/TnU=
Endpoint = host4.domain.tld:53421
AllowedIPs = 192.168.30.4/32
```

etc.
