import json

result=""
with open("./sellData.json",'r') as load_f:
    load_dict = json.load(load_f)
    i=0
    for price in load_dict["sell_price"]:

        result+=",(1,"+str(price)+",'"+load_dict["date"][i]+"')"
        i+=1

print(result)