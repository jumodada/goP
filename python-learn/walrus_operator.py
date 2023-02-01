list = ["a","b", "c"]

if(size := len(list) >= 3): print("总共有".format(size))

powers = [len(list)]

import re
desc = 'body: 18'

if(m := re.match('body:(.*)', desc)):
    age = m.group(1)
    print(age)




age = 1+ 1

name : bool = True

x: bytes = b"test"

print(age)

from typing import List,Set,Dict,Tuple

xx : List[str] = ['1']

xx.append('1')

dict1 : Dict[int, int] = {1: 1}

tuple1 : Tuple[int, ...] = (1,2,3)

print(len(tuple1))