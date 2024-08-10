# Trident 1.0.2 API Linux
import os

data = None
path = None
dae = open("./request.dat", "x")
called = False
length = 0

def CallData():
    global data, path
    with open("./data.d") as dat:
        if dat:
            lines = dat.readlines()
            if lines:
                data = lines[0].strip()
                if len(lines) > 1:
                    path = open(lines[1].strip(), "r")

def Segment(name: str, content: any):
    global length
    if called == True:
        dae.write(f"{name}: {content}\n")
    return 

def ContentType(type: str):
    global called
    dae.write(f"{type}\n\n")
    called = True
    return

def Content(content: any):
    global length
    if called == True:
        dae.write(f"""{content}""")
        length += content.__len__()

def Exit():
    global length
    dae.write("LENGTH="+str(length))
    dae.close()
