import json
data = {
    "stage" : "initial",
    "string" : "",
    "key" : ""
}
with open("info_for_golang.json", "w") as write_file:
    json.dump(data, write_file)
