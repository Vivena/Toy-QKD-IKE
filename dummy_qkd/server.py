import os
import uvicorn
from typing import Optional
from pydantic import BaseModel
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse

from dummy_key_gen import getkey, get_key_fom_id

app = FastAPI()
keySize = 256
SAE_ID = 'dummy_QKD'


class KeyRequest(BaseModel):
    number: Optional[str] = '1'
    size: Optional[str] = '1024'

class KeyRequestWithKeyID(BaseModel):
    key_ids: list = []


@app.get("/api/v1/keys/{slave_SAE_ID}/enc_keys")
def api_getkey_get(slave_SAE_ID,
                   number: int = 1, size: int = keySize):
    tmp = {'Keys': []}
    for i in range(number):
        key, keyid = getkey(size)
        if(key):
            tmpkey = {}

            tmpkey['key'] = key.hex()
            tmpkey['key_id'] = keyid
            tmp['Keys'].append(tmpkey)
        else:
            return {"message": "Unable to generate the keys requested."}
    return JSONResponse(content=tmp)



@app.get("/api/v1/keys/{master_SAE_ID}/dec_keys")
def api_getKeyWithID_get(master_SAE_ID, request: Request,
                         key_id: str = ""):

    tmp = {'Keys': []}
    print("key_id: " + key_id)
    key, keyid = get_key_fom_id(key_id)
    if(key):
        tmpkey = {}
        tmpkey['key'] = key.hex()
        tmpkey['key_id'] = keyid
        tmp['Keys'].append(tmpkey)
    else:
        return {"message": "Unable to generate the keys requested."}

    return JSONResponse(content=tmp)

    
if __name__ == '__main__':
    # TODO: use global variable instead of magic values for host, port and debug
    port = os.environ.get('PORT', '8000')
    uvicorn.run(app, host='0.0.0.0', port=int(port))