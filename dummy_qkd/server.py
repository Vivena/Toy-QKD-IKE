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



# @app.get("/api/v1/keys/{slave_SAE_ID}/status")
# def api_status(slave_SAE_ID, request: Request):
#     tmp = None
#     master_SAE_ID = request.client.host
#     try:
#             # tmpi = next(
#             #     (i for i in range(len(config[master_SAE_ID]))
#             #         if config[master_SAE_ID][i]['slave_SAE_ID']
#             #         == slave_SAE_ID), None)
#         tmpi = next(
#             (i for i in range(len(config[master_SAE_ID]))))
#     except KeyError as e:
#         print('KeyError: Can\'t find key in config list:', e)
#         return {"message": "configuration access error. KeyError "}
#     if(tmpi is None):
#         return {"message": "Configuration access error. tmpi is None"}
#     tmp = config[master_SAE_ID][tmpi]
#     if(slave_SAE_ID in keys[master_SAE_ID]):
#         try:
#             tmp['stored_key_count'] = len(keys[master_SAE_ID][slave_SAE_ID])
#         except KeyError as e:
#             print('KeyError: Can\'t find key in keys list:', e)
#             return {"message": "configuration access error. KeyError "}
#     else:
#         tmp['stored_key_count'] = 0
#     tmp['master_SAE_ID'] = master_SAE_ID
#     return JSONResponse(content=tmp)


# @app.post("/api/v1/keys/{slave_SAE_ID}/enc_keys")
# async def api_getKey_post(slave_SAE_ID, request: Request):
#     tmp = {'Keys': []}
#     keyRequest = await request.json()
#     master_SAE_ID = request.client.host
#     for i in range(int(keyRequest['number'])):
#         key, keyid = getKey(keySize, path)
#         if(key):
#             tmpkey = {}
#             tmpkey['key'] = key.hex()
#             tmpkey['key_id'] = keyid
#             tmp['Keys'].append(tmpkey)
#         else:
#             return {"message": "Unable to generate the keys requested."}
#     return JSONResponse(content=tmp)


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


# @app.post("/api/v1/keys/{master_SAE_ID}/dec_keys")
# async def api_getKeyWithID_post(master_SAE_ID, request: Request):
#                           # keyRequestWithKeyID: KeyRequestWithKeyID):
#     global keyList
#     tmp = {'Keys': []}
#     keyRequestWithKeyID = await request.json()
#     for keyid in keyRequestWithKeyID["key_ids"]:
#         print("key_id: " + keyid["key_id"])
#         key, keyid = getKeyFromID(keyid["key_id"], keySize)
#         if(key):
#             tmpkey = {}
#             tmpkey['key'] = key.hex()
#             tmpkey['key_id'] = keyid
#             tmp['Keys'].append(tmpkey)
#         else:
#             return {"message": "Unable to generate the keys requested."}
#     return JSONResponse(content=tmp)


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