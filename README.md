# CodinGame - Death First Search - Episode 1
https://www.codingame.com/training/medium/death-first-search-episode-1

Your virus has caused a backdoor to open on the Bobnet network enabling you 
to send new instructions in real time.

You decide to take action by **stopping Bobnet from communicating on its own 
internal network**.

Bobnet's network is divided into several smaller networks, in each sub-network 
is a Bobnet agent tasked with transferring information by moving from node to 
node along links and **accessing gateways leading to other sub-networks**.

Your mission is to reprogram the virus so it will **sever links** in such a way 
that the Bobnet Agent is unable to access another sub-network thus preventing 
information concerning the presence of our virus to reach Bobnet's central hub.

### Rules
For each test you are given:
- A map of the network.
- The position of the exit gateways.
- The starting position of the Bobnet agent.

**>>> Nodes can only be connected to up to a single gateway. <<<**

Each game turn:
- First off, you sever one of the given links in the network.
- Then the Bobnet agent moves from one Node to another accessible Node.