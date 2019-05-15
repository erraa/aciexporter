import requests
import json
import argparse
import pprint


from requests.packages.urllib3.exceptions import InsecureRequestWarning
requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

class Session(object):
    def __init__(self, ip, usr, pwd):
        'https://' + ip, usr, pwd
        self.usr = usr
        self.pwd = pwd
        self.url = 'https://' + ip
        self.verify = False
        self.cookies = None
        self._login()

    def _login(self):
        resp = self.post(
            '/api/aaaLogin.json',
            data={'aaaUser': {'attributes': {'name': self.usr, 'pwd': self.pwd}}}
        )
        self.cookies = {"APIC-Cookie": resp.cookies["APIC-cookie"]}


    #def __cookie__(self): return {'Set-Cookie' : self.token}

    def get(self, url):
        resp = requests.get(self.url + url, cookies=self.cookies, verify=self.verify)
        return resp.text

    def post(self, url, data):
        if not self.cookies:
            resp = requests.post(self.url + url, json.dumps(data), verify=self.verify)
        else:
            resp = requests.post(self.url + url, json.dumps(data), cookies=self.cookies, verify=self.verify)
        return resp

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("klass", help="add classname")
    args = parser.parse_args()

    s = Session("sandboxapicdc.cisco.com", "admin", "ciscopsdt")
    data = s.get("/api/class/{}.json".format(args.klass))
    print(data)

if __name__ == "__main__":
    main()
