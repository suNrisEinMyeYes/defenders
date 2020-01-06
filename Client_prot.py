import random
import json
def get_session_key():
	"""generate 10 char random string"""
	result = ""
	for i in range(1, 11):
		result += str(int(9 * random.random())+1)[0]
	return result

def get_hash_str():
	"""calculate initial hash string"""
	li = ""
	for i in range(5):
		li += str(int(int((6 * random.random()) + 1)))
	return li

# session protector class
class Session_protector:
	"""class used to protect web services from unauthorized access"""

	def __init__(self, hash_str):
		"""constructor"""
		self.__hash = hash_str

	def next_session_key(self, session_key):
		"""generate next session key"""
		## verify hashcode
		if self.__hash == "":
			raise Exception("Hash code is empty")

		for idx in range(len(self.__hash)):
			i = self.__hash[idx]
			if not str(i).isdigit():
				raise Exception("Hash code contains non-digit letter \"%c\"" % str(i))
		result = 0
		for idx in range(len(self.__hash)):
			i = self.__hash[idx]
			result += int(self.__calc_hash(session_key, int(i)))
		return ("0"*10 + str(result)[0:10])[-10:]

	def __calc_hash(self, session_key, val):
		"""calculate hash"""
		result = ""
		if val == 1:
			return ("00" + str(int(session_key[0:5]) % 97))[-2:]
		elif val == 2:
			for i in range(1, len(session_key)):
				result = result + session_key[i*(-1)]
			return str(result + session_key[0])
		elif val == 3:
			return session_key[-5:] + session_key[0:5]
		elif val == 4:
			num  = 0
			for i in range(1, 9):
				num += int(session_key[i]) + 41
			return str(num)
		elif val == 5:
			ch = ""
			num = 0
			for i in range(len(session_key)):
				ch = chr(ord(session_key[i]) ^ 43)
				if not ch.isdigit():
					ch = str(ord(ch))
				num += int(ch)
			return str(num)
		else:
			return str(int(session_key) + val)

with open("info_for_golang.json", "r") as read_file:
    data = json.load(read_file)

if data[stage] == "initial":
    hash_string = get_hash_str()
    skey_initial = get_session_key()
    protector = Session_protector(hash_string)
#stage 0
skey1 = protector1.next_session_key(skey_initial)
#stage 1 - each protector use it's own skey and compare with another's side
skey3 = protector1.next_session_key(skey1)
#stage 2
skey5 = protector1.next_session_key(skey3)
