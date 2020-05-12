import cypher as cy


def test1():
    out = cy.exec([1,1],[0,1])
    if out != "1":
        print("test1 incorrect")


def test2():
    out = cy.exec([2,1],[3,2])
    if out != "010":
        print("test2 incorrect")
        print(out)

def test3():
    out = cy.exec([18,29],[13,21])
    if out != "10010101":
        print("test3 incorrect")
        print(out)

def test4():
    out = cy.exec([1,2999],[0,1])
    if len(out) != 120:
        print("test4 incorrect")
        print(len(out))

def test5():
    out = cy.exec([1,0],[0,1])
    if out != "":
        print("test3 incorrect")
        print(out)

test1()
test2()
test3()
test4()
test5()