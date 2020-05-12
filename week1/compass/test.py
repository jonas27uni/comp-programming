anfang = int(input())
end = int(input())

if anfang > end:
    if anfang - end < 180:
        print(anfang-end)
    else:
        print(-360+anfang-end)
else:
    if end - anfang < 180:
        print(end-anfang)
    else:
        print(-360+end-anfang)

# if anfang > end:
#     print("groeser")