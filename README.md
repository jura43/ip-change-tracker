# IP change tracker
I have a problem where the time when my public IP address changes sometimes changes unexpectedly. This caused me a lot of trouble when doing e-sports. While my ISP says there is nothing they can do about it, I needed a solution. So I created this small app that will notify you about this change by sending me an email.
## How does it work?
The app checks every minute if the IP address has changed. Then it checks if the IP address has changed and, in case if it did, it also checks if the current time of the day is different from the time when the IP address last changed. And then, in case that time of the day is different, it will send me an email from my Gmail address to the same Gmail address.
## How to run?
To run this you only need Go runtime and to enter credentials for login into the SMTP server.


