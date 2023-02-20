import os
import uuid

keys = {}

def getkey(keysize):

    keyid = str(uuid.uuid4())
    key = os.urandom(int(keysize/8))
    if not key:
        return key, ""
    else:
        keys[keyid] = key
        return key, keyid
    

def get_key_fom_id(keyid):

    if keyid in keys:
        return keys[keyid], keyid
    else:
        # last_position = newLast_position
        return "", ""