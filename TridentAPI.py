# Trident 1.0.1 API Linux
import os

data = None
path = None
dae = open("./request.dat", "x")
called = False

def CallData():
    global data, path
    with open("./data.d") as dat:
        if dat:
            lines = dat.readlines()
            if lines:
                data = lines[0].strip()
                if len(lines) > 1:
                    path = lines[1].strip()

def Segment(name: str, content: any):
    if called == True:
        dae.write(f"{name}: {content}\n")
    return 

def ContentType(type: str):
    global called
    dae.write(f"{type}\n")
    called = True
    return

def Content(content: any):
    if called == True:
        dae.write(f"{content}")