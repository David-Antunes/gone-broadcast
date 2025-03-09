# GONE-Broadcast

GONE example showing the capability of broadcast traffic between nodes connected to the same bridge.

## Deployment

To observe the experiment, just execute the `network.sh` script.

This script builds the image, and deploys a network with 10 nodes, 3 bridges and 1 router.

This network shows the broadcast with different bridge connections

* 5 nodes in bridge-1
* 3 nodes in bridge-2
* 2 nodes in bridge-3

By executing docker logs between gone-{1-10} you will see the IP addresses they find on the network.
