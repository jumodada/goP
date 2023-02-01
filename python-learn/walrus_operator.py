list = ["a","b", "c"]

if(size := len(list) >= 3): print("总共有".format(size))

powers = [len(list)]

import re
desc = 'body: 18'

if(m := re.match('body:(.*)', desc)):
    age = m.group(1)
    print(age)



age = 1+ 1

print(age)