# Trident-XSS
A Hacking tool made for cross site scripting. V 1.0.2

To build just put: go build source.go

# Trident
Trident is a hacking tool for XSS attacks with simple python scripts that make the request, use at your own risk!

How it works:
So after specifying your flags, data and script, Trident will run the python script to create the request then it will connect with net.Dial() with a timeout (Default 10 seconds) after creating the request it will send it to the target.

Expectations for V 1.0.3:
    - Response.
    - Success?
