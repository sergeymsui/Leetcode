
def main():
    n, k = [int(e) for e in input().split(" ")]
    arr = [int(e) for e in input().split(" ")]

    i = 0
    arrCount = 0
    mdict = {0:1}
    csumm = 0

    while i < n:
        csumm += arr[i]
        arrCount += mdict.get((csumm - k), 0)
        mdict[ csumm ] = mdict.get(csumm, 0) + 1
        i += 1
    
    print(arrCount)


if __name__ == '__main__':
    main()