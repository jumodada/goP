def readFile():
    f= open('var_conv.py')
    print('start')
    sum = [1,2]
    try:
        for line in f:
            sum.append(line)
        _sum = sum
        return _sum
    finally:
        sum.append(12)
        print('end')



print(readFile())
