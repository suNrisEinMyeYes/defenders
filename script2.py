import json
with open("info_for_golang.json", "r") as read_file:
    data = json.load(read_file)
data["stage"] = "0"
with open("info_for_golang.json", "w") as write_file:
    json.dump(data, write_file)
